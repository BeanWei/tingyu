// from https://github.com/zu1k/nali
package iploc

import (
	_ "embed"
	"encoding/binary"
	"net"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

//go:embed qqwry.dat
var dat []byte

const (
	// redirectMode1 [IP][0x01][国家和地区信息的绝对偏移地址]
	redirectMode1 = 0x01
	// redirectMode2 [IP][0x02][信息的绝对偏移][...] or [IP][国家][...]
	redirectMode2 = 0x02
)

// Find 获取IP属地(中国IP返回省份其它IP返回国家)
func Find(s string) string {
	if s == "" {
		return ""
	}
	ip := net.ParseIP(s)
	if ip == nil {
		return ""
	}
	ipv4 := ip.To4()
	if ipv4 == nil {
		return ""
	}
	ipv4uint := binary.BigEndian.Uint32(ipv4)

	db := &ipdb{
		Data: dat,
	}
	offset := db.SearchIndex(ipv4uint)
	if offset <= 0 {
		return ""
	}

	var gbkLoc []byte

	mode := db.ReadMode(offset + 4)
	switch mode {
	case redirectMode1:
		locOffset := db.ReadUInt24()
		mode = db.ReadMode(locOffset)
		if mode == redirectMode2 {
			c := db.ReadUInt24()
			gbkLoc = db.ReadString(c)
			locOffset += 4
		} else {
			gbkLoc = db.ReadString(locOffset)
			locOffset += uint32(len(gbkLoc) + 1)
		}
	case redirectMode2:
		locOffset := db.ReadUInt24()
		gbkLoc = db.ReadString(locOffset)
	default:
		gbkLoc = db.ReadString(offset + 4)
	}

	loc, err := simplifiedchinese.GBK.NewDecoder().String(string(gbkLoc))
	if err != nil {
		return ""
	}
	return strings.ReplaceAll(loc, " CZ88.NET", "")
}

type ipdb struct {
	Data   []byte
	Offset uint32
}

// SearchIndex 查找索引位置
func (db *ipdb) SearchIndex(ip uint32) uint32 {
	header := db.ReadData(8, 0)

	start := binary.LittleEndian.Uint32(header[:4])
	end := binary.LittleEndian.Uint32(header[4:])

	mid := uint32(0)
	ipUint := uint32(0)

	for {
		mid = getMiddleOffset(start, end, 7)
		buf := db.ReadData(7, mid)
		ipUint = binary.LittleEndian.Uint32(buf[:4])

		if end-start == 7 {
			offset := byteToUInt32(buf[4:])
			buf = db.ReadData(7)
			if ip < binary.LittleEndian.Uint32(buf[:4]) {
				return offset
			}
			return 0
		}

		if ipUint > ip {
			end = mid
		} else if ipUint < ip {
			start = mid
		} else if ipUint == ip {
			return byteToUInt32(buf[4:])
		}
	}
}

// setOffset 设置偏移量
func (db *ipdb) SetOffset(offset uint32) {
	db.Offset = offset
}

// readString 获取字符串
func (db *ipdb) ReadString(offset uint32) []byte {
	db.SetOffset(offset)
	data := make([]byte, 0, 30)
	for {
		buf := db.ReadData(1)
		if buf[0] == 0 {
			break
		}
		data = append(data, buf[0])
	}
	return data
}

// readData 从文件中读取数据
func (db *ipdb) ReadData(length uint32, offset ...uint32) (rs []byte) {
	if len(offset) > 0 {
		db.SetOffset(offset[0])
	}

	end := db.Offset + length
	dataNum := uint32(len(db.Data))
	if db.Offset > dataNum {
		return nil
	}

	if end > dataNum {
		end = dataNum
	}
	rs = db.Data[db.Offset:end]
	db.Offset = end
	return
}

// readMode 获取偏移值类型
func (db *ipdb) ReadMode(offset uint32) byte {
	mode := db.ReadData(1, offset)
	return mode[0]
}

// ReadUInt24
func (db *ipdb) ReadUInt24() uint32 {
	buf := db.ReadData(3)
	return byteToUInt32(buf)
}

func getMiddleOffset(start uint32, end uint32, indexLen uint32) uint32 {
	records := ((end - start) / indexLen) >> 1
	return start + records*indexLen
}

func byteToUInt32(data []byte) uint32 {
	i := uint32(data[0]) & 0xff
	i |= (uint32(data[1]) << 8) & 0xff00
	i |= (uint32(data[2]) << 16) & 0xff0000
	return i
}

package biz

import "github.com/cloudwego/hertz/pkg/common/utils"

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func RespSuccess(data interface{}, total ...int) *Resp {
	if len(total) > 0 {
		return &Resp{
			Code: 0,
			Msg:  "success",
			Data: utils.H{
				"list":  data,
				"total": total[0],
			},
		}
	}
	return &Resp{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
}

func RespFail(bizCode int, errMsg ...string) *Resp {
	var msg string
	if len(errMsg) > 0 {
		msg = errMsg[0]
	} else {
		msg = CodeText(bizCode)
	}
	return &Resp{
		Code: bizCode,
		Msg:  msg,
	}
}

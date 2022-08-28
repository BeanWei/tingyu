package iploc

import (
	"testing"

	"github.com/ysmood/got"
)

func Test_Find(t *testing.T) {
	for _, data := range []struct {
		ip  string
		loc string
	}{
		{ip: "127.0.0.1", loc: "本机地址"},
		{ip: "117.169.96.198", loc: "江西省"},
		{ip: "114.44.227.87", loc: "台湾省台北市"},
		{ip: "123.23.3.0", loc: "越南"},
	} {
		got.T(t).Eq(Find(data.ip), data.loc)
	}
}

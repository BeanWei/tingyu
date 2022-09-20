package service

import (
	"testing"

	"github.com/duke-git/lancet/v2/random"
	"github.com/ysmood/got"
)

func Test_ValidateNickname(t *testing.T) {
	for _, c := range []struct {
		nickname string
		pass     bool
	}{
		{nickname: "", pass: false},
		{nickname: "ab", pass: false},
		{nickname: random.RandString(18), pass: false},
		{nickname: "你好_", pass: false},
		{nickname: "你好 o", pass: false},
		{nickname: "你好@世界", pass: false},
		{nickname: "你好世界", pass: true},
		{nickname: "你好Hello", pass: true},
		{nickname: "HelloWorld", pass: true},
		{nickname: "你好123", pass: true},
	} {
		got.T(t).Eq(ValidNickname(c.nickname) == nil, c.pass)
	}
}

package biz

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
)

type BizError struct {
	HttpCode int    // HTTP 状态码
	BizCode  int    // 业务码
	Msg      string // 错误描述
	Alert    bool   // 是否告警通知
}

type BizErrorOption func(be *BizError)

func (e *BizError) Error() string {
	return fmt.Sprintf(
		"HttpCode=%d\nBizCode=%d\nAlert=%t\nMsg=%s",
		e.HttpCode,
		e.BizCode,
		e.Alert,
		e.Msg,
	)
}

func Abort(c *app.RequestContext, bizCode int, err error, opts ...BizErrorOption) {
	c.Abort()
	c.Error(NewError(bizCode, err, opts...))
}

func AbortBizError(c *app.RequestContext, err *errors.Error) {
	c.Abort()
	c.Error(err)
}

func NewError(bizCode int, err error, opts ...BizErrorOption) *errors.Error {
	be := &BizError{
		HttpCode: Code2HttpCode(bizCode),
		BizCode:  bizCode,
		Msg:      CodeText(bizCode),
		Alert:    false,
	}
	for _, opt := range opts {
		opt(be)
	}
	return errors.New(err, errors.ErrorTypePublic, be)
}

func ErrWithAlert(alert bool) BizErrorOption {
	return func(be *BizError) {
		be.Alert = alert
	}
}

func ErrWithHttpCode(httpCode int) BizErrorOption {
	return func(be *BizError) {
		be.HttpCode = httpCode
	}
}

func ErrWithBizCode(bizCode int) BizErrorOption {
	return func(be *BizError) {
		be.BizCode = bizCode
	}
}

func ErrWithMsg(msg string) BizErrorOption {
	return func(be *BizError) {
		be.Msg = msg
	}
}

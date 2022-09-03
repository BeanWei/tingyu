package middleware

import (
	"context"

	"github.com/BeanWei/tingyu/g"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func ErrorHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Next(ctx)

		if err := c.Errors.Last(); err != nil {
			if g.Cfg().Server.IsDev {
				g.Dump(err)
			}
			if bizErr, ok := err.Meta.(*biz.BizError); ok {
				c.JSON(bizErr.HttpCode, biz.RespFail(
					bizErr.BizCode, bizErr.Msg,
				))
			} else {
				c.JSON(consts.StatusInternalServerError, biz.RespFail(
					biz.CodeServerError,
				))
			}
		}
	}
}

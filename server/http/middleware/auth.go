package middleware

import (
	"context"

	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Authentication 登录验证
func Authentication() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		ctxUser := shared.GetCtxUser(ctx)
		if ctxUser == nil {
			c.AbortWithError(
				consts.StatusUnauthorized,
				biz.NewError(biz.CodeNotAuthorized, errors.NewPublic("user not login")),
			)
		}
		c.Next(ctx)
	}
}

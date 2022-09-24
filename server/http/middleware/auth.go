package middleware

import (
	"context"

	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
)

// Authentication 登录验证
func Authentication() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		ctxUser := shared.GetCtxUser(ctx)
		if ctxUser == nil || ctxUser.Id == 0 {
			biz.Abort(c, biz.CodeNotAuthorized, errors.NewPublic("user not login"))
		}
		c.Next(ctx)
	}
}

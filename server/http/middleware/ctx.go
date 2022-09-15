package middleware

import (
	"context"

	"github.com/BeanWei/tingyu/http/jwt"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
)

// Ctx 自定义上下文初始化
func Ctx() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		claims, _ := jwt.ExtractClaims(ctx, c)
		if claims != nil {
			ctx = context.WithValue(ctx, shared.CtxSvcKey, &shared.Ctx{
				User: claims.UserInfo,
			})
		}
		c.Next(ctx)
	}
}

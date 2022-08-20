package middleware

import (
	"context"

	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

// Ctx 自定义上下文初始化
func Ctx() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		claims := jwt.ExtractClaims(ctx, c)
		if uid, ok := claims["id"].(int64); ok && uid > 0 {
			ctx = context.WithValue(ctx, shared.CtxSvcKey, &shared.Ctx{
				User: &shared.CtxUser{
					ID: uid,
				},
			})
		}
		c.Next(ctx)
	}
}

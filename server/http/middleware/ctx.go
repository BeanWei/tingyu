package middleware

import (
	"context"

	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

// Ctx 自定义上下文初始化
func Ctx(hzjwt *jwt.HertzJWTMiddleware) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		claims, _ := hzjwt.GetClaimsFromJWT(ctx, c)
		if exp, ok := claims["exp"].(float64); ok && int64(exp) >= hzjwt.TimeFunc().Unix() {
			if uid, ok := claims["id"].(float64); ok && uid > 0 {
				ctx = context.WithValue(ctx, shared.CtxSvcKey, &shared.Ctx{
					User: &shared.CtxUser{
						Id: int64(uid),
					},
				})
			}
		}
		c.Next(ctx)
	}
}

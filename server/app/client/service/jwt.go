package service

import (
	"context"
	"time"

	"github.com/BeanWei/tingyu/app/client/types"
	"github.com/BeanWei/tingyu/g"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

func JWT() *jwt.HertzJWTMiddleware {
	j, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "tingyu jwt",
		Key:         []byte(g.Cfg().JWT.SecretKey),
		Timeout:     time.Hour * 24 * time.Duration(g.Cfg().JWT.TimeoutDays),
		MaxRefresh:  time.Hour * 24,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*shared.CtxUser); ok {
				return jwt.MapClaims{
					"id":       v.ID,
					"is_admin": v.IsAdmin,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &shared.CtxUser{
				ID:      claims["id"].(int64),
				IsAdmin: claims["is_admin"].(bool),
			}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var req types.UserLoginReq
			if err := c.BindAndValidate(&req); err != nil {
				return nil, biz.NewError(biz.CodeParamBindError, err)
			}
			usr, err := UserLoginOrSignIn(ctx, &req)
			if err != nil {
				c.AbortWithError(consts.StatusBadRequest, err)
				return nil, err
			}
			return &shared.CtxUser{
				ID:      usr.ID,
				IsAdmin: usr.IsAdmin,
			}, nil
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if v, ok := data.(*shared.CtxUser); ok && v.ID != 0 {
				return true
			}
			return false
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			if err := c.Errors.Last(); err != nil {
				return
			}
			c.JSON(code, biz.RespFail(biz.CodeNotAuthorized))
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(code, biz.RespSuccess(utils.H{
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			}))
		},
		LogoutResponse: func(ctx context.Context, c *app.RequestContext, code int) {
			c.JSON(code, biz.RespSuccess(utils.H{}))
		},
	})
	if err != nil {
		panic(err)
	}
	return j
}

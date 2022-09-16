package api

import (
	"context"
	"fmt"
	"time"

	"github.com/BeanWei/tingyu/app/admin/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/user"
	"github.com/BeanWei/tingyu/http/jwt"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/BeanWei/tingyu/pkg/cryptor"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// UserLogin 管理员登录
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var req types.UserLoginReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeParamBindError, err))
		return
	}
	usr, err := ent.DB().User.Query().Where(user.UsernameEQ(req.Username)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeUserNotFound, err))
		} else {
			c.AbortWithError(consts.StatusInternalServerError, biz.NewError(biz.CodeServerError, err))
		}
		return
	}
	if !usr.IsAdmin {
		c.AbortWithError(consts.StatusForbidden, biz.NewError(biz.CodeForbidden, fmt.Errorf("user %s is no admin user", req.Username)))
		return
	}
	if cryptor.HashUserPwd(req.Password, usr.Salt) != usr.Password {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeInvalidPassword, errors.NewPublic("invalid password")))
		return
	}
	if token, expire, err := jwt.CreateToken(&shared.CtxUser{
		Id:      usr.ID,
		IsAdmin: usr.IsAdmin,
	}); err != nil {
		c.AbortWithError(biz.CodeServerError, err)
	} else {
		c.JSON(consts.StatusOK, biz.RespSuccess(utils.H{
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		}))
	}
}

// GetUserInfo 获取当前用户信息
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	uid := shared.GetCtxUser(ctx).Id
	usr, err := ent.DB().User.Get(ctx, uid)
	if err != nil {
		if ent.IsNotFound(err) {
			c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeUserNotFound, fmt.Errorf("user %d not found", uid)))
		} else {
			c.AbortWithError(consts.StatusInternalServerError, biz.NewError(biz.CodeServerError, err))
		}
		return
	}
	c.JSON(consts.StatusOK, biz.RespSuccess(utils.H{
		"id":       usr.ID,
		"username": usr.Username,
		"nickname": usr.Nickname,
		"avatar":   usr.Avatar,
	}))
}

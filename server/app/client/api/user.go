package api

import (
	"context"
	"fmt"

	"github.com/BeanWei/tingyu/app/client/service"
	"github.com/BeanWei/tingyu/app/client/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/duke-git/lancet/v2/random"
)

// SignIn 用户注册
func SignIn(ctx context.Context, c *app.RequestContext) {
	var req types.UserSignInReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeParamBindError, err))
		return
	}

	if err := service.ValidateUsername(ctx, req.Username); err != nil {
		c.AbortWithError(consts.StatusBadRequest, err)
		return
	}

	if err := service.ValidPassword(req.Password); err != nil {
		c.AbortWithError(consts.StatusBadRequest, err)
	}

	salt := random.RandString(10)
	ent.DB().User.Create().
		SetUsername(req.Username).
		SetNickname(req.Username).
		SetPassword(service.HashUserPwd(req.Password, salt)).
		SetSalt(salt).
		ExecX(ctx)

	c.JSON(consts.StatusOK, biz.RespSuccess(utils.H{}))
}

// GetUserInfo 获取用户基本信息
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var req types.GetUserInfoReq
	if err := c.Bind(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, err)
		return
	}

	uid := req.ID
	if uid == 0 {
		ctxuser := shared.GetCtxUser(ctx)
		if ctxuser != nil {
			uid = ctxuser.ID
		}
	}
	if uid <= 0 {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeUserNotFound, fmt.Errorf("user id %d is invalid", uid)))
	}
	usr, err := ent.DB().User.Get(ctx, uid)
	if err != nil {
		if ent.IsNotFound(err) {
			c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeUserNotFound, fmt.Errorf("user %d not found", uid)))
		} else {
			c.AbortWithError(consts.StatusInternalServerError, err)
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

// ChangePassword 修改密码

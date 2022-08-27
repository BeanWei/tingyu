package api

import (
	"context"
	"fmt"

	"github.com/BeanWei/tingyu/app/client/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetUserInfo 获取用户基本信息
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var req types.GetUserInfoReq
	if err := c.Bind(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, err)
		return
	}

	uid := req.ID
	if uid == 0 {
		ctxUser := shared.GetCtxUser(ctx)
		if ctxUser != nil {
			uid = ctxUser.ID
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

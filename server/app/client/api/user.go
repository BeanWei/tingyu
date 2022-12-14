package api

import (
	"context"
	"fmt"
	"time"

	"github.com/BeanWei/tingyu/app/client/dto"
	"github.com/BeanWei/tingyu/app/client/service"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/user"
	"github.com/BeanWei/tingyu/http/jwt"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// UserLogin 用户登录
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var req dto.UserLoginReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}
	usr, err := service.UserLoginOrSignIn(ctx, &req)
	if err != nil {
		biz.AbortBizError(c, err)
		return
	}
	if token, expire, err := jwt.CreateToken(&shared.CtxUser{
		Id:      usr.ID,
		IsAdmin: usr.IsAdmin,
	}); err != nil {
		biz.Abort(c, biz.CodeServerError, err)
	} else {
		c.JSON(200, biz.RespSuccess(utils.H{
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		}))
	}
}

// RefreshToken 刷新 Token
func RefreshToken(ctx context.Context, c *app.RequestContext) {
	token, expire, err := jwt.RefreshToken(ctx, c)
	if err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}
	c.JSON(200, biz.RespSuccess(utils.H{
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	}))
}

// GetUserInfo 获取用户基本信息
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var req dto.GetUserInfoReq
	if err := c.Bind(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	uid := req.Id
	if uid == 0 {
		ctxUser := shared.GetCtxUser(ctx)
		if ctxUser != nil {
			uid = ctxUser.Id
		}
	}
	if uid <= 0 {
		biz.Abort(c, biz.CodeNotAuthorized, fmt.Errorf("user id %d is invalid", uid))
		return
	}
	usr, err := ent.DB().User.Get(ctx, uid)
	if err != nil {
		if ent.IsNotFound(err) {
			biz.Abort(c, biz.CodeUserNotFound, fmt.Errorf("user %d not found", uid))
		} else {
			biz.Abort(c, biz.CodeServerError, err)
		}
		return
	}
	c.JSON(200, biz.RespSuccess(utils.H{
		"id":              usr.ID,
		"username":        usr.Username,
		"nickname":        usr.Nickname,
		"avatar":          usr.Avatar,
		"headline":        usr.Headline,
		"count_post":      usr.CountPost,
		"count_topic":     usr.CountTopic,
		"count_following": usr.CountFollowing,
		"count_follower":  usr.CountFollower,
	}))
}

// UpdateUserInfo 更新个人信息
func UpdateUserInfo(ctx context.Context, c *app.RequestContext) {
	var req dto.UpdateUserInfoReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	uid := shared.GetCtxUser(ctx).Id
	if ent.DB().User.Query().Where(user.IDNEQ(uid), user.NicknameEqualFold(req.Nickname)).ExistX(ctx) {
		biz.Abort(c, biz.CodeNicknameExisted, fmt.Errorf("nickname %s existed", req.Nickname))
		return
	}
	if err := service.ValidNickname(req.Nickname); err != nil {
		biz.AbortBizError(c, err)
		return
	}

	ent.DB().User.UpdateOneID(uid).
		SetAvatar(req.Avatar).
		SetNickname(req.Nickname).
		SetHeadline(req.Headline).
		ExecX(ctx)

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

// ChangePassword 修改密码

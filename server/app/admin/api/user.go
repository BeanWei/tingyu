package api

import (
	"context"

	"github.com/BeanWei/tingyu/app/admin/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/user"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ListUser 用户列表
func ListUser(ctx context.Context, c *app.RequestContext) {
	var req types.ListUserReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeParamBindError, err))
		return
	}

	query := ent.DB().User.Query().Where(user.DeletedAtEQ(0))
	if req.Filter.Status > 0 {
		query.Where(user.StatusEQ(req.Filter.Status))
	}
	if isAdmin := req.Filter.IsAdmin; isAdmin != nil {
		if *isAdmin == 0 {
			query.Where(user.IsAdminEQ(false))
		} else if *isAdmin == 1 {
			query.Where(user.IsAdminEQ(true))
		}
	}
	total := query.CountX(ctx)
	if total == 0 {
		c.JSON(consts.StatusOK, biz.RespSuccess(nil, total))
		return
	}
	topics := query.Limit(req.Limit).Offset(req.Offset()).AllX(ctx)

	c.JSON(consts.StatusOK, biz.RespSuccess(topics, total))
}

// UpdateUser 更新帖子
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	var req types.UpdateUserReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeParamBindError, err))
		return
	}

	ent.DB().User.UpdateOneID(req.Id).
		SetStatus(req.Status).
		ExecX(ctx)

	c.JSON(consts.StatusOK, biz.RespSuccess(utils.H{}))
}

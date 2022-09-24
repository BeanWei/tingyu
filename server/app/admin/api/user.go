package api

import (
	"context"

	"github.com/BeanWei/tingyu/app/admin/dto"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/user"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// ListUser 用户列表
func ListUser(ctx context.Context, c *app.RequestContext) {
	var req dto.ListUserReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
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
		c.JSON(200, biz.RespSuccess(nil, total))
		return
	}
	topics := query.Limit(req.Limit).Offset(req.Offset()).AllX(ctx)

	c.JSON(200, biz.RespSuccess(topics, total))
}

// UpdateUser 更新帖子
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	var req dto.UpdateUserReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	ent.DB().User.UpdateOneID(req.Id).
		SetStatus(req.Status).
		ExecX(ctx)

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

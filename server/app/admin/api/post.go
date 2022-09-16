package api

import (
	"context"

	"github.com/BeanWei/tingyu/app/admin/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/post"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// ListPost 帖子列表
func ListPost(ctx context.Context, c *app.RequestContext) {
	var req types.ListPostReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	query := ent.DB().Post.Query().Where(post.DeletedAtEQ(0))
	if req.Filter.Status > 0 {
		query.Where(post.StatusEQ(req.Filter.Status))
	}
	total := query.CountX(ctx)
	if total == 0 {
		c.JSON(200, biz.RespSuccess(nil, total))
		return
	}
	topics := query.Limit(req.Limit).Offset(req.Offset()).AllX(ctx)

	c.JSON(200, biz.RespSuccess(topics, total))
}

// UpdatePost 更新帖子
func UpdatePost(ctx context.Context, c *app.RequestContext) {
	var req types.UpdatePostReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	ent.DB().Post.UpdateOneID(req.Id).
		SetStatus(req.Status).
		ExecX(ctx)

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

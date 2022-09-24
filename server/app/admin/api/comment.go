package api

import (
	"context"

	"github.com/BeanWei/tingyu/app/admin/dto"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/comment"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// ListComment 评论列表
func ListComment(ctx context.Context, c *app.RequestContext) {
	var req dto.ListCommentReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	query := ent.DB().Comment.Query().Where(comment.DeletedAtEQ(0))
	if req.Filter.Status > 0 {
		query.Where(comment.StatusEQ(req.Filter.Status))
	}
	total := query.CountX(ctx)
	if total == 0 {
		c.JSON(200, biz.RespSuccess(nil, total))
		return
	}
	topics := query.Limit(req.Limit).Offset(req.Offset()).AllX(ctx)

	c.JSON(200, biz.RespSuccess(topics, total))
}

// UpdateComment 更新评论
func UpdateComment(ctx context.Context, c *app.RequestContext) {
	var req dto.UpdateCommentReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	ent.DB().Comment.UpdateOneID(req.Id).
		SetStatus(req.Status).
		ExecX(ctx)

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

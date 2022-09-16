package api

import (
	"context"

	"github.com/BeanWei/tingyu/app/admin/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/commentreply"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// ListCommentReply 回复列表
func ListCommentReply(ctx context.Context, c *app.RequestContext) {
	var req types.ListCommentReplyReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	query := ent.DB().CommentReply.Query().Where(commentreply.DeletedAtEQ(0))
	if req.Filter.Status > 0 {
		query.Where(commentreply.StatusEQ(req.Filter.Status))
	}
	total := query.CountX(ctx)
	if total == 0 {
		c.JSON(200, biz.RespSuccess(nil, total))
		return
	}
	topics := query.Limit(req.Limit).Offset(req.Offset()).AllX(ctx)

	c.JSON(200, biz.RespSuccess(topics, total))
}

// UpdateCommentReply 更新回复
func UpdateCommentReply(ctx context.Context, c *app.RequestContext) {
	var req types.UpdateCommentReplyReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	ent.DB().Post.UpdateOneID(req.Id).
		SetStatus(req.Status).
		ExecX(ctx)

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

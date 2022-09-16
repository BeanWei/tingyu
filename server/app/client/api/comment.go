package api

import (
	"context"
	"fmt"
	"time"

	"github.com/BeanWei/tingyu/app/client/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/comment"
	"github.com/BeanWei/tingyu/data/ent/commentreply"
	"github.com/BeanWei/tingyu/data/enums"
	"github.com/BeanWei/tingyu/g"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/BeanWei/tingyu/pkg/iploc"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// ListComment 评论列表
func ListComment(ctx context.Context, c *app.RequestContext) {
	var req types.ListCommentReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	query := ent.DB().Comment.Query().Where(
		comment.PostIDEQ(req.PostId),
		comment.DeletedAtEQ(0),
		comment.StatusEQ(enums.CommentStatusPass),
	)
	total := query.CountX(ctx)
	if total == 0 {
		c.JSON(200, biz.RespSuccess(nil, total))
		return
	}
	comments := query.WithUser().WithCommentReplies(
		func(crq *ent.CommentReplyQuery) {
			crq.WithUser().Order(ent.Asc(commentreply.FieldCreatedAt))
		},
	).Limit(req.Limit).Offset(req.Offset()).AllX(ctx)

	c.JSON(200, biz.RespSuccess(comments, total))
}

// CreateComment 发表评论
func CreateComment(ctx context.Context, c *app.RequestContext) {
	var req types.CreateCommentReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	var (
		postData = ent.DB().Post.GetX(ctx, req.PostId)
		ip       = c.ClientIP()
		uid      = shared.GetCtxUser(ctx).Id
	)

	status := enums.CommentStatusPass
	if g.Cfg().Operation.Audit {
		status = enums.CommentStatusAuditing
	}

	ent.DB().Comment.Create().
		SetStatus(status).
		SetPostID(postData.ID).
		SetUserID(uid).
		SetIP(ip).
		SetIPLoc(iploc.Find(ip)).
		SetContent(req.Content).
		SetIsPoster(postData.UserID == uid).
		ExecX(ctx)

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

// DeleteComment 删除评论
func DeleteComment(ctx context.Context, c *app.RequestContext) {
	var req types.DeleteCommentReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	var (
		commentData = ent.DB().Comment.GetX(ctx, req.Id)
		ctxUser     = shared.GetCtxUser(ctx)
	)
	if commentData.UserID != ctxUser.Id && !ctxUser.IsAdmin {
		biz.Abort(c, biz.CodeForbidden, fmt.Errorf("user %d forbidden to delete comment %d", ctxUser.Id, req.Id))
		return
	}
	commentData.Update().SetDeletedAt(time.Now().Unix()).ExecX(ctx)

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

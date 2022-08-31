package api

import (
	"context"
	"fmt"
	"time"

	"github.com/BeanWei/tingyu/app/client/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/comment"
	"github.com/BeanWei/tingyu/data/ent/commentreply"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/BeanWei/tingyu/pkg/iploc"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ListComment 评论列表
func ListComment(ctx context.Context, c *app.RequestContext) {
	var req types.ListCommentReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeParamBindError, err))
		return
	}

	query := ent.DB().Comment.Query().Where(
		comment.PostIDEQ(req.PostId),
		comment.DeletedAtEQ(0),
	)
	total := query.CountX(ctx)
	if total == 0 {
		c.JSON(consts.StatusOK, biz.RespSuccess(nil, total))
		return
	}
	comments := query.WithUser().WithCommentReplies(
		func(crq *ent.CommentReplyQuery) {
			// 带上前5条回复
			crq.Limit(5).Order(ent.Asc(commentreply.FieldCreatedAt))
		},
	).Limit(req.Limit).Offset(req.Offset()).AllX(ctx)

	c.JSON(consts.StatusOK, biz.RespSuccess(comments, total))
}

// CreateComment 发表评论
func CreateComment(ctx context.Context, c *app.RequestContext) {
	var req types.CreateCommentReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeParamBindError, err))
		return
	}

	var (
		postData = ent.DB().Post.GetX(ctx, req.PostId)
		ip       = c.ClientIP()
		uid      = shared.GetCtxUser(ctx).ID
	)
	ent.DB().Comment.Create().
		SetPostID(postData.ID).
		SetUserID(uid).
		SetIP(ip).
		SetIPLoc(iploc.Find(ip)).
		SetContent(req.Content).
		SetIsPoster(postData.UserID == uid).
		ExecX(ctx)

	c.JSON(consts.StatusOK, biz.RespSuccess(utils.H{}))
}

// DeleteComment 删除评论
func DeleteComment(ctx context.Context, c *app.RequestContext) {
	var req types.DeleteCommentReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeParamBindError, err))
		return
	}

	var (
		commentData = ent.DB().Comment.GetX(ctx, req.Id)
		ctxUser     = shared.GetCtxUser(ctx)
	)
	if commentData.UserID != ctxUser.ID && !ctxUser.IsAdmin {
		c.AbortWithError(consts.StatusForbidden, biz.NewError(
			biz.CodeForbidden,
			fmt.Errorf("user %d forbidden to delete comment %d", ctxUser.ID, req.Id)),
		)
		return
	}
	commentData.Update().SetDeletedAt(time.Now().Unix()).ExecX(ctx)

	c.JSON(consts.StatusOK, biz.RespSuccess(utils.H{}))
}

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

// ListCommentReply 回复列表
func ListCommentReply(ctx context.Context, c *app.RequestContext) {
	var req types.ListCommentReplyReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	query := ent.DB().CommentReply.Query().Where(
		commentreply.CommentIDEQ(req.CommentId),
		commentreply.DeletedAtEQ(0),
		commentreply.StatusEQ(enums.CommentReplyStatusPass),
	)
	total := query.CountX(ctx)
	if total == 0 {
		c.JSON(200, biz.RespSuccess(nil, total))
		return
	}
	replies := query.WithUser().Limit(req.Limit).Offset(req.Offset()).AllX(ctx)

	c.JSON(200, biz.RespSuccess(replies, total))
}

// CreateCommentReply 发表回复
func CreateCommentReply(ctx context.Context, c *app.RequestContext) {
	var req types.CreateCommentReplyReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	var (
		commentData = ent.DB().Comment.Query().
				Where(comment.IDEQ(req.CommentId)).
				WithPost().
				OnlyX(ctx)
		ip  = c.ClientIP()
		uid = shared.GetCtxUser(ctx).Id
	)

	status := enums.CommentReplyStatusPass
	if g.Cfg().Operation.Audit {
		status = enums.CommentReplyStatusAuditing
	}

	res := ent.DB().CommentReply.Create().
		SetStatus(status).
		SetUserID(uid).
		SetIP(ip).
		SetIPLoc(iploc.Find(ip)).
		SetContent(req.Content).
		SetCommentID(commentData.ID).
		SetToUserID(req.ToUserId).
		SetToReplyID(req.ToReplyId).
		SetIsPoster(commentData.Edges.Post.UserID == uid).
		SaveX(ctx)

	c.JSON(200, biz.RespSuccess(res))
}

// DeleteCommentReply 删除回复
func DeleteCommentReply(ctx context.Context, c *app.RequestContext) {
	var req types.DeleteCommentReplyReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	var (
		replyData = ent.DB().CommentReply.GetX(ctx, req.Id)
		ctxUser   = shared.GetCtxUser(ctx)
	)
	if replyData.UserID != ctxUser.Id && !ctxUser.IsAdmin {
		biz.Abort(c, biz.CodeForbidden, fmt.Errorf("user %d forbidden to delete reply %d", ctxUser.Id, req.Id))
		return
	}
	replyData.Update().SetDeletedAt(time.Now().Unix()).ExecX(ctx)

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

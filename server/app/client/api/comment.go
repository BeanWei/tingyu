package api

import (
	"context"
	"fmt"
	"time"

	"github.com/BeanWei/tingyu/app/client/dto"
	"github.com/BeanWei/tingyu/app/client/service"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/comment"
	"github.com/BeanWei/tingyu/data/ent/commentreply"
	"github.com/BeanWei/tingyu/data/ent/userreaction"
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
	var req dto.ListCommentReq
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
	records := query.WithUser().WithCommentReplies(
		func(crq *ent.CommentReplyQuery) {
			crq.WithUser().Order(ent.Asc(commentreply.FieldCreatedAt))
		},
	).Limit(req.Limit).Offset(req.Offset()).AllX(ctx)

	ids := make([]int64, 0)
	for _, record := range records {
		ids = append(ids, record.ID)
		for _, record2 := range record.Edges.CommentReplies {
			ids = append(ids, record2.ID)
		}
	}
	reactions, err := service.GetReactionsForManySubject(
		ctx, shared.GetCtxUser(ctx).Id, ids,
	)
	if err != nil {
		biz.Abort(c, biz.CodeServerError, err)
		return
	}
	results := make([]*dto.Comment, len(records))
	for i, record := range records {
		results2 := make([]*dto.CommentReply, len(record.Edges.CommentReplies))
		for j, record2 := range record.Edges.CommentReplies {
			results2[j] = &dto.CommentReply{
				ID:        record2.ID,
				CreatedAt: record2.CreatedAt,
				UpdatedAt: record.UpdatedAt,
				UserID:    record2.UserID,
				IPLoc:     record2.IPLoc,
				Content:   record2.Content,
				CommentID: record2.CommentID,
				ToUserID:  record2.ToUserID,
				ToReplyID: record2.ToReplyID,
				IsPoster:  record2.IsPoster,
				User:      record2.Edges.User,
				Reactions: reactions[record2.ID],
			}
			results[i] = &dto.Comment{
				ID:             record.ID,
				CreatedAt:      record.CreatedAt,
				UpdatedAt:      record.UpdatedAt,
				PostID:         record.PostID,
				UserID:         record.UserID,
				IPLoc:          record.IPLoc,
				Content:        record.Content,
				ReplyCount:     record.ReplyCount,
				IsPoster:       record.IsPoster,
				User:           record.Edges.User,
				Reactions:      reactions[record.ID],
				CommentReplies: results2,
			}
		}
	}

	c.JSON(200, biz.RespSuccess(results, total))
}

// CreateComment 发表评论
func CreateComment(ctx context.Context, c *app.RequestContext) {
	var req dto.CreateCommentReq
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

	g.Pool().Submit(func() {
		ent.DB().Post.UpdateOneID(postData.ID).SetLatestRepliedAt(time.Now().Unix()).ExecX(context.Background())
	})

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

// ReactComment 收藏或点赞评论
func ReactComment(ctx context.Context, c *app.RequestContext) {
	var req dto.ReactCommentReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	uid := shared.GetCtxUser(ctx).Id

	if reaction := ent.DB().UserReaction.Query().Where(
		userreaction.SubjectTypeEQ(userreaction.SubjectTypeComment),
		userreaction.SubjectIDEQ(req.Id),
		userreaction.UserIDEQ(uid),
		userreaction.ReactCodeEQ(req.Code),
	).FirstX(ctx); reaction != nil {
		ent.DB().UserReaction.DeleteOneID(reaction.ID).ExecX(ctx)
	} else {
		ent.DB().UserReaction.Create().
			SetUserID(uid).
			SetSubjectType(userreaction.SubjectTypeComment).
			SetSubjectID(req.Id).
			SetReactCode(req.Code).
			ExecX(ctx)
	}

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

// DeleteComment 删除评论
func DeleteComment(ctx context.Context, c *app.RequestContext) {
	var req dto.DeleteCommentReq
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

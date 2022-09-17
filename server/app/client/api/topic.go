package api

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/BeanWei/tingyu/app/client/service"
	"github.com/BeanWei/tingyu/app/client/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/topic"
	"github.com/BeanWei/tingyu/data/ent/user"
	"github.com/BeanWei/tingyu/data/enums"
	"github.com/BeanWei/tingyu/g"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// ListTopic 话题列表
func ListTopic(ctx context.Context, c *app.RequestContext) {
	var req types.ListTopicReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	query := ent.DB().Topic.Query().Where(topic.DeletedAtEQ(0), topic.StatusEQ(enums.TopicStatusOnline))
	if req.IsRec {
		query.Where(topic.IsRecEQ(true))
	}
	if req.CategoryId != 0 {
		query.Where(topic.TopicCategoryIDEQ(req.CategoryId))
	}
	if req.UserId != 0 {
		query.Where(topic.HasUsersWith(user.IDEQ(req.UserId)))
	} else if ctxUser := shared.GetCtxUser(ctx); ctxUser != nil && ctxUser.Id > 0 {
		query.Where(func(s *sql.Selector) {
			s.Where(sql.P(func(b *sql.Builder) {
				b.WriteString(`"topics"."id" NOT IN (SELECT "user_topics"."topic_id" FROM "user_topics" WHERE "user_topics"."user_id" = `).Arg(ctxUser.Id).WriteString(")")
			}))
		})
	}
	total := query.CountX(ctx)
	if total == 0 {
		c.JSON(200, biz.RespSuccess(nil, total))
		return
	}
	topics := query.Order(
		ent.Asc(topic.FieldRecRank), ent.Desc(topic.FieldCreatedAt),
	).Limit(req.Limit).Offset(req.Offset()).AllX(ctx)

	c.JSON(200, biz.RespSuccess(topics, total))
}

// SearchTopic 搜索话题
func SearchTopic(ctx context.Context, c *app.RequestContext) {
	var req types.SearchTopicReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	query := ent.DB().Topic.Query().Where(topic.DeletedAtEQ(0), topic.StatusEQ(enums.TopicStatusOnline))
	if req.Keyword != "" {
		query.Where(topic.TitleContainsFold(req.Keyword))
	}
	topics := query.Order(
		ent.Asc(topic.FieldRecRank), ent.Desc(topic.FieldCreatedAt),
	).Limit(20).AllX(ctx)

	c.JSON(200, biz.RespSuccess(topics, len(topics)))
}

// CreateTopic 创建话题
func CreateTopic(ctx context.Context, c *app.RequestContext) {
	var req types.CreateTopicReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}
	if titleLength := len(req.Title); titleLength > 20 {
		biz.Abort(c, biz.CodeInvalidTopicTitle, fmt.Errorf("topic length is %d, over than 20", titleLength))
		return
	}

	status := enums.TopicStatusOnline
	if g.Cfg().Operation.Audit {
		status = enums.TopicStatusAuditing
	}

	res := ent.DB().Topic.Create().
		SetTitle(req.Title).
		SetIcon(req.Icon).
		SetDescription(req.Description).
		SetCreatorID(shared.GetCtxUser(ctx).Id).
		SetStatus(status).
		SaveX(ctx)
	g.Pool().Submit(func() {
		ent.DB().User.UpdateOneID(res.CreatorID).AddCountTopic(1).AddTopicIDs(res.ID).ExecX(context.Background())
	})

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

// FollowTopic 关注话题
func FollowTopic(ctx context.Context, c *app.RequestContext) {
	var req types.FollowTopicReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	if !ent.DB().Topic.Query().Where(topic.IDEQ(req.Id)).ExistX(ctx) {
		biz.Abort(c, biz.CodeParamBindError, fmt.Errorf("topic %d is not found in db", req.Id))
		return
	}

	uid := shared.GetCtxUser(ctx).Id
	if service.TopicIsFollowed(ctx, req.Id, uid) {
		biz.Abort(c, biz.CodeTopicIsFollowed, fmt.Errorf("user %d followed topic %d repeat", uid, req.Id))
		return
	}
	ent.DB().User.UpdateOneID(uid).AddCountTopic(1).AddTopicIDs(req.Id).ExecX(ctx)

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

// UnFollowTopic 取关话题
func UnFollowTopic(ctx context.Context, c *app.RequestContext) {
	var req types.UnFollowTopicReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	if !ent.DB().Topic.Query().Where(topic.IDEQ(req.Id)).ExistX(ctx) {
		biz.Abort(c, biz.CodeParamBindError, fmt.Errorf("topic %d is not found in db", req.Id))
		return
	}

	uid := shared.GetCtxUser(ctx).Id
	if !service.TopicIsFollowed(ctx, req.Id, uid) {
		biz.Abort(c, biz.CodeTopicIsNotFollowed, fmt.Errorf("user %d not followed topic %d", uid, req.Id))
		return
	}
	ent.DB().User.UpdateOneID(uid).AddCountTopic(-1).RemoveTopicIDs(req.Id).ExecX(ctx)

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

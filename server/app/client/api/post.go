package api

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/BeanWei/tingyu/app/client/service"
	"github.com/BeanWei/tingyu/app/client/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/post"
	"github.com/BeanWei/tingyu/data/ent/topic"
	"github.com/BeanWei/tingyu/data/ent/userreaction"
	"github.com/BeanWei/tingyu/data/enums"
	"github.com/BeanWei/tingyu/g"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/BeanWei/tingyu/pkg/iploc"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/meilisearch/meilisearch-go"
)

// ListPost 帖子列表
func ListPost(ctx context.Context, c *app.RequestContext) {
	var req types.ListPostReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	query := ent.DB().Post.Query().Where(post.DeletedAtEQ(0), post.StatusEQ(enums.PostStatusPass))
	if req.TopicId != 0 {
		query.Where(post.HasTopicsWith(topic.IDEQ(req.TopicId)))
	}
	if req.UserId != 0 {
		if req.Pinned {
			query.Where(func(s *sql.Selector) {
				s.Where(sql.P(func(b *sql.Builder) {
					b.WriteString(`"posts"."id" IN (
						SELECT "user_reactions"."subject_id"
						FROM "user_reactions"
						WHERE "user_reactions"."subject_type" = 'post'
						AND "user_reactions"."react_code" = 'action-pin'
						AND "user_reactions"."user_id" = `).Arg(req.UserId).WriteString(")")
				}))
			})
		} else if req.Starred {
			query.Where(func(s *sql.Selector) {
				s.Where(sql.P(func(b *sql.Builder) {
					b.WriteString(`"posts"."id" IN (
						SELECT "user_reactions"."subject_id"
						FROM "user_reactions"
						WHERE "user_reactions"."subject_type" = 'post'
						AND "user_reactions"."react_code" = 'action-star'
						AND "user_reactions"."user_id" = `).Arg(req.UserId).WriteString(")")
				}))
			})
		} else {
			query.Where(post.UserIDEQ(req.UserId))
		}
	}
	total := query.CountX(ctx)
	if total == 0 {
		c.JSON(200, biz.RespSuccess(nil, total))
		return
	}
	if req.SortType == 1 {
		// TODO: 热度值计算
		query.Order(ent.Desc(post.FieldCommentCount))
	} else if req.SortType == 2 {
		query.Order(ent.Desc(post.FieldCreatedAt))
	}
	records := query.WithUser().
		Limit(req.Limit).
		Offset(req.Offset()).
		AllX(ctx)

	ids := make([]int64, len(records))
	for i, record := range records {
		ids[i] = record.ID
	}
	reactions, err := service.GetReactionsForManySubject(
		ctx, shared.GetCtxUser(ctx).Id, userreaction.SubjectTypePost, ids,
	)
	if err != nil {
		biz.Abort(c, biz.CodeServerError, err)
		return
	}
	results := make([]*types.Post, len(records))
	for i, record := range records {
		results[i] = &types.Post{
			ID:              record.ID,
			CreatedAt:       record.CreatedAt,
			UpdatedAt:       record.UpdatedAt,
			CommentCount:    record.CommentCount,
			IsTop:           record.IsTop,
			IsExcellent:     record.IsExcellent,
			IsLock:          record.IsLock,
			LatestRepliedAt: record.LatestRepliedAt,
			IPLoc:           record.IPLoc,
			Content:         record.Content,
			UserID:          record.UserID,
			User:            record.Edges.User,
			Reactions:       reactions[record.ID],
		}
	}
	c.JSON(200, biz.RespSuccess(results, total))
}

// GetPost 帖子详情
func GetPost(ctx context.Context, c *app.RequestContext) {
	var req types.GetPostReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	record := ent.DB().Post.Query().Where(post.IDEQ(req.Id)).WithUser().OnlyX(ctx)
	reactions, err := service.GetReactionsForOneSubject(
		ctx, shared.GetCtxUser(ctx).Id, userreaction.SubjectTypePost, record.ID,
	)
	if err != nil {
		biz.Abort(c, biz.CodeServerError, err)
		return
	}
	result := &types.Post{
		ID:              record.ID,
		CreatedAt:       record.CreatedAt,
		UpdatedAt:       record.UpdatedAt,
		CommentCount:    record.CommentCount,
		IsTop:           record.IsTop,
		IsExcellent:     record.IsExcellent,
		IsLock:          record.IsLock,
		LatestRepliedAt: record.LatestRepliedAt,
		IPLoc:           record.IPLoc,
		Content:         record.Content,
		UserID:          record.UserID,
		User:            record.Edges.User,
		Reactions:       reactions,
	}

	c.JSON(200, biz.RespSuccess(result))
}

// CreatePost 发表帖子
func CreatePost(ctx context.Context, c *app.RequestContext) {
	var req types.CreatePostReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	status := enums.PostStatusPass
	if g.Cfg().Operation.Audit {
		status = enums.PostStatusAuditing
	}
	ip := c.ClientIP()
	uid := shared.GetCtxUser(ctx).Id

	res := ent.DB().Post.Create().
		SetStatus(status).
		SetUserID(uid).
		SetIP(ip).
		SetIPLoc(iploc.Find(ip)).
		SetContent(req.Content).
		AddTopicIDs(req.TopicIds...).
		SaveX(ctx)
	g.Meili().Index(post.Table).AddDocuments([]map[string]any{
		{
			"id":      res.ID,
			"content": req.ContentText,
		},
	})
	if len(req.TopicIds) > 0 {
		g.Pool().Submit(func() {
			ent.DB().Topic.Update().Where(topic.IDIn(req.TopicIds...)).AddPostCount(1).Exec(context.Background())
		})
	}
	g.Pool().Submit(func() {
		ent.DB().User.UpdateOneID(uid).AddCountPost(1).Exec(context.Background())
	})

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

// DeletePost 删除帖子

// SearchPost 搜索帖子
func SearchPost(ctx context.Context, c *app.RequestContext) {
	var req types.SearchPostReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	res, err := g.Meili().Index(post.Table).Search(req.Keyword, &meilisearch.SearchRequest{
		Limit:  int64(req.Limit),
		Offset: int64(req.Offset()),
	})
	if err != nil {
		biz.Abort(c, biz.CodeServerError, err)
		return
	}
	if res.EstimatedTotalHits == 0 {
		c.JSON(200, biz.RespSuccess(nil, 0))
	}
	ids := make([]int64, len(res.Hits))
	for i, hit := range res.Hits {
		id, _ := convertor.ToInt(hit.(map[string]any)["id"])
		ids[i] = id
	}
	records := ent.DB().Post.Query().Unique(false).Where(
		post.DeletedAtEQ(0),
		post.StatusEQ(enums.PostStatusPass),
		post.IDIn(ids...),
	).WithUser().Order(func(s *sql.Selector) {
		s.OrderExpr(sql.P(func(b *sql.Builder) {
			b.WriteString("array_position(").Arg(ids).WriteString(", id)")
		}))
	}).AllX(ctx)

	c.JSON(200, biz.RespSuccess(records, int(res.EstimatedTotalHits)))
}

// ReactPost 收藏或点赞帖子
func ReactPost(ctx context.Context, c *app.RequestContext) {
	var req types.ReactPostReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	uid := shared.GetCtxUser(ctx).Id

	if reaction := ent.DB().UserReaction.Query().Where(
		userreaction.SubjectTypeEQ(userreaction.SubjectTypePost),
		userreaction.SubjectIDEQ(req.Id),
		userreaction.UserIDEQ(uid),
		userreaction.ReactCodeEQ(req.Code),
	).FirstX(ctx); reaction != nil {
		ent.DB().UserReaction.DeleteOneID(reaction.ID).ExecX(ctx)
	} else {
		ent.DB().UserReaction.Create().
			SetUserID(uid).
			SetSubjectType(userreaction.SubjectTypePost).
			SetSubjectID(req.Id).
			SetReactCode(req.Code).
			ExecX(ctx)
	}

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

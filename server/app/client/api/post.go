package api

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/BeanWei/tingyu/app/client/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/post"
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
		query.Where(post.ContentContains(fmt.Sprintf(`"mentionName":"%d"`, req.TopicId)))
	}
	total := query.CountX(ctx)
	if total == 0 {
		c.JSON(200, biz.RespSuccess(nil, total))
		return
	}
	if req.SortType == 0 {
		query.Order(ent.Desc(post.FieldCreatedAt))
	} else if req.SortType == 1 {
		// TODO: 热度值计算
		query.Order(ent.Desc(post.FieldCommentCount))
	}
	records := query.WithUser().
		Limit(req.Limit).
		Offset(req.Offset()).
		AllX(ctx)

	c.JSON(200, biz.RespSuccess(records, total))
}

// GetPost 帖子详情
func GetPost(ctx context.Context, c *app.RequestContext) {
	var req types.GetPostReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	res := ent.DB().Post.Query().Where(post.IDEQ(req.Id)).WithUser().OnlyX(ctx)

	c.JSON(200, biz.RespSuccess(res))
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
	res := ent.DB().Post.Create().
		SetStatus(status).
		SetUserID(shared.GetCtxUser(ctx).Id).
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

// CollectPost 收藏帖子

// ReactPost .

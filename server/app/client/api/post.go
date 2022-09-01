package api

import (
	"context"

	"github.com/BeanWei/tingyu/app/client/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/post"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/BeanWei/tingyu/pkg/iploc"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ListPost 帖子列表
func ListPost(ctx context.Context, c *app.RequestContext) {
	var req types.ListPostReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeParamBindError, err))
		return
	}

	query := ent.DB().Post.Query()
	total := query.CountX(ctx)
	if total == 0 {
		c.JSON(consts.StatusOK, biz.RespSuccess(nil, total))
		return
	}
	if req.SortType == 0 {
		query.Order(ent.Desc(post.FieldCreatedAt))
	}
	posts := query.WithUser().
		Limit(req.Limit).
		Offset(req.Offset()).
		AllX(ctx)

	c.JSON(consts.StatusOK, biz.RespSuccess(posts, total))
}

// GetPost 帖子详情
func GetPost(ctx context.Context, c *app.RequestContext) {
	var req types.GetPostReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeParamBindError, err))
		return
	}

	post := ent.DB().Post.Query().Where(post.IDEQ(req.Id)).WithUser().OnlyX(ctx)

	c.JSON(consts.StatusOK, biz.RespSuccess(post))
}

// CreatePost 发表帖子
func CreatePost(ctx context.Context, c *app.RequestContext) {
	var req types.CreatePostReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeParamBindError, err))
		return
	}

	ip := c.ClientIP()
	ent.DB().Post.Create().
		SetUserID(shared.GetCtxUser(ctx).ID).
		SetIP(ip).
		SetIPLoc(iploc.Find(ip)).
		SetContent(req.Content).
		AddTopicIDs(req.TopicIds...).
		ExecX(ctx)

	c.JSON(consts.StatusOK, biz.RespSuccess(utils.H{}))
}

// DeletePost 删除帖子

// SearchPost 搜索帖子

// CollectPost 收藏帖子

// ReactPost .

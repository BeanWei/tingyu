package api

import (
	"context"

	"github.com/BeanWei/tingyu/app/admin/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/topic"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ListTopic 话题列表
func ListTopic(ctx context.Context, c *app.RequestContext) {
	var req types.ListTopicReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeParamBindError, err))
		return
	}

	query := ent.DB().Topic.Query().Where(topic.DeletedAtEQ(0))
	if req.Filter.Status > 0 {
		query.Where(topic.StatusEQ(req.Filter.Status))
	}
	total := query.CountX(ctx)
	if total == 0 {
		c.JSON(consts.StatusOK, biz.RespSuccess(nil, total))
		return
	}
	topics := query.WithTopicCategory().Limit(req.Limit).Offset(req.Offset()).AllX(ctx)

	c.JSON(consts.StatusOK, biz.RespSuccess(topics, total))
}

// CreateTopic 创建话题
func CreateTopic(ctx context.Context, c *app.RequestContext) {
	var req types.CreateTopicReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeParamBindError, err))
		return
	}

	ent.DB().Topic.Create().
		SetTitle(req.Title).
		SetIcon(req.Icon).
		SetDescription(req.Description).
		SetStatus(req.Status).
		SetTopicCategoryID(req.TopicCategoryId).
		SetIsRec(req.IsRec).
		SetRecRank(req.RecRank).
		ExecX(ctx)

	c.JSON(consts.StatusOK, biz.RespSuccess(utils.H{}))
}

// UpdateTopic 更新话题
func UpdateTopic(ctx context.Context, c *app.RequestContext) {
	var req types.UpdateTopicReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeParamBindError, err))
		return
	}

	ent.DB().Topic.UpdateOneID(req.Id).
		SetTitle(req.Title).
		SetIcon(req.Icon).
		SetDescription(req.Description).
		SetStatus(req.Status).
		SetTopicCategoryID(req.TopicCategoryId).
		SetIsRec(req.IsRec).
		SetRecRank(req.RecRank).
		ExecX(ctx)

	c.JSON(consts.StatusOK, biz.RespSuccess(utils.H{}))
}

package api

import (
	"context"

	"github.com/BeanWei/tingyu/app/admin/dto"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/topiccategory"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// ListTopicCategory 话题分类列表
func ListTopicCategory(ctx context.Context, c *app.RequestContext) {
	var req dto.ListTopicCategoryReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	query := ent.DB().TopicCategory.Query().Where(topiccategory.DeletedAtEQ(0))
	total := query.CountX(ctx)
	if total == 0 {
		c.JSON(200, biz.RespSuccess(nil, total))
		return
	}
	categories := query.Limit(req.Limit).Offset(req.Offset()).AllX(ctx)

	c.JSON(200, biz.RespSuccess(categories, total))
}

// CreateTopicCategory 创建话题分类
func CreateTopicCategory(ctx context.Context, c *app.RequestContext) {
	var req dto.CreateTopicCategoryReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	ent.DB().TopicCategory.Create().
		SetName(req.Name).
		SetRank(req.Rank).
		ExecX(ctx)

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

// UpdateTopicCategory 创建话题分类
func UpdateTopicCategory(ctx context.Context, c *app.RequestContext) {
	var req dto.UpdateTopicCategoryReq
	if err := c.BindAndValidate(&req); err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	ent.DB().TopicCategory.UpdateOneID(req.Id).
		SetName(req.Name).
		SetRank(req.Rank).
		ExecX(ctx)

	c.JSON(200, biz.RespSuccess(utils.H{}))
}

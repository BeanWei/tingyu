package api

import (
	"context"

	"github.com/BeanWei/tingyu/app/client/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/topiccategory"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/cloudwego/hertz/pkg/app"
)

// ListTopicCategory 话题分类列表
func ListTopicCategory(ctx context.Context, c *app.RequestContext) {
	var req types.ListTopicCategoryReq
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
	categories := query.AllX(ctx)

	c.JSON(200, biz.RespSuccess(categories, total))
}

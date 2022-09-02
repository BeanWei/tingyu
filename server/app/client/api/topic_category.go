package api

import (
	"context"

	"github.com/BeanWei/tingyu/app/client/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/topiccategory"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ListTopicCategory 话题分类列表
func ListTopicCategory(ctx context.Context, c *app.RequestContext) {
	var req types.ListTopicCategoryReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeParamBindError, err))
		return
	}

	query := ent.DB().TopicCategory.Query().Where(topiccategory.DeletedAtEQ(0))
	total := query.CountX(ctx)
	if total == 0 {
		c.JSON(consts.StatusOK, biz.RespSuccess(nil, total))
		return
	}
	categories := query.Limit(req.Limit).Offset(req.Offset()).AllX(ctx)

	c.JSON(consts.StatusOK, biz.RespSuccess(categories, total))
}

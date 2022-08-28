package api

import (
	"context"
	"fmt"

	"github.com/BeanWei/tingyu/app/client/types"
	"github.com/BeanWei/tingyu/data/ent"
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

	query := ent.DB().Topic.Query()
	total := query.CountX(ctx)
	if total == 0 {
		c.JSON(consts.StatusOK, biz.RespSuccess(nil, total))
		return
	}
	topics := query.Limit(req.Limit).Offset(req.Offset()).AllX(ctx)

	c.JSON(consts.StatusOK, biz.RespSuccess(topics, total))
}

// CreateTopic 创建话题
func CreateTopic(ctx context.Context, c *app.RequestContext) {
	var req types.CreateTopicReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(biz.CodeParamBindError, err))
		return
	}
	if titleLength := len(req.Title); titleLength > 20 {
		c.AbortWithError(consts.StatusBadRequest, biz.NewError(
			biz.CodeInvalidTopicTitle, fmt.Errorf("topic length is %d, over than 20", titleLength),
		))
		return
	}

	ent.DB().Topic.Create().
		SetTitle(req.Title).
		SetIcon(req.Icon).
		SetDescription(req.Description).
		SetNotice(req.Notice).
		ExecX(ctx)

	c.JSON(consts.StatusOK, biz.RespSuccess(utils.H{}))
}

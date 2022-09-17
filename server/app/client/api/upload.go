package api

import (
	"context"
	"strings"

	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/g"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/BeanWei/tingyu/pkg/oss"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func Upload(ctx context.Context, c *app.RequestContext) {
	file, err := c.FormFile("file")
	if err != nil {
		biz.Abort(c, biz.CodeParamBindError, err)
		return
	}

	fileType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(fileType, "image") {
		biz.Abort(c, biz.CodeFileTypeUnSupport, err)
		return
	}

	if max := g.Cfg().OSS.MaxMB; max > 0 && file.Size > int64(max*1024*1024) {
		biz.Abort(c, biz.CodeFileSizeTooLarge, err)
		return
	}

	output, err := oss.New().PutObject(ctx, file)
	if err != nil {
		biz.Abort(c, biz.CodeServerError, err)
		return
	}
	ent.DB().File.Create().
		SetUserID(shared.GetCtxUser(ctx).Id).
		SetSize(file.Size).
		SetType(fileType).
		SetName(output.Filename).
		SetURL(output.URL).
		SetOrigName(file.Filename).
		SetOssType(output.OSSType).
		SetOssDomain(output.OSSDomain).
		SetOssBucket(output.OSSBucket).
		ExecX(ctx)

	c.JSON(200, biz.RespSuccess(utils.H{
		"url": output.URL,
	}))
}

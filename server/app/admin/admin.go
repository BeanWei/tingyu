package admin

import (
	"context"

	"github.com/BeanWei/tingyu/app/admin/api"
	"github.com/BeanWei/tingyu/http/middleware"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Register(s *server.Hertz) {
	sg := s.Group("/api/admin")

	sg.POST("/auth/login", api.UserLogin)

	sg.Use(middleware.Ctx(), middleware.Authentication(), func(ctx context.Context, c *app.RequestContext) {
		if !shared.GetCtxUser(ctx).IsAdmin {
			c.AbortWithError(
				consts.StatusForbidden,
				biz.NewError(biz.CodeForbidden, errors.NewPublic("no permission")),
			)
		}
		c.Next(ctx)
	})

	sg.GET("/auth/info", api.GetUserInfo)
	sg.GET("/category/list", api.ListTopicCategory)
	sg.POST("/category/create", api.CreateTopicCategory)
	sg.PUT("/category/update", api.UpdateTopicCategory)
	sg.GET("/topic/list", api.ListTopic)
	sg.POST("/topic/create", api.CreateTopic)
	sg.PUT("/topic/update", api.UpdateTopic)
	sg.GET("/post/list", api.ListPost)
	sg.PUT("/post/update", api.UpdatePost)
	sg.GET("/comment/list", api.ListComment)
	sg.PUT("/comment/update", api.UpdateComment)
	sg.GET("/reply/list", api.ListCommentReply)
	sg.PUT("/reply/update", api.UpdateCommentReply)
	sg.GET("/user/list", api.ListUser)
	// sg.PUT("/user/update", api.UpdateUser)
}

package client

import (
	"github.com/BeanWei/tingyu/app/client/api"
	"github.com/BeanWei/tingyu/app/client/service"
	"github.com/BeanWei/tingyu/http/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(s *server.Hertz) {
	jwt := service.JWT()

	apiv1 := s.Group("/api/v1")
	{
		// user
		apiv1.POST("/user/login", jwt.LoginHandler)
		apiv1.GET("/user/get", middleware.Ctx(jwt), api.GetUserInfo)
		// post
		apiv1.GET("/post/list", api.ListPost)
		apiv1.GET("/post/get", api.GetPost)
		// topic
		apiv1.GET("/topic/list", api.ListTopic)

		apiv1.Use(middleware.Ctx(jwt), middleware.Authentication())

		// user
		apiv1.POST("/user/logout", jwt.LogoutHandler)
		apiv1.GET("/user/refresh_token", jwt.RefreshHandler)
		// post
		apiv1.POST("/post/create", api.CreatePost)
		// topic
		apiv1.POST("/topic/create", api.CreateTopic)
	}
}

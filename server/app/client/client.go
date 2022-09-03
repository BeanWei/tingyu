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
		// topic
		apiv1.GET("/topic/list", middleware.Ctx(jwt), api.ListTopic)
		apiv1.GET("/topic/search", middleware.Ctx(jwt), api.SearchTopic)
		// post
		apiv1.GET("/post/list", api.ListPost)
		apiv1.GET("/post/get", api.GetPost)
		// comment
		apiv1.GET("/comment/list", api.ListComment)
		// reply
		apiv1.GET("/reply/list", api.ListCommentReply)

		apiv1.Use(middleware.Ctx(jwt), middleware.Authentication())

		// user
		apiv1.POST("/user/logout", jwt.LogoutHandler)
		apiv1.GET("/user/refresh_token", jwt.RefreshHandler)
		// topic
		apiv1.POST("/topic/create", api.CreateTopic)
		apiv1.POST("/topic/follow", api.FollowTopic)
		apiv1.DELETE("/topic/unfollow", api.UnFollowTopic)
		// post
		apiv1.POST("/post/create", api.CreatePost)
		// comment
		apiv1.POST("/comment/create", api.CreateComment)
		apiv1.DELETE("/comment/delete", api.DeleteComment)
		// reply
		apiv1.POST("/reply/create", api.CreateCommentReply)
		apiv1.DELETE("/reply/delete", api.DeleteCommentReply)
	}
}

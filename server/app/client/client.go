package client

import (
	"github.com/BeanWei/tingyu/app/client/api"
	"github.com/BeanWei/tingyu/http/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(s *server.Hertz) {
	apiv1 := s.Group("/api/v1")
	{
		// user
		apiv1.POST("/user/login", api.UserLogin)
		apiv1.GET("/user/get", middleware.Ctx(), api.GetUserInfo)
		// category
		apiv1.GET("/category/list", api.ListTopicCategory)
		// topic
		apiv1.GET("/topic/list", middleware.Ctx(), api.ListTopic)
		apiv1.GET("/topic/search", middleware.Ctx(), api.SearchTopic)
		// post
		apiv1.GET("/post/list", api.ListPost)
		apiv1.GET("/post/search", api.SearchPost)
		apiv1.GET("/post/get", api.GetPost)
		// comment
		apiv1.GET("/comment/list", api.ListComment)
		// reply
		apiv1.GET("/reply/list", api.ListCommentReply)

		apiv1.Use(middleware.Ctx(), middleware.Authentication())

		// user
		apiv1.GET("/user/refresh_token", api.RefreshToken)
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

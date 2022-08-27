package client

import (
	"github.com/BeanWei/tingyu/app/client/api"
	"github.com/BeanWei/tingyu/app/client/service"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(s *server.Hertz) {
	jwt := service.JWT()

	apiv1 := s.Group("/api/v1")
	apiv1.POST("/user/login", jwt.LoginHandler)
	apiv1.GET("/user/get", api.GetUserInfo)

	apiv1.Use(jwt.MiddlewareFunc())

	apiv1.POST("/user/logout", jwt.LogoutHandler)
	apiv1.GET("/user/refresh_token", jwt.RefreshHandler)
}

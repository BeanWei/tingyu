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
	apiv1.POST("/user/login", jwt.LoginHandler)
	apiv1.GET("/user/get", middleware.Ctx(jwt), api.GetUserInfo)

	apiv1.Use(middleware.Ctx(jwt), middleware.Authentication())

	apiv1.POST("/user/logout", jwt.LogoutHandler)
	apiv1.GET("/user/refresh_token", jwt.RefreshHandler)
}

package client

import (
	clientapi "github.com/BeanWei/tingyu/app/client/api"
	clientsvc "github.com/BeanWei/tingyu/app/client/service"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Rgister(s *server.Hertz) {
	jwt := clientsvc.JWT()

	apiv1 := s.Group("/api/v1")
	apiv1.POST("/user/signin", clientapi.SignIn)
	apiv1.POST("/user/login", jwt.LoginHandler)
	apiv1.GET("/user/get", clientapi.GetUserInfo)

	apiv1.Use(jwt.MiddlewareFunc())

	apiv1.POST("/user/logout", jwt.LogoutHandler)
	apiv1.GET("/user/refresh_token", jwt.RefreshHandler)
}

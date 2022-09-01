package app

import (
	"time"

	"github.com/BeanWei/tingyu/app/client"
	"github.com/BeanWei/tingyu/g"
	"github.com/BeanWei/tingyu/http/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/panjf2000/ants/v2"

	_ "github.com/BeanWei/tingyu/data/hook"
)

func NewHTTPServer() {
	defer ants.Release()

	svr := server.Default(func() []config.Option {
		opts := make([]config.Option, 0)
		if g.Cfg().Server.Address != "" {
			opts = append(opts, server.WithHostPorts(g.Cfg().Server.Address))
		}
		if g.Cfg().Server.IsDev {
			opts = append(opts, server.WithExitWaitTime(time.Second*0))
		}
		return opts
	}()...)

	// 注册全局中间件
	svr.Use(
		middleware.ErrorHandler(),
		middleware.CORS(),
	)

	client.Register(svr)

	svr.Spin()
}

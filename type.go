package rest

import (
	"context"
	"fmt"
	"net/http"
)

type (

	// handlerFund 定义实际处理请求的函数
	HandlerFunc func(*Context) error

	// Route 最终挂载给 http 服务的函数
	Route struct {
		Method  string
		Path    string
		Handler HandlerFunc
	}

	Context interface {
		context.Context

		OkJSON()
		ErrorJSON()
		GetRequestID() string
		Set(string, interface{})
		Get(string) (interface{}, bool)

		init(http.ResponseWriter, *http.Request)
		setHandlers(...HandlerFunc)
		next()
	}

	// Middleware 中间件函数
	Middleware func(next HandlerFunc) HandlerFunc

	// Server 服务实体
	Server interface {
		// Run 启动
		Run() error
		// Use 使用中间件
		Use(middlewares ...HandlerFunc) error
		// AddRoute 挂载路由
		AddRoute(route ...Route) error
	}
)

// NewServer 创建一个服务实例
func NewServer(host string, port int) Server {
	server := &serverImp{
		httpServer: &http.Server{
			Addr: fmt.Sprintf("%s:%d", host, port),
		},
		enginez: newEnginez(),
	}
	server.httpServer.Handler = server.enginez
	return server
}

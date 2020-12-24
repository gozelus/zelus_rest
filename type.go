package rest

import (
	"fmt"
	"github.com/gozelus/zelus_rest/core"
	"github.com/gozelus/zelus_rest/logger"
	"net/http"
	path2 "path"
	"time"
)

type (
	// Route 最终挂载给 http 服务的函数
	Route struct {
		Method  string
		Path    string
		Handler http.HandlerFunc
	}

	// Middleware 中间件函数
	Middleware func(next http.HandlerFunc) http.HandlerFunc

	// Server 服务实体
	Server interface {
		// Run 启动
		Run() error
		// Use 使用中间件
		Use(middlewares ...Middleware) error
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
		engine: core.NewEnginez(),
	}
	server.httpServer.Handler = server.engine
	return server
}

var (
	LoggerMiddleware Middleware = func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			method := request.Method
			path := path2.Clean(request.URL.Path)
			now := time.Now()
			next(writer, request)
			logger.Infof("[%s]-[%s]-[%dms]", method, path, now.Sub(time.Now()).Milliseconds())
		}
	}
)

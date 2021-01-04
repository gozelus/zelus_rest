package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gozelus/zelus_rest/logger"
)

type (
	// StatusError
	// 用于拓展错误，GetCode 返回值将会作为 httpCode
	StatusError interface {
		error
		GetCode() int
		GetMessage() string
	}

	// handlerFund 定义实际处理请求的函数
	HandlerFunc func(Context)

	// Route 最终挂载给 http 服务的函数
	Route struct {
		Method  string
		Path    string
		Handler HandlerFunc
	}

	Context interface {
		context.Context

		Headers() map[string][]string
		Method() string
		Path() string
		JSONBodyBind(v interface{}) error
		JSONQueryBind(v interface{}) error

		Next()

		// RenderOkJSON
		// 表示成功处理请求，返回 HttpStatusOk
		RenderOkJSON(data interface{})

		// RenderErrorJSON
		// 内部尝试断言 error 为 StatusError，若断言成功
		// HttpCode 会使用此 error 的 GetCode 返回值
		// 若不成功，会使用 InternalServerErrorCode
		RenderErrorJSON(data interface{}, err error)

		GetRequestID() string
		GetError() error
		Set(string, interface{})
		Get(string) (interface{}, bool)

		init(http.ResponseWriter, *http.Request)
		setHandlers(...HandlerFunc)
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

type Option = func(imp *Plugin)
type Plugin struct {
	Logger   HandlerFunc
	Recovery HandlerFunc
}

// NewServer 创建一个服务实例
func NewServer(port int, opts ...Option) Server {
	server := &serverImp{
		httpServer: &http.Server{
			Addr: fmt.Sprintf("%s:%d", "localhost", port),
		},
		enginez: newEnginez(),
	}
	p := &Plugin{
		Logger:   nil,
		Recovery: nil,
	}
	for _, o := range opts {
		o(p)
	}
	server.httpServer.Handler = server.enginez
	server.plugin = p

	// 启动之前，检查下是否注入了 Logger 和 Recovery
	if server.plugin.Logger == nil {
		server.plugin.Logger = func(c Context) {
			now := time.Now()
			c.Next()
			if err := c.GetError(); err != nil {
				logger.WarnfWithContext(c, "method : %s | path : %s | duration : %d ms | err : %T -> %+v", c.Method(), c.Path(), now.Sub(now).Milliseconds(), err, err)
			} else {
				logger.InfofWithContext(c, "method : %s | path : %s | duration : %d ms", c.Method(), c.Path(), now.Sub(now).Milliseconds())
			}
		}
		server.use(server.plugin.Logger)
	}
	if server.plugin.Recovery == nil {
		server.plugin.Recovery = func(c Context) {
			defer func() {
				if err := recover(); err != nil {
					logger.ErrorfWithStackWithContext(c, "recover err : %s", err)
					c.RenderErrorJSON(nil, statusInternalServerError)
				}
			}()
			c.Next()
		}
		server.use(server.plugin.Recovery)
	}
	return server
}

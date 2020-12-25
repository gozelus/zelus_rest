package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gozelus/zelus_rest/logger"
)

// 一些预定好的错误
var (
	StatusMethodNotAllowed = Error{
		Code:    http.StatusMethodNotAllowed,
		Message: "method not allowed",
	}
	StatusNotFound = Error{
		Code:    http.StatusNotFound,
		Message: "not found",
	}
	StatusInternalServerError = Error{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
	}
)

func (e Error) ErrorCode() int {
	return e.Code
}
func (e Error) ErrorMessage() string {
	return e.Message
}

type (
	// handlerFund 定义实际处理请求的函数
	HandlerFunc func(Context) ErrorInterface

	// Route 最终挂载给 http 服务的函数
	Route struct {
		Method  string
		Path    string
		Handler HandlerFunc
	}

	Error struct {
		Code    int    `json:"code_code"`
		Message string `json:"error_message"`
	}

	// ErrorInterface 用于扩展 Error
	ErrorInterface interface {
		// ErrorCode 此方法的返回值将会作为 http code
		// 也会映射为结构体中的 error_code 字段
		ErrorCode() int
		// ErrorMessage 映射为返回结构体中的 error_message 字段
		ErrorMessage() string
	}

	Context interface {
		context.Context

		Headers() map[string][]string
		Method() string
		Path() string

		Next()
		OkJSON(interface{})
		ErrorJSON(ErrorInterface)
		GetRequestID() string
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
func NewServer(host string, port int, opts ...Option) Server {
	server := &serverImp{
		httpServer: &http.Server{
			Addr: fmt.Sprintf("%s:%d", host, port),
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
		server.plugin.Logger = func(c Context) ErrorInterface {
			now := time.Now()
			c.Next()
			logger.InfofWithContext(c, "method : %s | path : %s | duration : %d ms", c.Method(), c.Path(), now.Sub(now).Milliseconds())
			return nil
		}
		server.use(server.plugin.Logger)
	}
	if server.plugin.Recovery == nil {
		server.plugin.Recovery = func(c Context) ErrorInterface {
			defer func() {
				if err := recover(); err != nil {
					logger.ErrorfWithStackWithContext(c, "recover err : %s", err)
					c.ErrorJSON(StatusInternalServerError)
				}
			}()
			return nil
		}
		server.use(server.plugin.Recovery)
	}
	return server
}

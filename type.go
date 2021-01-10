package rest

import (
	"context"
	"fmt"
	"github.com/gozelus/zelus_rest/core"
	"github.com/pkg/errors"
	"io"
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
		GetReason() interface {
			GetReasonMessage() string
			GetReasonCode() int
		}
	}

	// handlerFund 定义实际处理请求的函数
	HandlerFunc func(Context)

	// Route 最终挂载给 http 服务的函数
	Route struct {
		Method  string
		Path    string
		Handler HandlerFunc
	}

	// 用于控制流
	Context interface {
		context.Context

		Headers() map[string][]string
		Method() string
		Path() string
		// @Authored handler 的返回值才会有意义
		UserID() int64
		// 如果内部调用了此方法，会尝试生成 or 刷新一个 jwt 给客户端
		SetUserID(int64)

		File(name string) (io.Reader, error)
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
		setUserID(int64)
		setJwtUtils(utils jwtUtils)
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
	Authored HandlerFunc // 默认实现为 jwt
	// 用于设置 jwt ak
	JwtAk func() (string, string)
}

// 初始化一个 context
func BackgroundContext() Context {
	c := newContext()
	c.init(nil, nil)
	return c
}

// NewServer 创建一个服务实例
func NewServer(port int, opts ...Option) Server {
	server := &serverImp{
		httpServer: &http.Server{
			Addr: fmt.Sprintf("%s:%d", "0.0.0.0", port),
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
				logger.WarnfWithContext(c, "method : %s | path : %s | duration : %d ms | err : %T -> %+v", c.Method(), c.Path(), now.Sub(now).Milliseconds(), errors.Cause(err), err)
			} else {
				duration := time.Now().Sub(now).Milliseconds()
				if duration > 300 {
					logger.WarnfWithContext(c, "slow request -> method : %s | path : %s | duration : %d ms", c.Method(), c.Path(), time.Now().Sub(now).Milliseconds())
				} else {
					logger.InfofWithContext(c, "ok request -> method : %s | path : %s | duration : %d ms", c.Method(), c.Path(), time.Now().Sub(now).Milliseconds())
				}
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
	if server.plugin.Authored == nil {
		if server.plugin.JwtAk == nil {
			server.plugin.JwtAk = func() (s string, s2 string) {
				return "hi", "hello"
			}
		}
		server.jwtUtils = core.NewJwtUtils(server.plugin.JwtAk())
		server.plugin.Authored = func(c Context) {
			c.setJwtUtils(server.jwtUtils)
			if token, ok := c.Headers()["Authorization"]; ok && len(token) > 0 && len(token[0]) > 0 {
				userID, newTokenStr, err := server.jwtUtils.ValidateToken(token[0])
				if err != nil {
					c.RenderErrorJSON(nil, statusUnauthorized)
					return
				}
				c.setUserID(userID)
				c.Next()
				c.Set("jwt-token", newTokenStr)
				return
			}
			c.RenderErrorJSON(nil, statusUnauthorized)
			return
		}
		server.use(server.plugin.Authored)
	}
	return server
}

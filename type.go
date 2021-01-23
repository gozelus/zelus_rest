package rest

import (
	"context"
	"fmt"
	"github.com/gozelus/zelus_rest/core"
	"github.com/gozelus/zelus_rest/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"net/http"
	"time"
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
		WithReason(string, int) StatusError
	}

	// handlerFund 定义实际处理请求的函数
	HandlerFunc func(Context)

	// Route 最终挂载给 http 服务的函数
	Route struct {
		Method             string
		Path               string
		Handler            HandlerFunc
		NeedAuthentication bool
	}

	// 用于控制流
	Context interface {
		context.Context

		Headers() map[string][]string
		Method() string
		Path() string
		HttpCode() int
		QueryMap() map[string]string
		RequestBodyJsonStr() string
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

		// private
		init(http.ResponseWriter, *http.Request, *time.Duration)
		setHandlers(...HandlerFunc)
		setUserID(int64)
		setJwtToken(string)
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
	Logger      HandlerFunc
	Recovery    HandlerFunc
	Authored    func(HandlerFunc) HandlerFunc // 默认实现为 jwt
	Metrics     HandlerFunc                   // prometheus
	MetricsPort int

	// 用于设置 jwt ak 过期时间
	JwtAk func() (string, int64, int64)
	// 设置请求限时，默认1000ms
	ReqTimeOut *time.Duration
}

// 初始化一个 context
func BackgroundContext() Context {
	c := newContext()
	c.init(nil, nil, nil)
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

	if server.plugin.MetricsPort == 0 {
		server.plugin.MetricsPort = 8888
	}

	// 开启 prometheus 监听
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		logger.Infof("prometheus list in %s ...", server.plugin.MetricsPort)
		if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", server.plugin.MetricsPort), nil); err != nil {
			panic(err)
		}
	}()

	if server.plugin.ReqTimeOut == nil {
		timeOut := time.Millisecond * 1000
		server.plugin.ReqTimeOut = &timeOut
	}
	server.enginez.timeout = server.plugin.ReqTimeOut

	// 启动之前，检查下是否注入了 Logger 和 Recovery
	if server.plugin.Logger == nil {
		server.plugin.Logger = loggerMiddleware
		server.use(server.plugin.Logger)
	}
	if server.plugin.Recovery == nil {
		server.plugin.Recovery = recoverMiddleware
		server.use(server.plugin.Recovery)
	}
	if server.plugin.Metrics == nil {
		server.plugin.Metrics = httpMetricsMiddleware
		server.use(server.plugin.Metrics)
	}
	if server.plugin.Authored == nil {
		if server.plugin.JwtAk == nil {
			server.plugin.JwtAk = func() (s2 string, exp int64, min int64) {
				return "hello", 7 * 24 * 60 * 60, 3 * 24 * 60 * 60
			}
		}
		server.jwtUtils = core.NewJwtUtils(server.plugin.JwtAk())
		server.enginez.jwtUtils = server.jwtUtils
		server.plugin.Authored = authorMiddleware(server)
		// 默认不适用，判断当 route 需要鉴权时再加入
		// server.use(server.plugin.Authored)
	}
	return server
}

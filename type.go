package rest

import (
	"context"
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
		AllowCORS          bool
		TimeOut            *time.Duration
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
		ResponseBodyJsonStr() string
		// @Authored handler 的返回值才会有意义
		UserID() int64
		// 如果内部调用了此方法，会尝试生成 or 刷新一个 jwt 给客户端
		SetUserID(int64)
		// 设置响应头
		SetResponseHeader(key, value string)

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

		// private
		setTimeout(duration *time.Duration)
		setUserID(int64)
		setRequestID(string)
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

type Plugin struct {
	Logger       HandlerFunc
	Recovery     HandlerFunc
	RequestIdGen HandlerFunc
	Cors         HandlerFunc                   // 跨域
	Authored     func(HandlerFunc) HandlerFunc // 默认实现为 jwt
	Metrics      HandlerFunc                   // prometheus
	MetricsPort  int

	// 用于设置 jwt ak 过期时间
	JwtAk func() (string, int64, int64)
}

// 初始化一个 context
func BackgroundContext() Context {
	c := newContext(nil)
	return c
}

// NewServer 创建一个服务实例
func NewServer(port int, opts ...Option) Server {
	return newServerImp(port, opts...)
}

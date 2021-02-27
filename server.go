package rest

import (
	"fmt"
	"github.com/gozelus/zelus_rest/core"
	"github.com/gozelus/zelus_rest/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

type eng interface {
	use(middlrewares ...HandlerFunc)
	addRoute(method, path string, timeout time.Duration, f HandlerFunc) error
	setJwtUtils(*core.JwtUtils)
	run(port int) error
}
type serverImp struct {
	eng      eng
	port     int
	plugin   *Plugin
	jwtUtils *core.JwtUtils
}

type Option = func(imp *Plugin)

func newServerImp(port int, opts ...Option) *serverImp {
	server := &serverImp{
		port: port,
	}
	p := &Plugin{
		Logger:   nil,
		Recovery: nil,
	}
	for _, o := range opts {
		o(p)
	}
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

	if server.plugin.RequestIdGen == nil {
		server.plugin.RequestIdGen = requestIdGenMiddleware
		server.eng.use(server.plugin.RequestIdGen)
	}
	if server.plugin.Logger == nil {
		server.plugin.Logger = loggerMiddleware
		server.eng.use(server.plugin.Logger)
	}
	if server.plugin.Recovery == nil {
		server.plugin.Recovery = recoverMiddleware
		server.eng.use(server.plugin.Recovery)
	}
	if server.plugin.Cors == nil {
		server.plugin.Cors = corsMiddleware
		server.eng.use(server.plugin.Cors)
	}
	if server.plugin.Metrics == nil {
		server.plugin.Metrics = httpMetricsMiddleware
		server.eng.use(server.plugin.Metrics)
	}
	if server.plugin.Authored == nil {
		if server.plugin.JwtAk == nil {
			server.plugin.JwtAk = func() (s2 string, exp int64, min int64) {
				return "hello", 7 * 24 * 60 * 60, 3 * 24 * 60 * 60
			}
		}
		server.jwtUtils = core.NewJwtUtils(server.plugin.JwtAk())
		server.eng.setJwtUtils(server.jwtUtils)
		server.plugin.Authored = authorMiddleware(server)
		// 默认不适用，判断当 route 需要鉴权时再加入
		// server.use(server.plugin.Authored)
	}
	return server
}

// Use 加载中间件
func (s *serverImp) Use(middlrewares ...HandlerFunc) error {
	s.eng.use(middlrewares...)
	return nil
}

// AddRoute 挂载路由
func (s *serverImp) AddRoute(routes ...Route) error {
	for _, r := range routes {
		if r.NeedAuthentication && s.plugin.Authored != nil {
			if err := s.eng.addRoute(r.Method, r.Path, r.TimeOut, s.plugin.Authored(r.Handler)); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := s.eng.addRoute(r.Method, r.Path, r.TimeOut, r.Handler); err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}

// Run
func (s *serverImp) Run() error {
	return s.eng.run(s.port)
}

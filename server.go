package rest

import (
	"context"
	"github.com/gozelus/zelus_rest/core"
	"log"
	"net/http"
)

type serverImp struct {
	engine     *core.Enginez
	httpServer *http.Server
}

// Use 加载中间件
func (s *serverImp) Use(middlrewares ...Middleware) error {
	for _, m := range middlrewares {
		s.engine.Use(func(next http.HandlerFunc) http.HandlerFunc {
			return m(next)
		})
	}
	return nil
}

// AddRoute 挂载路由
func (s *serverImp) AddRoute(routes ...Route) error {
	for _, r := range routes {
		if err := s.engine.AddRoute(r.Method, r.Path, r.Handler); err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

// Run
func (s *serverImp) Run() error {
	server := s.httpServer
	defer server.Shutdown(context.Background())
	return server.ListenAndServe()
}

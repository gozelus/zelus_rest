package rest

import (
	"context"
	"log"
	"net/http"
)

type serverImp struct {
	*enginez
	httpServer *http.Server
	plugin     *Plugin
}

// Use 加载中间件
func (s *serverImp) Use(middlrewares ...HandlerFunc) error {
	for _, m := range middlrewares {
		s.use(m)
	}
	return nil
}

// AddRoute 挂载路由
func (s *serverImp) AddRoute(routes ...Route) error {
	for _, r := range routes {
		if err := s.addRoute(r.Method, r.Path, r.Handler); err != nil {
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

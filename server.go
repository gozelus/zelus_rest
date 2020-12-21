package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	engine engine
}

type engine interface {
	addRoute(method, path string, handler http.HandlerFunc) error
	use(middlewares ...Middleware)

	ServeHTTP(http.ResponseWriter, *http.Request)
}

func NewServer() *Server {
	ege := newEnginez(NewRouterz())
	return newServer(ege)
}

func newServer(e engine) *Server {
	return &Server{
		engine: e,
	}
}

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func (s *Server) Use(middlrewares ...Middleware) *Server {
	s.engine.use(middlrewares...)
	return s
}
func (s *Server) AddRoute(routes ...Route) *Server {
	for _, r := range routes {
		if err := s.engine.addRoute(r.Method, r.Path, r.Handler); err != nil {
			log.Fatal(err)
		}
	}
	return s
}
func (s *Server) Run(host string, port int) error {
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", host, port),
	}
	server.Handler = s.engine
	defer server.Shutdown(context.Background())
	return server.ListenAndServe()
}

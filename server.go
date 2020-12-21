package rest

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	engine engine
}

type engine interface {
	addRoute(method, path string, handler http.HandlerFunc) error
	http.Handler
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

func (s *Server) Run(host string, port int, routes ...Route) error {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: s.engine,
	}

	for _, r := range routes {
		s.engine.addRoute(r.Method, r.Path, r.Handler)
	}

	defer server.Shutdown(context.Background())
	return server.ListenAndServe()
}

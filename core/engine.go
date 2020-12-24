package core

import (
	"net/http"
	"path"
)

type Enginez struct {
	middlewares []func(handlerFunc http.HandlerFunc) http.HandlerFunc
	router      *routerz
}

func NewEnginez() *Enginez {
	e := &Enginez{
		router: NewRouterz(),
	}
	return e
}
func (e *Enginez) AddRoute(method, path string, handler http.HandlerFunc) error {
	return e.router.AddRoute(method, path, handler)
}
func (e *Enginez) Search(method, path string) (http.HandlerFunc, error) {
	return e.router.Search(method, path)
}
func (e *Enginez) Use(middlewares ...func(next http.HandlerFunc) http.HandlerFunc) {
	for _, m := range middlewares {
		e.middlewares = append(e.middlewares, m)
	}
}
func (e *Enginez) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqPath := path.Clean(req.URL.Path)
	h, _ := e.router.Search(req.Method, reqPath)
	for _, middleware := range e.middlewares {
		h = middleware(h)
	}
	h(w, req)
}

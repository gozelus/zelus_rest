package rest

import (
	"net/http"
	"path"
)

type router interface {
	addRoute(method, path string, handler http.HandlerFunc) error
	search(method, path string) (http.HandlerFunc, error)
}

type enginez struct {
	middlewares []Middleware
	router
}

func newEnginez(r router) *enginez {
	e := &enginez{
		router: r,
	}
	return e
}

func (e *enginez) use(middlewares ...Middleware) {
	for _, m := range  middlewares {
		e.middlewares = append(e.middlewares, m)
	}
}


func (e *enginez) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqPath := path.Clean(req.URL.Path)

	h, _ := e.search(req.Method, reqPath)
	for _, middleware := range e.middlewares {
		h = middleware(h)
	}
	h(w, req)
}

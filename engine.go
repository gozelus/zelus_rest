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
	router
}

func newEnginez(r router) *enginez {
	e :=  &enginez{
		router: r,
	}
	return e
}


func (e *enginez) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqPath := path.Clean(req.URL.Path)
	h , _ := e.search(req.Method, reqPath)
	h(w, req)
}

package core

import (
	"net/http"
)

type routerz struct {
	notFoundHandler         http.HandlerFunc
	methodNotAllowedHandler http.HandlerFunc
	trees                   map[string]*Tree
}

func NewRouterz(hooks ...http.HandlerFunc) *routerz {
	r := &routerz{
		trees: make(map[string]*Tree),
	}
	if len(hooks) > 0 {
		r.notFoundHandler = hooks[0]
		if len(hooks) > 1 {
			r.methodNotAllowedHandler = hooks[1]
		}
	}
	return r
}

func (r *routerz) AddRoute(method, path string, handler http.HandlerFunc) error {
	var tree *Tree
	if t, ok := r.trees[method]; ok {
		tree = t
	} else {
		tree = NewTree()
		r.trees[method] = tree
	}
	return tree.Add(path, handler)
}

func (r *routerz) Search(method, path string) (http.HandlerFunc, error) {
	if tree, ok := r.trees[method]; ok {
		if result, ok := tree.Search(path); ok {
			return result.Item.(http.HandlerFunc), nil
		} else {
			if r.notFoundHandler != nil {
				return r.notFoundHandler, nil
			}
			return http.NotFoundHandler().ServeHTTP, nil
		}
	}
	if r.methodNotAllowedHandler != nil {
		return r.methodNotAllowedHandler, nil
	}
	return http.NotFoundHandler().ServeHTTP, nil
}

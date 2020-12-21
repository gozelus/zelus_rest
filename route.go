package rest

import (
	"net/http"

	searchTree "github.com/gozelus/zelus_rest/core/search_tree"
)

type routerz struct {
	notFoundHandler         http.HandlerFunc
	methodNotAllowedHandler http.HandlerFunc
	trees                   map[string]*searchTree.Tree
}

func NewRouterz(hooks ...http.HandlerFunc) *routerz {
	r := &routerz{
		trees: make(map[string]*searchTree.Tree),
	}
	if len(hooks) > 0 {
		r.notFoundHandler = hooks[0]
		if len(hooks) > 1 {
			r.methodNotAllowedHandler = hooks[1]
		}
	}
	return r
}

func (r *routerz) addRoute(method, path string, handler http.HandlerFunc) error {
	var tree *searchTree.Tree
	if t, ok := r.trees[method]; ok {
		tree = t
	} else {
		tree = searchTree.NewTree()
		r.trees[method] = tree
	}
	return tree.Add(path, handler)
}

func (r *routerz) search(method, path string) (http.HandlerFunc, error) {
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

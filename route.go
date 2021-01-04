package rest

import (
	"github.com/gozelus/zelus_rest/core"
)

type routerz struct {
	trees                   map[string]*core.Tree
}

func newRouterz() *routerz {
	r := &routerz{
		trees: make(map[string]*core.Tree),
	}
	return r
}

func (r *routerz) addRoute(method, path string, handler HandlerFunc) error {
	var tree *core.Tree
	if t, ok := r.trees[method]; ok {
		tree = t
	} else {
		tree = core.NewTree()
		r.trees[method] = tree
	}
	return tree.Add(path, handler)
}

func (r *routerz) search(method, path string) (HandlerFunc, error) {
	if tree, ok := r.trees[method]; ok {
		if result, ok := tree.Search(path); ok {
			return result.Item.(HandlerFunc), nil
		}
	}
	return func(context Context) {
		context.RenderErrorJSON(nil, statusNotFound)
	}, nil
}

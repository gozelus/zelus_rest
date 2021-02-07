package rest

import (
	"net/http"
	"path"
	"sync"
	"time"
)

type enginez struct {
	// routerz 实现路由搜索添加的功能
	*routerz

	// 挂载在引擎上的全局中间件
	middlewares []HandlerFunc

	// 用于取用 Context 实例
	pool sync.Pool

	// jwt
	jwtUtils jwtUtils

	// timeout default 1000 ms
	timeout *time.Duration
}

func newEnginez() *enginez {
	e := &enginez{
		routerz: newRouterz(),
	}
	e.pool.New = func() interface{} {
		return e.allocateContext()
	}
	return e
}

func (e *enginez) addRoute(method, path string, handler HandlerFunc) error {
	return e.routerz.addRoute(method, path, handler)
}

func (e *enginez) search(method, path string) (HandlerFunc, error) {
	return e.routerz.search(method, path)
}

func (e *enginez) use(middlewares ...HandlerFunc) {
	for _, m := range middlewares {
		e.middlewares = append(e.middlewares, m)
	}
}

func (e *enginez) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqPath := path.Clean(req.URL.Path)
	h, _ := e.search(req.Method, reqPath)
	// TODO need check method and path is ok
	// means u should check search func called err result

	// core code
	// 1. get ctx from pool
	ctx := e.pool.Get().(Context)
	// 2. reset the ctx
	ctx.init(w, req, e.timeout)
	ctx.setJwtUtils(e.jwtUtils)
	ctx.setHandlers(append(e.middlewares, h)...)
	// 3. pass ctx to handler and run it
	ctx.Next()
	e.pool.Put(ctx)
}

func (e *enginez) allocateContext() Context {
	return newContext()
}

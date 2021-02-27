package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gozelus/zelus_rest/core"
	"net/http"
	"time"
)

type engz struct {
	jwtUtils *core.JwtUtils
	ginEng   *gin.Engine
}

func newEngz() *engz {
	ginEng := gin.New()
	gin.SetMode(gin.ReleaseMode)
	return &engz{
		ginEng: ginEng,
	}
}

func (e *engz) use(middlrewares ...HandlerFunc) {
	for _, m := range middlrewares {
		e.ginEng.Use(func(context *gin.Context) {
			m(newContext(context))
		})
	}
}

func (e *engz) addRoute(method, path string, timeout *time.Duration, f HandlerFunc) error {
	var wrap = func(ctx *gin.Context) {
		c := newContext(ctx)
		c.setJwtUtils(e.jwtUtils)
		c.setTimeout(timeout)
		f(c)
	}
	switch method {
	case http.MethodGet:
		e.ginEng.GET(path, wrap)
	case http.MethodPost:
		e.ginEng.POST(path, wrap)
	case http.MethodOptions:
		e.ginEng.OPTIONS(path, wrap)
	case http.MethodDelete:
		e.ginEng.DELETE(path, wrap)
	case http.MethodPut:
		e.ginEng.PUT(path, wrap)
	case http.MethodHead:
		e.ginEng.HEAD(path, wrap)
	default:
		return errors.New("invalid method : %s", method)
	}
	return nil
}

func (e *engz) setJwtUtils(jwt *core.JwtUtils) {
	e.jwtUtils = jwt
}

func (e *engz) run(port int) error {
	return e.ginEng.Run(fmt.Sprintf("127.0.0.1:%d", port))
}

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
	return &engz{
		ginEng: gin.New(),
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
		f(newContext(ctx))
	}
	switch method {
	case http.MethodGet:
		e.ginEng.GET(path, wrap)
	case http.MethodPost:
	case http.MethodOptions:
	case http.MethodDelete:
	case http.MethodPut:
	case http.MethodHead:
	}
}

func (e *engz) setJwtUtils(jwt *core.JwtUtils) {
	e.jwtUtils = jwt
}

func (e *engz) run(port int) error {
	return e.ginEng.Run(fmt.Sprintf("127.0.0.1:%d", port))
}

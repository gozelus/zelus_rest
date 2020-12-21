package rest

import (
	"net/http"
	path2 "path"
	"time"

	"github.com/gozelus/zelus_rest/logger"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc

var (
	LoggerMiddleware Middleware = func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			method := request.Method
			path := path2.Clean(request.URL.Path)
			now := time.Now()
			next(writer, request)
			logger.Infof("[%s]-[%s]-[%dms]", method, path, now.Sub(time.Now()).Milliseconds())
		}
	}
)

package rest

import (
	"github.com/gozelus/zelus_rest/core/metric"
	"github.com/gozelus/zelus_rest/logger"
	"github.com/hashicorp/go-uuid"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

const serverNamespace = "http_server"

var (
	metricServerReqDur = metric.NewHistogramVec(&metric.HistogramVecOpts{
		Namespace: serverNamespace,
		Subsystem: "requests",
		Name:      "duration_ms",
		Help:      "http server requests duration(ms).",
		Labels:    []string{"path"},
		Buckets:   []float64{5, 10, 25, 50, 100, 250, 500, 1000},
	})

	metricServerReqCodeTotal = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: serverNamespace,
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "http server requests error count.",
		Labels:    []string{"path", "code"},
	})
)
var httpMetricsMiddleware = func(c Context) {
	startTime := time.Now()
	c.Next()
	duration := time.Now().Sub(startTime).Milliseconds()
	metricServerReqDur.Observe(int64(duration), c.Path())
	metricServerReqCodeTotal.Inc(c.Path(), strconv.Itoa(c.HttpCode()))
}

var requestIdGenMiddleware = func(c Context) {
	requestID, err := uuid.GenerateUUID()
	if err != nil {
		c.RenderErrorJSON(nil, err)
	}
	c.setRequestID(requestID)
}
var authorMiddleware = func(server *serverImp) func(HandlerFunc) HandlerFunc {
	return func(handlerFunc HandlerFunc) HandlerFunc {
		return func(ctx Context) {
			if token, ok := ctx.Headers()["Authorization"]; ok && len(token) > 0 && len(token[0]) > 0 {
				userID, newTokenStr, err := server.jwtUtils.ValidateToken(token[0])
				if err != nil {
					ctx.RenderErrorJSON(nil, statusUnauthorized)
					return
				}
				ctx.setUserID(userID)
				ctx.setJwtToken(newTokenStr)
				handlerFunc(ctx)
				return
			}
			ctx.RenderErrorJSON(nil, statusUnauthorized)
			return
		}
	}
}
var recoverMiddleware = func(c Context) {
	defer func() {
		if err := recover(); err != nil {
			logger.ErrorfWithStackWithContext(c, "recover err : %s", err)
			c.RenderErrorJSON(nil, statusInternalServerError)
		}
	}()
	c.Next()
}
var loggerMiddleware = func(c Context) {
	now := time.Now()
	logger.InfofWithContext(c, "query map : %v | json body : %s", c.QueryMap(), c.RequestBodyJsonStr())
	c.Next()
	logger.InfofWithContext(c, "response body : %s", c.ResponseBodyJsonStr())
	if err := c.GetError(); err != nil {
		logger.WarnfWithContext(c, "bad request -> [%d] | method : %s | path : %s | duration : %d ms | err : %T -> %+v", c.HttpCode(), c.Method(), c.Path(), now.Sub(now).Milliseconds(), errors.Cause(err), err)
	} else {
		duration := time.Now().Sub(now).Milliseconds()
		if duration > 300 {
			logger.WarnfWithContext(c, "slow request -> [%d] | method : %s | path : %s | duration : %d ms", c.HttpCode(), c.Method(), c.Path(), time.Now().Sub(now).Milliseconds())
		} else {
			logger.InfofWithContext(c, "ok request -> [%d] | method : %s | path : %s | duration : %d ms", c.HttpCode(), c.Method(), c.Path(), time.Now().Sub(now).Milliseconds())
		}
	}
}

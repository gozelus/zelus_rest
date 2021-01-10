package rest

import (
	"github.com/gozelus/zelus_rest/logger"
	"github.com/pkg/errors"
	"time"
)

var authorMiddleware = func(server *serverImp) func(HandlerFunc) HandlerFunc {
	return func(handlerFunc HandlerFunc) HandlerFunc {
		return func(ctx Context) {
			ctx.setJwtUtils(server.jwtUtils)
			if token, ok := ctx.Headers()["Authorization"]; ok && len(token) > 0 && len(token[0]) > 0 {
				userID, newTokenStr, err := server.jwtUtils.ValidateToken(token[0])
				if err != nil {
					ctx.RenderErrorJSON(nil, statusUnauthorized)
					return
				}
				ctx.setUserID(userID)
				handlerFunc(ctx)
				ctx.Set("jwt-token", newTokenStr)
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
	c.Next()
	if err := c.GetError(); err != nil {
		logger.WarnfWithContext(c, "method : %s | path : %s | duration : %d ms | err : %T -> %+v", c.Method(), c.Path(), now.Sub(now).Milliseconds(), errors.Cause(err), err)
	} else {
		duration := time.Now().Sub(now).Milliseconds()
		if duration > 300 {
			logger.WarnfWithContext(c, "slow request -> method : %s | path : %s | duration : %d ms", c.Method(), c.Path(), time.Now().Sub(now).Milliseconds())
		} else {
			logger.InfofWithContext(c, "ok request -> method : %s | path : %s | duration : %d ms", c.Method(), c.Path(), time.Now().Sub(now).Milliseconds())
		}
	}
}

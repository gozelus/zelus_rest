package router

import (
	"github.com/gozelus/zelus_rest"
	"net/http"
)

type UserControllerInter interface {
	GetUser(w http.ResponseWriter, req *http.Request)
}
type Router struct {
	user UserControllerInter
}

func NewRouter(user UserControllerInter) *Router {
	return &Router{user: user}
}

func (r *Router) Routes() []rest.Route {
	return []rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/user/get",
			Handler: r.user.GetUser,
		},
	}
}

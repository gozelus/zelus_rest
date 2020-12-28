package router

import (
	"github.com/gozelus/zelus_rest"
	"net/http"
)

type UserControllerInter interface {
	Register(ctx rest.Context)
}
type Router struct {
	user UserControllerInter
}

func NewRouter(user UserControllerInter) *Router {
	return &Router{user: user}
}

var UserGroup = func(r *Router) []rest.Route {
	return []rest.Route{
		//{
		//	Method:  http.MethodGet,
		//	Path:    "/user/get",
		//	Handler: r.user.GetUser,
		//},
		{
			Method:  http.MethodPost,
			Path:    "/user/register",
			Handler: r.user.Register,
		},
	}
}

func (r *Router) Routes() []rest.Route {
	return UserGroup
}

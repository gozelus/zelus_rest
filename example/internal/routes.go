package internal

import (
	"github.com/gozelus/zelus_rest"
	"net/http"
)

type userController interface {
	GetUser(w http.ResponseWriter, req *http.Request)
}
type Router struct {
	user userController
}

func NewRouter(user userController) *Router {
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

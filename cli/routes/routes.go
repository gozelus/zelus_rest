package router

import (
	"github.com/gozelus/zelus_rest"
	"net/http"
)


type UserControllerInter interface {
	 UserGet(w http.ResponseWriter, req *http.Request)
	 UserCreate(w http.ResponseWriter, req *http.Request)
	
}


type Router struct {
	 User UserControllerInter 
}

func NewRouter(
User UserControllerInter,) *Router {
	return &Router{ 
		User: User, 
	}
}

func (r *Router) Routes() []rest.Route {
	return []rest.Route{
		
		
		{
			Method:  "GET",
			Path:    "/user/info",
			Handler: r.User.UserGet,
		},
        
		{
			Method:  "POST",
			Path:    "/user/create",
			Handler: r.User.UserCreate,
		},
        
		
	}
}

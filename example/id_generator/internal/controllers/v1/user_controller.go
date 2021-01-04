package v1

import (
	"github.com/gozelus/zelus_rest"
)

type UserService interface { 
	UserInfo(ctx rest.Context, req *api.User) (*api.User, error) 
	UserCreate(ctx rest.Context, req *api.User) (*api.User, error) 
}
type UserController struct {
	service UserService
}
func NewUserController(service UserService) *UserController {
	return &UserController{service : service}
}


func (c *UserController) UserInfo(ctx rest.Context) {
	res := &api.User{}
	req := &api.User{}
	var err error 
	if err := ctx.JSONBodyBind(req); err != nil {
		ctx.RenderErrorJSON(nil, apiErrors.BadRequest.WithReason(err.Error()))
		return
	}
	if res, err = c.service.UserInfo(ctx, req);err!=nil{
		ctx.RenderErrorJSON(nil, err)
		return
	}
	ctx.RenderOkJSON(res)
}

func (c *UserController) UserCreate(ctx rest.Context) {
	res := &api.User{}
	req := &api.User{}
	var err error 
	if err := ctx.JSONQueryBind(req); err != nil {
		ctx.RenderErrorJSON(nil, apiErrors.BadRequest.WithReason(err.Error()))
		return
	}
	if res, err = c.service.UserCreate(ctx, req);err!=nil{
		ctx.RenderErrorJSON(nil, err)
		return
	}
	ctx.RenderOkJSON(res)
}


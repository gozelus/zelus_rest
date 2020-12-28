package user

import (
	rest "github.com/gozelus/zelus_rest"
)

type Controller struct {
	user UserServiceInterface
}

type UserServiceInterface interface {
	Register(ctx rest.Context, req *RegisterRequest) error
	Info(ctx rest.Context, req *InfoRequest) error
}

func NewController(user UserServiceInterface) *Controller {
	return &Controller{
		user: user,
	}
}
func (c *Controller) Register(ctx rest.Context) {
	req := RegisterRequest{}
	if err := ctx.JSONBodyBind(&req); err != nil {
		ctx.RenderJSON(rest.StatusBadRequestResp())
	}
	//if err := c.user.Register(ctx, &req); err != nil {
	//	return rest.StatusInternalServerError
	//}
}
func (c *Controller) Info(ctx rest.Context) {
}

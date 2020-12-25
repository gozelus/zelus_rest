package user

import (
	rest "github.com/gozelus/zelus_rest"
)

type Controller struct {
	user userServiceInterface
}

type userServiceInterface interface {
	Register(ctx rest.Context, req *RegisterRequest) error
	Info(ctx rest.Context, req *InfoRequest) error
}

func NewController(user userServiceInterface) *Controller {
	return &Controller{
		user: user,
	}
}
func (c *Controller) Register(ctx rest.Context) rest.ErrorInterface {
	req := RegisterRequest{}
	if err := ctx.JSONBodyBind(ctx); err != nil {
		return rest.StatusBadRequest
	}
	if err := c.user.Register(ctx, &req); err != nil {
		return rest.StatusInternalServerError
	}
	return nil
}
func (c *Controller) Info(ctx rest.Context) rest.ErrorInterface {
	return nil
}

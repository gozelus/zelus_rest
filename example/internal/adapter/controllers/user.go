package controllers

import (
	rest "github.com/gozelus/zelus_rest"
)

type RegisterRequest struct {
	Age      int    `json:"age" validate:"gte=18 lte=99 required"`
	Nickname string `json:"nickname" validate:"required"`
	Avatar   string `json:"avatar"`
}
type RegisterResponse struct {
	User interface{} `json:"user"`
}
type InfoRequest struct {
	UserID int64 `json:"user_id",range:"[1,2]"`
}
type InfoResponse struct {
	User interface{} `json:"user"`
}

type Controller struct {
	user UserDomain
}

type UserDomain interface {
	Register(ctx rest.Context, req *RegisterRequest) error
	Info(ctx rest.Context, req *InfoRequest) error
}

func NewController(user UserDomain) *Controller {
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

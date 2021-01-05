package v1_controllers

import (
	rest "github.com/gozelus/zelus_rest"

	api "github.com/gozelus/zelus_rest/internal"

	v1_services "github.com/gozelus/zelus_rest/internal/services/v1"
)

type UserController struct {
	service *v1_services.UserService
}

func NewUserController(service *v1_services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) LoginByPhoneCode(ctx rest.Context) {
	res := &api.LoginByPhoneCodeResponse{}
	req := &api.LoginByPhoneCodeRequest{}
	var err error
	if err := ctx.JSONBodyBind(req); err != nil {
		ctx.RenderErrorJSON(nil, apiErrors.BadRequest.WithReason(err.Error()))
		return
	}
	if res, err = c.service.LoginByPhoneCode(ctx, req); err != nil {
		ctx.RenderErrorJSON(nil, err)
		return
	}
	ctx.RenderOkJSON(res)
}

func (c *UserController) SendLoginOrRegisterPhoneCode(ctx rest.Context) {
	res := &api.SendLoginOrRegisterResponse{}
	req := &api.SendLoginOrRegisterRequest{}
	var err error
	if err := ctx.JSONBodyBind(req); err != nil {
		ctx.RenderErrorJSON(nil, apiErrors.BadRequest.WithReason(err.Error()))
		return
	}
	if res, err = c.service.SendLoginOrRegisterPhoneCode(ctx, req); err != nil {
		ctx.RenderErrorJSON(nil, err)
		return
	}
	ctx.RenderOkJSON(res)
}

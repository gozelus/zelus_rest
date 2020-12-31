package controllers

import "github.com/gozelus/zelus_rest"

type UserService interface {
	RegisterOrLogin(ctx rest.Context, phone, code string) error
	SendPhoneCode(ctx rest.Context, phone string) error
}
type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}
func (c *UserController) RegisterOrLogin(ctx rest.Context) {
}
func (c *UserController) SendPhoneCode(ctx rest.Context) {
}

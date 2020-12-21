package controller

import (
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/internal/types"
	"net/http"
)

type UserServiceInter interface {
	GetUser(req *types.GetUserRequest) (*types.GetUserResponse, error)
}
type UserController struct {
	userService UserServiceInter
}

func NewUserController(userService UserServiceInter) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) GetUser(w http.ResponseWriter, req *http.Request) {
	getUserRequest := types.GetUserRequest{}
	// todo check err
	rest.JsonBodyFromRequest(req, &getUserRequest)
	// todo check err
	res, _ := c.userService.GetUser(&getUserRequest)
	rest.OkJson(w, res)
}

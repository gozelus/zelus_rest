package controller

import (
	"net/http"

	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/cli/types"
)

type UserServiceInterface interface {
	UserGet(*types.UserGetRequest) (*types.UserGetResponse, error)
	UserCreate(*types.UserCreateRequest) (*types.UserCreateResponse, error)
}

type User struct {
	UserService UserServiceInterface
}

func NewUser(UserService UserServiceInterface) *User {
	return &User{
		UserService: UserService,
	}
}

func (c *User) UserGet(w http.ResponseWriter, req *http.Request) {
	param := &types.UserGetRequest{}
	rest.JsonBodyFromRequest(req, param)
	res, _ := c.UserService.UserGet(param)
	rest.OkJson(w, res)
}

func (c *User) UserCreate(w http.ResponseWriter, req *http.Request) {
	param := &types.UserCreateRequest{}
	rest.JsonBodyFromRequest(req, param)
	res, _ := c.UserService.UserCreate(param)
	rest.OkJson(w, res)
}

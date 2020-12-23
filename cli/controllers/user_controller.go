package controller

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"

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

func (c *User) checkReq(request types.UserGetRequest) {
	t := reflect.TypeOf(request)
	for i := 0; i < t.NumField(); i++ {
		filed := t.Field(i)
		tag := filed.Tag
		v := reflect.ValueOf(request).Field(i)
		c.handleTag(request, i, tag, filed.Type, v)
	}
}

func (c *User) handleTag(originValue interface{}, fieldIdx int, tag reflect.StructTag, fieldType reflect.Type, value reflect.Value) error {
	if defaultValue, ok := tag.Lookup("default"); ok && value.IsZero() {
		if value.Kind() == reflect.Int {
			reflect.
		}
	}
	return nil
}

func handleDefaultTag(defaultValue string, typeName string) (value interface{}, err error) {
	switch typeName {
	case "string":
		return defaultValue, nil
	case "int64":
		return strconv.ParseInt(defaultValue, 10, 64)
	case "int":
		return strconv.ParseInt(defaultValue, 10, 64)
	case "int8":
		return strconv.ParseInt(defaultValue, 10, 8)
	case "int32":
		return strconv.ParseInt(defaultValue, 10, 32)
	case "float32":
		return strconv.ParseFloat(defaultValue, 32)
	case "float64":
		return strconv.ParseFloat(defaultValue, 64)
	default:
		return nil, errors.New("")
	}
}

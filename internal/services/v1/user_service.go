package v1_services

import (
	"errors"

	rest "github.com/gozelus/zelus_rest"
	api "github.com/gozelus/zelus_rest/internal"
)

type UserService struct {
	// 以后放入要依赖度的对象
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) LoginByPhoneCode(ctx rest.Context, request *api.LoginByPhoneCodeRequest) (*api.LoginByPhoneCodeResponse, error) {
	return nil, errors.New("no imp")
}

func (s *UserService) SendLoginOrRegisterPhoneCode(ctx rest.Context, request *api.SendLoginOrRegisterRequest) (*api.SendLoginOrRegisterResponse, error) {
	return nil, errors.New("no imp")
}

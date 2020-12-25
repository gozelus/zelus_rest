package service

import (
	controller "github.com/gozelus/zelus_rest/cli/controllers"
	"github.com/gozelus/zelus_rest/cli/types"
	"github.com/gozelus/zelus_rest/example/internal/entity"
)

type UserRepoInter interface {
	MGetUser(userIDs []int64) (map[int64]*entity.User, error)
}

var _ controller.UserServiceInterface = &UserService{}

type UserService struct {
	userRepo UserRepoInter
}

func (s *UserService) UserGet(request *types.UserGetRequest) (*types.UserGetResponse, error) {
	panic("implement me")
}

func (s *UserService) UserCreate(request *types.UserCreateRequest) (*types.UserCreateResponse, error) {
	panic("implement me")
}

func NewUserService(userRepo UserRepoInter) *UserService {
	return &UserService{userRepo: userRepo}
}

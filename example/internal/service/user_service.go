package service

import (
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/internal/controller/user"
	"github.com/gozelus/zelus_rest/example/internal/entity"
)

type UserRepoInter interface {
	MGetUser(context *rest.Context, userIDs []int64) (map[int64]*entity.User, error)
	Create(context *rest.Context, user *entity.User) (int64, error)
	Update(context *rest.Context, userID int64, attr map[string]interface{}) error
}

var _ user.UserServiceInterface = &UserService{}

type UserService struct {
	userRepo UserRepoInter
}

func (UserService) Register(ctx rest.Context, req *user.RegisterRequest) error {
	panic("implement me")
}

func (UserService) Info(ctx rest.Context, req *user.InfoRequest) error {
	panic("implement me")
}

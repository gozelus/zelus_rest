package service

import (
	"github.com/gozelus/zelus_rest/example/entity"
	"github.com/gozelus/zelus_rest/example/internal/types"
)

type userRepo interface {
	MGetUser(userIDs []int64) (map[int64]*entity.User, error)
}

type UserService struct {
	userRepo userRepo
}

func NewUserService(userRepo userRepo) *UserService {
	return &UserService{userRepo: userRepo}
}
func (s *UserService) GetUser(req *types.GetUserRequest) (*types.GetUserResponse, error) {
	users, err := s.userRepo.MGetUser([]int64{req.UserID})

	if err != nil {
		return nil, err
	}
	res := &types.GetUserResponse{}

	if val, ok := users[req.UserID]; ok {
		res.UserID = val.UserID
		res.UserName = val.NickName
		return res, nil
	}
	return nil, err
}

package user

import (
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/internal/controllers"
)

type UserBindRepo interface {
	First(ctx rest.Context, bindValue string, bindType uint8) (userID int64, err error)
	FirstOrCreate(ctx rest.Context, bindUserID int64, bindValue string, bindType uint8) (userID int64, err error)
}
type UserRepo interface {
}

type UserService struct {
	bindRepo UserBindRepo
}

var _ controllers.UserService = &UserService{}

func (u *UserService) RegisterOrLogin(ctx rest.Context, phone, code string) error {
	if code != "123456" {
	}
	userID, err := u.bindRepo.First(ctx, phone, 1)
	if err != nil {
		return err
	}
	panic("implement me")
}

func (UserService) SendPhoneCode(ctx rest.Context, phone string) error {
	return nil
}

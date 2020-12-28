package entity

import (
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/internal/domain/user/po"
)

type UserEntity struct {
	po.UserInfoPo
}

func (u *UserEntity) SetNickname(ctx rest.Context, newName string) error {
	u.Nickname = newName
	return nil
}
func (u *UserEntity) SetAge(ctx rest.Context, age int) error {
	u.Age = age
	return nil
}

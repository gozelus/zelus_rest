package model

import "github.com/gozelus/zelus_rest/example/internal/biz/repos"

type User struct {
	ID         int64
	NickName   string
	AvatarGuid string
	Phone      string

	bindRepo     *repos.UserBindsModelRepoImp
	userInfoRepo *repos.UsersModelRepoImp
}

func NewUser(bindRepo *repos.UserBindsModelRepoImp, userInfoRepo *repos.UsersModelRepoImp) *User {
	return &User{
		bindRepo:     bindRepo,
		userInfoRepo: userInfoRepo,
	}
}

func (u *User) Save() error {
	if u.ID == 0 {
	}
}

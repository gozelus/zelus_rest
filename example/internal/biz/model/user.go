package model

import (
	"errors"
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/core/db"
	"github.com/gozelus/zelus_rest/example/internal/biz/repos"
	"github.com/gozelus/zelus_rest/example/internal/data/db"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         int64
	NickName   string
	AvatarGuid string
	Phone      string

	db           db.MySQLDb
	bindRepo     *repos.UserBindsModelRepoImp
	userInfoRepo *repos.UsersModelRepoImp
}

func NewUser(db db.MySQLDb, bindRepo *repos.UserBindsModelRepoImp, userInfoRepo *repos.UsersModelRepoImp) *User {
	return &User{
		db:           db,
		bindRepo:     bindRepo,
		userInfoRepo: userInfoRepo,
	}
}

func (u *User) Update() error {
	return nil
}
func (u *User) Save(ctx rest.Context) error {
	var err error
	if u.ID == 0 && len(u.Phone) > 0 {
		u.ID = time.Now().Unix()
		tx := u.db.Begin()
		bind := &models.UserBindsModel{}
		if bind, err = u.bindRepo.FindOneWithBindCodeBindTypeByTx(ctx, tx, u.Phone, 1); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				tx.Rollback()
				return err
			}
		}
	}
}

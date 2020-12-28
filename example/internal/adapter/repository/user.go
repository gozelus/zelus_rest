package repository

import (
	"database/sql"
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/internal/domain/user"
	"github.com/gozelus/zelus_rest/example/internal/domain/user/entity"
)

type UserRepo struct {
	sql *sql.DB
}

var _ user.Repo = &UserRepo{}

func (UserRepo) New(ctx rest.Context) *entity.UserEntity {
	panic("implement me")
}

func (UserRepo) Save(ctx rest.Context, userEntity *entity.UserEntity) {
	panic("implement me")
}

func NewUserPepo(sql *sql.DB) *UserRepo {
	return &UserRepo{
		sql: sql,
	}
}

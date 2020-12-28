package repo

import (
	"database/sql"
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/example/internal/entity"
	"github.com/gozelus/zelus_rest/example/internal/service"
)

type UserRepo struct {
	sql *sql.DB
}

var _ service.UserRepoInter = &UserRepo{}

func (repo *UserRepo) MGetUser(context *rest.Context, userIDs []int64) (map[int64]*entity.User, error) {
	panic("implement me")
}

func (repo *UserRepo) Create(context *rest.Context, user *entity.User) (int64, error) {
	panic("implement me")
}

func (repo *UserRepo) Update(context *rest.Context, userID int64, attr map[string]interface{}) error {
	panic("implement me")
}

func NewUserRepo(sql *sql.DB) *UserRepo {
	return &UserRepo{sql: sql}
}

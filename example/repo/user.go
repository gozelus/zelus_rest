package repo

import (
	"database/sql"

	"github.com/gozelus/zelus_rest/example/entity"
)

type UserRepo struct {
	sql *sql.DB
}

func NewUserRepo(sql *sql.DB) *UserRepo {
	return &UserRepo{sql: sql}
}

func (repo *UserRepo) MGetUser(userIDs []int64) (map[int64]*entity.User, error) {
	rows, err := repo.sql.Query("select * from users where user_id in (?)", userIDs)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*entity.User)
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user); err != nil {
			return nil, err
		}
	}
	return result, rows.Err()
}

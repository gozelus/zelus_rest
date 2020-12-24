package repo

import (
	"database/sql"

	"github.com/gozelus/zelus_rest/example/internal/entity"
)

type UserRepo struct {
	sql *sql.Conn
}

func NewUserRepo(sql *sql.Conn) *UserRepo {
	return &UserRepo{sql: sql}
}

func (repo *UserRepo) FindOne(userID int64) (*entity.User, error) {
}

func (repo *UserRepo) MGetUser(userIDs []int64) (map[int64]*entity.User, error) {
	rows, err := repo.sql.QueryContext(nil, "select * from users where user_id in (?)", userIDs)
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

package db

import (
	"github.com/gozelus/zelus_rest"
	"gorm.io/gorm"
)

type insertSQL interface {
	Insert(dest interface{}) error
}

var _ insertSQL = &insertSQLImp{}

type insertSQLImp struct {
	db *gorm.DB
}

func (i *insertSQLImp) Insert(dest interface{}) error {
	// 新建一个 Session 用于构建 SQL
	db := i.db.Session(&gorm.Session{DryRun: true})
	stmt := db.Create(dest).Statement
	sql := db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
	ctx := i.db.Statement.Context.(rest.Context)

	return exec(ctx, sql, func() (int64, error) {
		result := i.db.Create(dest)
		return result.RowsAffected, result.Error
	})
}

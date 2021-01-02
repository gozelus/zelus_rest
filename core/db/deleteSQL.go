package db

import (
	"github.com/gozelus/zelus_rest"
	"gorm.io/gorm"
)

type deleteSQL interface {
	Delete(dest interface{}) error
}
type deleteSQLImp struct {
	db *gorm.DB
}

func (d *deleteSQLImp) Delete(dest interface{}) error {
	// 新建一个 Session 用于构建 SQL
	db := d.db.Session(&gorm.Session{DryRun: true})
	stmt := db.Delete(dest).Statement
	sql := db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
	ctx := d.db.Statement.Context.(rest.Context)

	return exec(ctx, sql, func() (int64, error) {
		result := d.db.Delete(dest)
		return result.RowsAffected, result.Error
	})
}

var _ deleteSQL = &deleteSQLImp{}

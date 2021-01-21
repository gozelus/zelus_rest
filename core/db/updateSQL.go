package db

import (
	"gorm.io/gorm"
)

type updateSQL interface {
	Updates(attrs map[string]interface{}) error
}

type updateSQLImp struct {
	db *gorm.DB
}

func (u *updateSQLImp) Updates(attrs map[string]interface{}) error {
	// 新建一个 Session 用于构建 SQL
	db := u.db.Session(&gorm.Session{DryRun: true})
	stmt := db.Updates(attrs).Statement
	sql := db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
	ctx := u.db.Statement.Context

	return exec(ctx, sql, func() (int64, error) {
		result := u.db.Updates(attrs)
		return result.RowsAffected, result.Error
	})

}

var _ updateSQL = &updateSQLImp{}

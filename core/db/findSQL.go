package db

import (
	"gorm.io/gorm"
)

type findSQL interface {
	Find(dest interface{}) error
}

var _ findSQL = &findSQLImp{}

type findSQLImp struct {
	db *gorm.DB
}

func (f *findSQLImp) Find(dest interface{}) error {
	// 新建一个 Session 用于构建 SQL
	db := f.db.Session(&gorm.Session{DryRun: true})
	stmt := db.Find(dest).Statement
	sql := db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
	ctx := f.db.Statement.Context

	return exec(ctx, sql, func() (int64, error) {
		result := f.db.Find(dest)
		return result.RowsAffected, result.Error
	})
}

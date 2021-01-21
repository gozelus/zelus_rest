package db

import (
	"gorm.io/gorm"
)

type firstSQL interface {
	First(dest interface{}) error
}

type firstSQLImp struct {
	db *gorm.DB
}

func (f *firstSQLImp) First(dest interface{}) error {
	// 新建一个 Session 用于构建 SQL
	db := f.db.Session(&gorm.Session{DryRun: true})
	stmt := db.First(dest).Statement
	sql := db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
	ctx := f.db.Statement.Context
	return exec(ctx, sql, func() (i int64, e error) {
		result := f.db.First(dest)
		return result.RowsAffected, result.Error
	})

}

var _ firstSQL = &firstSQLImp{}

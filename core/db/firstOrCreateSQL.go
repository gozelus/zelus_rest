package db

import (
	"github.com/gozelus/zelus_rest"
	"gorm.io/gorm"
)

type firstOrCreateSQL interface {
	FirstOrCreate(dest interface{}) error
}
type firstOrCreateSQLImp struct {
	db *gorm.DB
}

func (f *firstOrCreateSQLImp) FirstOrCreate(dest interface{}) error {
	// 新建一个 Session 用于构建 SQL
	db := f.db.Session(&gorm.Session{DryRun: true})
	stmt := db.FirstOrCreate(dest).Statement
	sql := db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
	ctx := f.db.Statement.Context.(rest.Context)
	return exec(ctx, sql, func() (i int64, e error) {
		result := f.db.FirstOrCreate(dest)
		return result.RowsAffected, result.Error
	})
}

var _ firstOrCreateSQL = &firstOrCreateSQLImp{}

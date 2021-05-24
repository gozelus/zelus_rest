package db

import (
	"context"
	"gorm.io/gorm"
)

type ExecSQL interface {
}

var _ ExecSQL = &execSQLImp{}

type execSQLImp struct {
	db *gorm.DB
}

func (f *execSQLImp) Exec(ctx context.Context, sql string, values ...interface{}) error {
	return f.db.WithContext(ctx).Exec(sql, values...).Error
}

package db

import (
	"context"
	"gorm.io/gorm"
)

type execSQL interface {
}

var _ execSQL = &execSQLImp{}

type execSQLImp struct {
	db *gorm.DB
}

func (f *execSQLImp) Exec(ctx context.Context, sql string, values ...interface{}) error {
	return f.db.WithContext(ctx).Exec(sql, values...).Error
}

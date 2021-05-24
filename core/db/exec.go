package db

import "gorm.io/gorm"

type execSQL interface {
}

var _ execSQL = &execSQLImp{}

type execSQLImp struct {
	db *gorm.DB
}

func (f *execSQLImp) Exec(sql string, values ...interface{}) error {
	return f.db.Exec(sql, values...).Error
}

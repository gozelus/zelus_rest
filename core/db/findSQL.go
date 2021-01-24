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
	return f.db.Find(dest).Error
}

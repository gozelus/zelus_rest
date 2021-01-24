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
	return f.db.First(dest).Error
}

var _ firstSQL = &firstSQLImp{}

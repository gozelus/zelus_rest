package db

import (
	"gorm.io/gorm"
)

type firstOrCreateSQL interface {
	FirstOrCreate(dest interface{}) error
}
type firstOrCreateSQLImp struct {
	db *gorm.DB
}

func (f *firstOrCreateSQLImp) FirstOrCreate(dest interface{}) error {
	return f.db.FirstOrCreate(dest).Error
}

var _ firstOrCreateSQL = &firstOrCreateSQLImp{}

package db

import "gorm.io/gorm"

type insertSQL interface {
	Insert(dest interface{}) error
}

var _ insertSQL = &insertSQLImp{}

type insertSQLImp struct {
	db *gorm.DB
}

func (i *insertSQLImp) Insert(dest interface{}) error {
	return i.db.Create(dest).Error
}

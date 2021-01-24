package db

import (
	"gorm.io/gorm"
)

type deleteSQL interface {
	Delete(dest interface{}) error
}
type deleteSQLImp struct {
	db *gorm.DB
}

func (d *deleteSQLImp) Delete(dest interface{}) error {
	return d.db.Delete(dest).Error
}

var _ deleteSQL = &deleteSQLImp{}

package db

import (
	"gorm.io/gorm"
)

type updateSQL interface {
	Updates(attrs map[string]interface{}) error
}

type updateSQLImp struct {
	db *gorm.DB
}

func (u *updateSQLImp) Updates(attrs map[string]interface{}) error {
	return u.db.Updates(attrs).Error
}

var _ updateSQL = &updateSQLImp{}

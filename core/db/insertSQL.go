package db

import (
	"gorm.io/gorm"
)

type insertSQL interface {
	Insert(dests []interface{}) error
}

var _ insertSQL = &insertSQLImp{}

type insertSQLImp struct {
	db *gorm.DB
}

func (i *insertSQLImp) Insert(dest []interface{}) error {
	// 新建一个 Session 用于构建 SQL
	return i.db.Create(dest).Error
}

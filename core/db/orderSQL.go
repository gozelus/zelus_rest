package db

import "gorm.io/gorm"

type orderSQL interface {
	Order(string) limitSQL
}
type orderSQLImp struct {
	db *gorm.DB
}

func (o *orderSQLImp) Order(orderField string) limitSQL {
	db := o.db.Order(orderField)
	return &limitSQLImp{db: db}
}

var _ orderSQL = &orderSQLImp{}

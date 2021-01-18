package db

import "gorm.io/gorm"

type orderSQL interface {
	Order(string) interface {
		limitSQL
		findSQL
	}
}
type orderSQLImp struct {
	db *gorm.DB
}

func (o *orderSQLImp) Order(orderField string) interface{
	limitSQL
	findSQL
} {
	return struct {
		limitSQL
		findSQL
	}{
		limitSQL: &limitSQLImp{
			db: o.db,
		},
		findSQL: &findSQLImp{
			db: o.db,
		},
	}
}

var _ orderSQL = &orderSQLImp{}

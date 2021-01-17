package db

import "gorm.io/gorm"

type selectSQL interface {
	Select(args ...string) whereSQL
}
type selectSQLImp struct {
	db *gorm.DB
}

func (s *selectSQLImp) Select(args ...string) whereSQL {
	return &whereSQLImp{db: s.db.Select(args)}
}

var _ selectSQL = &selectSQLImp{}

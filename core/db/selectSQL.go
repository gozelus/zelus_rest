package db

import "gorm.io/gorm"

type selectSQL interface {
	Select(args ...string) whereSQL
}
type selectSQLImp struct {
	db *gorm.DB
}

func (s *selectSQLImp) Select(args ...string) whereSQL {
	return &whereSQLImp{db: s.db.Select(args).WithContext(s.db.Statement.Context)}
}

var _ selectSQL = &selectSQLImp{}

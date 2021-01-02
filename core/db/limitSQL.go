package db

import "gorm.io/gorm"

type limitSQL interface {
	Limit(int) findSQL
}

var _ limitSQL = &limitSQLImp{}

type limitSQLImp struct {
	db *gorm.DB
}

func (l *limitSQLImp) Limit(limit int) findSQL {
	db := l.db.Limit(limit)
	i := &findSQLImp{db: db}
	return i
}

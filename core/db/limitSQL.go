package db

import "gorm.io/gorm"

type limitSQL interface {
	Limit(int) interface {
		offsetSQL
		findSQL
	}
}

var _ limitSQL = &limitSQLImp{}

type limitSQLImp struct {
	db *gorm.DB
}

func (l *limitSQLImp) Limit(limit int) interface {
	offsetSQL
	findSQL
} {
	return &struct {
		*findSQLImp
		*offsetSQLImp
	}{
		findSQLImp:   &findSQLImp{db: l.db.Limit(limit)},
		offsetSQLImp: &offsetSQLImp{db: l.db.Limit(limit)},
	}
}

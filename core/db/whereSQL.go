package db

import (
	"gorm.io/gorm"
)

type whereSQL interface {
	Where(query interface{}, args ...interface{}) interface {
		whereSQL
		orderSQL
		endSQL
	}
}
type whereSQLImp struct {
	db *gorm.DB
}

var _ whereSQL = &whereSQLImp{}

func (w *whereSQLImp) Where(query interface{}, args ...interface{}) interface {
	whereSQL
	orderSQL
	endSQL
} {
	w.db = w.db.Where(query, args...)
	return struct {
		orderSQL
		whereSQL
		endSQL
	}{
		whereSQL: &whereSQLImp{db: w.db},
		orderSQL: &orderSQLImp{db: w.db},
		endSQL: &struct {
			*findSQLImp
			*updateSQLImp
			*deleteSQLImp
			*insertSQLImp
			*firstOrCreateSQLImp
			*firstSQLImp
		}{
			firstSQLImp:         &firstSQLImp{db: w.db},
			firstOrCreateSQLImp: &firstOrCreateSQLImp{db: w.db},
			findSQLImp:          &findSQLImp{db: w.db},
			updateSQLImp:        &updateSQLImp{db: w.db},
			deleteSQLImp:        &deleteSQLImp{db: w.db},
			insertSQLImp:        &insertSQLImp{db: w.db},
		},
	}
}

var _ whereSQL = &whereSQLImp{}

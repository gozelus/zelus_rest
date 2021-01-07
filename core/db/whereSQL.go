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
	db := w.db.Where(query, args...).WithContext(w.db.Statement.Context)
	return struct {
		orderSQL
		whereSQL
		endSQL
	}{
		whereSQL: &whereSQLImp{db: db},
		orderSQL: &orderSQLImp{db: db},
		endSQL: &struct {
			*findSQLImp
			*updateSQLImp
			*deleteSQLImp
			*insertSQLImp
			*firstOrCreateSQLImp
			*firstSQLImp
		}{
			firstSQLImp:         &firstSQLImp{db: db},
			firstOrCreateSQLImp: &firstOrCreateSQLImp{db: db},
			findSQLImp:          &findSQLImp{db: db},
			updateSQLImp:        &updateSQLImp{db: db},
			deleteSQLImp:        &deleteSQLImp{db: db},
			insertSQLImp:        &insertSQLImp{db: db},
		},
	}
}

var _ whereSQL = &whereSQLImp{}

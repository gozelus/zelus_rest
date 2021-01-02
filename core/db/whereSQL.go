package db

import "gorm.io/gorm"

type whereSQL interface {
	Where(query interface{}, args ...interface{}) interface {
		whereSQL
		endSQL
	}
}
type whereSQLImp struct {
	db *gorm.DB
}
var _ whereSQL = &whereSQLImp{}

func (w *whereSQLImp) Where(query interface{}, args ...interface{}) interface {
	whereSQL
	endSQL
} {
	db := w.db.Where(query, args...)
	return struct {
		whereSQL
		endSQL
	}{
		whereSQL: &whereSQLImp{db: db},
		endSQL: &struct {
			*findSQLImp
			*updateSQLImp
			*deleteSQLImp
			*insertSQLImp
		}{
			findSQLImp:   &findSQLImp{db: db},
			updateSQLImp: &updateSQLImp{db: db},
			deleteSQLImp: &deleteSQLImp{db: db},
			insertSQLImp: &insertSQLImp{db: db},
		},
	}
}

var _ whereSQL = &whereSQLImp{}

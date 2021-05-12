package db

import "gorm.io/gorm"

type offsetSQL interface {
	Offset(int64 int64) findSQL
}
type offsetSQLImp struct {
	db *gorm.DB
}

func (o *offsetSQLImp) Offset(offset int64) findSQL {
	i := &findSQLImp{db: o.db.Offset(int(offset))}
	return i
}

var _ offsetSQL = &offsetSQLImp{}

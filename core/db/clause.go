package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type clausesSQL interface {
	Clauses(conds ...clause.Expression) whereSQL
}

var _ clausesSQL = &clausesSQLImp{}

type clausesSQLImp struct {
	db *gorm.DB
}

func (f *clausesSQLImp) Clauses(conds ...clause.Expression) whereSQL {
	return &whereSQLImp{
		db: f.db.Clauses(conds...),
	}
}

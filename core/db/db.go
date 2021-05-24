package db

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDb interface {
	execSQL
	Table(ctx context.Context, name string) interface {
		whereSQL
		insertSQL
		selectSQL
		findSQL
		orderSQL
		clausesSQL
	}
	Begin() interface {
		MySQLDb
		Rollback()
		Commit()
	}
}

var _ MySQLDb = &dbImp{}

type dbImp struct {
	db *gorm.DB
	*whereSQLImp
	*insertSQLImp
	*selectSQLImp
	*findSQLImp
	*orderSQLImp
	*clausesSQLImp
	*execSQLImp
}

func (d *dbImp) Commit() {
	d.db.Commit()
}
func (d *dbImp) Rollback() {
	d.db.Rollback()
}
func (d *dbImp) Begin() interface {
	MySQLDb
	Rollback()
	Commit()
} {
	return &dbImp{db: d.db.Begin()}
}

func (d *dbImp) Table(ctx context.Context, name string) interface {
	insertSQL
	selectSQL
	orderSQL
	findSQL
	whereSQL
	clausesSQL
} {
	result := &dbImp{}
	result.db = d.db.WithContext(ctx).Session(&gorm.Session{NewDB: true}).Table(name)
	result.selectSQLImp = &selectSQLImp{db: result.db}
	result.findSQLImp = &findSQLImp{db: result.db}
	result.whereSQLImp = &whereSQLImp{db: result.db}
	result.insertSQLImp = &insertSQLImp{db: result.db}
	result.orderSQLImp = &orderSQLImp{db: result.db}
	result.clausesSQLImp = &clausesSQLImp{db: result.db}
	return result
}

func Open(dsn string) (MySQLDb, error) {
	m := &dbImp{}
	var err error
	m.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: NewDbLogger(),
	})
	return m, err
}

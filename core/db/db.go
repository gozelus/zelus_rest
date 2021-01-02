package db

import (
	"github.com/gozelus/zelus_rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySQLDb interface {
	Table(ctx rest.Context, name string) interface {
		whereSQL
		insertSQL
	}
}

var _ MySQLDb = &dbImp{}

type dbImp struct {
	db *gorm.DB
	*whereSQLImp
	*insertSQLImp
}

func (d *dbImp) Table(ctx rest.Context, name string) interface {
	insertSQL
	whereSQL
} {
	d.db = d.db.WithContext(ctx).Session(&gorm.Session{NewDB: true}).Table(name)
	d.whereSQLImp = &whereSQLImp{db: d.db}
	d.insertSQLImp = &insertSQLImp{db: d.db}
	return d
}

func Open(dsn string) (MySQLDb, error) {
	m := &dbImp{}
	var err error
	m.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	return m, err
}

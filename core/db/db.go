package db

import (
	"github.com/gozelus/zelus_rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySQLDb interface {
	Table(ctx rest.Context, name string) whereSQL
}

var _ MySQLDb = &dbImp{}

type dbImp struct {
	db *gorm.DB
	*whereSQLImp
}

func (d *dbImp) Table(ctx rest.Context, name string) whereSQL {
	d.db = d.db.WithContext(ctx).Session(&gorm.Session{NewDB: true}).Table(name)
	w := &whereSQLImp{db:d.db}
	d.whereSQLImp = w
	return d.whereSQLImp
}

func Open(dsn string) (MySQLDb, error) {
	m := &dbImp{}
	var err error
	m.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	return m, err
}

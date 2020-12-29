package sqlz

import (
	"database/sql"
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/logger"
)

func NewCommonSqlConn() Session {
	return &commonSqlConn{}
}

// Session 用于操作 MySQL 的接口
type Session interface {
	// Exec 执行一个 SQL
	// example UPDATE INSERT DELETE
	Exec(ctx rest.Context, query string, args ...interface{}) (sql.Result, error)
	// QueryRows 查询多行
	// example SELECT
	QueryRows(ctx rest.Context, v interface{}, query string, args ...interface{}) error
}

func NewSession(driverName, datasource string) Session {
	return &commonSqlConn{
		driverName: driverName,
		datasource: datasource,
	}
}

var _ Session = &commonSqlConn{}

type commonSqlConn struct {
	driverName string
	datasource string
}

func (db *commonSqlConn) Exec(ctx rest.Context, q string, args ...interface{}) (sql.Result, error) {
	var conn *sql.DB
	var err error
	conn, err = getSqlConn(db.driverName, db.datasource)
	if err != nil {
		logger.ErrorfWithStackWithContext(ctx, "get conn err for %s", err)
		return nil, err
	}
	return execContext(ctx, conn, q, args...)
}

func (db *commonSqlConn) QueryRows(ctx rest.Context, v interface{}, q string, args ...interface{}) error {
	var conn *sql.DB
	var err error
	conn, err = getSqlConn(db.driverName, db.datasource)
	if err != nil {
		logger.ErrorfWithStackWithContext(ctx, "get conn err for %s", err)
		return err
	}
	rows, err := queryContext(ctx, conn, q, args...)
	if err != nil {
		return err
	}
	return unmarshalRows(v, rows, false)
}

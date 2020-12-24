package sqlz

import "database/sql"

// Session 用于操作 MySQL 的接口
type Session interface {
	// Exec 执行一个 SQL
	// example UPDATE INSERT DELETE
	Exec(query string, args ...interface{}) (sql.Result, error)

	// QueryRow 仅查询一行
	// example SELECT
	QueryRow(v interface{}, query string, args ...interface{}) error

	// QueryRows 查询多行
	// example SELECT
	QueryRows(v interface{}, query string, args ...interface{}) error
}

type commonSqlConn struct {
	driverName string
	datasource string
}

func newCommonSqlConn() *commonSqlConn {
	return &commonSqlConn{}
}

func (db *commonSqlConn) Exec(q string, args ...interface{}) (sql.Result, error) {
	var conn *sql.DB
	var err error
	conn, err = getSqlConn(db.driverName, db.datasource)
	if err != nil {
		//logInstanceError(db.datasource, err)
		return nil, err
	}
	result, err = exec(conn, q, args...)
	return nil, err
}

func gtetSqlConn(driverName, datasource string) (*sql.DB, error) {
	return nil, nil
}

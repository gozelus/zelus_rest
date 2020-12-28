package sqlz

import (
	"database/sql"
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/core/syncz"
	"io"
	"sync"
	"time"
)

// Session 用于操作 MySQL 的接口
type Session interface {
	// Exec 执行一个 SQL
	// example UPDATE INSERT DELETE
	Exec(ctx rest.Context, query string, args ...interface{}) (*sql.Result, error)

	// QueryRow 仅查询一行
	// example SELECT
	QueryRow(ctx rest.Context, v interface{}, query string, args ...interface{}) error

	// QueryRows 查询多行
	// example SELECT
	QueryRows(ctx rest.Context, v interface{}, query string, args ...interface{}) error
}

type commonSqlConn struct {
	driverName string
	datasource string
}

func (db *commonSqlConn) QueryRow(ctx rest.Context, v interface{}, q string, args ...interface{}) error {
	var conn *sql.DB
	var err error
	conn, err = getSqlConn(db.driverName, db.datasource)
	if err != nil {
		//logInstanceError(db.datasource, err)
		return err
	}
	rows, err := queryContext(ctx, conn, q, args...)
	if err != nil {
		return err
	}
	return unmarshalRows(v, rows, false)
}

func (db *commonSqlConn) QueryRows(v interface{}, query string, args ...interface{}) error {
	panic("implement me")
}

func (db *commonSqlConn) Exec(q string, args ...interface{}) (sql.Result, error) {
	var conn *sql.DB
	var err error
	conn, err = getSqlConn(db.driverName, db.datasource)
	if err != nil {
		//logInstanceError(db.datasource, err)
		return nil, err
	}
	return exec(conn, q, args...)
}

func NewCommonSqlConn() Session {
	return &commonSqlConn{}
}

type pingedDB struct {
	*sql.DB
	once sync.Once
}

const (
	maxIdleConns = 64
	maxOpenConns = 64
	maxLifetime  = time.Minute
)

var connManager = syncz.NewResourceManager()

func getCachedSqlConn(driverName, server string) (*pingedDB, error) {
	val, err := connManager.GetResource(server, func() (io.Closer, error) {
		conn, err := newDBConnection(driverName, server)
		if err != nil {
			return nil, err
		}

		return &pingedDB{
			DB: conn,
		}, nil
	})
	if err != nil {
		return nil, err
	}

	return val.(*pingedDB), nil
}

func getSqlConn(driverName, server string) (*sql.DB, error) {
	pdb, err := getCachedSqlConn(driverName, server)
	if err != nil {
		return nil, err
	}

	pdb.once.Do(func() {
		err = pdb.Ping()
	})
	if err != nil {
		return nil, err
	}

	return pdb.DB, nil
}
func newDBConnection(driverName, datasource string) (*sql.DB, error) {
	conn, err := sql.Open(driverName, datasource)
	if err != nil {
		return nil, err
	}

	// we need to do this until the issue https://github.com/golang/go/issues/9851 get fixed
	// discussed here https://github.com/go-sql-driver/mysql/issues/257
	// if the discussed SetMaxIdleTimeout methods added, we'll change this behavior
	// 8 means we can't have more than 8 goroutines to concurrently access the same database.
	conn.SetMaxIdleConns(maxIdleConns)
	conn.SetMaxOpenConns(maxOpenConns)
	conn.SetConnMaxLifetime(maxLifetime)

	return conn, nil
}

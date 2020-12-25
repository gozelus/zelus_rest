// 此文件包含了用于统计需求的代码
// 内部会执行真正的 SQL
package sqlz

import (
	"database/sql"
)


// session 此文件中使用的 db 抽象
type session interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

// exec 用于执行真正的 exec
// 内部会进行诸如慢日志的统计
func exec(s session, q string, args ...interface{}) (sql.Result, error) {
	return s.Exec(q, args...)
}

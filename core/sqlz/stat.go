// 此文件包含了用于统计需求的代码
// 内部会执行真正的 SQL
package sqlz

import (
	"context"
	"database/sql"
	"github.com/gozelus/zelus_rest"
)

// session 此文件中使用的 db 抽象
type session interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

// exec 用于执行真正的 exec
// 内部会进行诸如慢日志的统计
func execContext(ctx rest.Context, s session, query string, args ...interface{}) (sql.Result, error) {
	return s.ExecContext(ctx, query, args...)
}
func queryContext(ctx rest.Context, s session, query string, args ...interface{}) (*sql.Rows, error) {
	return s.QueryContext(ctx, query, args...)
}

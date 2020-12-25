// 此文件包含了用于统计需求的代码
// 内部会执行真正的 SQL
package sqlz

import (
	"database/sql"
	"fmt"
	"strings"
)

// logger 用于打印譬如慢日志等
type logger interface {
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	InfofWithStack(string, ...interface{})
	WarnfWithStack(string, ...interface{})
	ErrorfWithStack(string, ...interface{})
}

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

// format 用于格式化 SQL
// TODO ?? 是否有防止 SQL 注入的效果?
func format(query string, args ...interface{}) (string, error) {
	numArgs := len(args)
	if numArgs == 0 {
		return query, nil
	}

	var b strings.Builder
	argIndex := 0

	for _, ch := range query {
		if ch == '?' {
			if argIndex >= numArgs {
				return "", fmt.Errorf("error: %d ? in sql, but less arguments provided", argIndex)
			}

			arg := args[argIndex]
			argIndex++

			switch v := arg.(type) {
			case bool:
				if v {
					b.WriteByte('1')
				} else {
					b.WriteByte('0')
				}
			case string:
				b.WriteByte('\'')
				b.WriteString(escape(v))
				b.WriteByte('\'')
			default:
				b.WriteString(mapping.Repr(v))
			}
		} else {
			b.WriteRune(ch)
		}
	}

	if argIndex < numArgs {
		return "", fmt.Errorf("error: %d ? in sql, but more arguments provided", argIndex)
	}

	return b.String(), nil
}

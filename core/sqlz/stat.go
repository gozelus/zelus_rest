// 此文件包含了用于统计需求的代码
// 内部会执行真正的 SQL
package sqlz

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/core/mapping"
	"github.com/gozelus/zelus_rest/logger"
	"github.com/gozelus/zelus_rest/timez"
	"strings"
	"time"
)

const slowThreshold = time.Millisecond * 500

// session 此文件中使用的 db 抽象
type session interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

// exec 用于执行真正的 exec
// 内部会进行诸如慢日志的统计
func execContext(ctx rest.Context, s session, query string, args ...interface{}) (sql.Result, error) {
	stmt, err := format(query, args...)
	if err != nil {
		return nil, err
	}
	logger.DebugfWithContext(ctx, "[SQL] : %s calls ", stmt)
	startTime := timez.Now()
	rows, err := s.ExecContext(ctx, query, args...)
	duration := timez.Since(startTime)
	if duration > slowThreshold {
		logger.WarnfWithContext(ctx, "[SQL] : %s slow call", stmt)
	}
	if err != nil {
		logger.ErrorfWithStackWithContext(ctx, "[SQL] : %s calls err for %s", stmt, err)
	}
	return rows, err
}
func queryContext(ctx rest.Context, s session, query string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := format(query, args...)
	if err != nil {
		return nil, err
	}
	logger.DebugfWithContext(ctx, "[SQL] : %s calls ", stmt)
	startTime := timez.Now()
	rows, err := s.QueryContext(ctx, query, args...)
	duration := timez.Since(startTime)
	if duration > slowThreshold {
		logger.WarnfWithContext(ctx, "[SQL] : %s slow call", stmt)
	}
	if err != nil {
		logger.ErrorfWithStackWithContext(ctx, "[SQL] : %s calls err for %s", stmt, err)
	}
	return rows, err
}
func escape(input string) string {
	var b strings.Builder

	for _, ch := range input {
		switch ch {
		case '\x00':
			b.WriteString(`\x00`)
		case '\r':
			b.WriteString(`\r`)
		case '\n':
			b.WriteString(`\n`)
		case '\\':
			b.WriteString(`\\`)
		case '\'':
			b.WriteString(`\'`)
		case '"':
			b.WriteString(`\"`)
		default:
			b.WriteRune(ch)
		}
	}

	return b.String()
}

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

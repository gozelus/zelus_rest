package db

import (
	"github.com/gozelus/zelus_rest"
	"github.com/gozelus/zelus_rest/logger"
	"time"
)

func exec(ctx rest.Context, sql string, f func() (int64, error)) error {
	now := time.Now()
	rows, err := f()
	consume := time.Now().Sub(now).Milliseconds()
	if err != nil {
		logger.ErrorfWithContext(ctx, "[%dms] [rows:%d] FAIL SQL : `%s` error for %s", consume, rows, sql, err)
		return err
	}
	if consume <= 200 { // 200ms
		logger.InfofWithContext(ctx, "[%dms] [rows:%d] SQL : `%s`", consume, rows, sql)
		return nil
	}
	logger.WarnfWithStackWithContext(ctx, "[%dms] [rows:%d] SLOW SQL : `%s`", rows, consume, sql)
	return nil
}

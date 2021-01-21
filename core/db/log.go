package db

import (
	"context"
	"errors"
	"github.com/gozelus/zelus_rest/logger"
	"gorm.io/gorm"
	"time"
)

func exec(ctx context.Context, sql string, f func() (int64, error)) error {
	now := time.Now()
	rows, err := f()
	consume := time.Now().Sub(now).Milliseconds()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.ErrorfWithContext(ctx, "[%dms] [rows:%d] FAIL SQL : `%s` error for %s", consume, rows, sql, err)
		}
		return err
	}
	if consume <= 200 { // 200ms
		logger.InfofWithContext(ctx, "[%dms] [rows:%d] SQL : `%s`", consume, rows, sql)
		return nil
	}
	logger.WarnfWithContext(ctx, "[%dms] [rows:%d] SLOW SQL : `%s`", consume, rows, sql)
	return nil
}

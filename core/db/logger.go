package db

import (
	"context"
	"errors"
	"github.com/gozelus/zelus_rest/logger"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"time"
)

type DBLogger struct {
	SlowSqlTime time.Duration
}

func NewDbLogger() *DBLogger {
	return &DBLogger{
		SlowSqlTime: time.Millisecond * 200,
	}
}

func (d *DBLogger) LogMode(glogger.LogLevel) glogger.Interface {
	return d
}

func (d *DBLogger) Info(ctx context.Context, str string, args ...interface{}) {
	logger.InfofWithContext(ctx, str, args...)
}

func (d *DBLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	logger.WarnfWithContext(ctx, str, args...)
}

func (d *DBLogger) Error(ctx context.Context, str string, args ...interface{}) {
	logger.ErrorfWithContext(ctx, str, args...)
}

func (d *DBLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil:
		sql, rows := fc()
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.ErrorfWithContext(ctx, "[%.3fms] [rows:%v] %s err for %s", float64(elapsed.Nanoseconds())/1e6, rows, sql, err)
		} else {
			logger.WarnfWithContext(ctx, "[%.3fms] [rows:%v] %s err for %s", float64(elapsed.Nanoseconds())/1e6, rows, sql, err)
		}
	case elapsed > d.SlowSqlTime && d.SlowSqlTime != 0:
		sql, rows := fc()
		logger.WarnfWithContext(ctx, "[%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
	default:
		sql, rows := fc()
		logger.InfofWithContext(ctx, "[%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}
}

var _ glogger.Interface = &DBLogger{}

package db

import (
	"context"
	"github.com/gozelus/zelus_rest/logger"
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
		logger.ErrorfWithContext(ctx, "err -> %s \n [%.3fms] [rows:%v] %s", err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
	case elapsed > d.SlowSqlTime && d.SlowSqlTime != 0:
		sql, rows := fc()
		logger.WarnfWithContext(ctx, "err -> %s \n [%.3fms] [rows:%v] %s", err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
	default:
		sql, rows := fc()
		logger.InfofWithContext(ctx, "err -> %s \n [%.3fms] [rows:%v] %s", err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}
}

var _ glogger.Interface = &DBLogger{}

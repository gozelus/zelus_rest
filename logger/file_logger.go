package logger

import (
	"fmt"
	"time"

	"github.com/gozelus/zelus_rest/timez"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

const (
	dateFormat      = "2006-01-02"
	timeFormat      = "2006-01-02T15:04:05.000Z07"
	hoursPerDay     = 24
	bufferSize      = 100
	defaultDirMode  = 0755
	defaultFileMode = 0600
)

// 用于外部定制 options
type Option func(o *options)
type options struct {
	// prefix 日志文件前缀
	prefix string
	// path  日志文件路径
	path string

	// MaxAge 日志文件最大保存时间
	MaxAge time.Duration
	// RotationTime 分隔时间
	RotationTime time.Duration
}

type fileLogger struct {
	// 当前正在写入的文件， fp -> 文件描述符
	fp      *rotatelogs.RotateLogs
	options options
}

func newFileLogger(path, prefix string, opts ...Option) (*fileLogger, error) {
	options := &options{
		prefix:       prefix,
		path:         path,
		MaxAge:       time.Hour * 24 * 10, // 默认保存10天
		RotationTime: time.Hour * 24,      // 默认一天一分隔
	}
	for _, o := range opts {
		o(options)
	}
	f := &fileLogger{
		options: *options,
	}

	var err error
	if f.fp, err = rotatelogs.New(fmt.Sprintf("%s/%s.%s.log", path, prefix, "%Y-%m-%d"), rotatelogs.WithMaxAge(f.options.MaxAge), rotatelogs.WithRotationTime(f.options.RotationTime)); err != nil {
		return nil, err
	}
	return f, nil
}

func (f *fileLogger) Write(data []byte) (int, error) {
	return f.fp.Write(data)
}

func getTimestamp() string {
	return timez.Time().Format(timeFormat)
}
func getNowDate() string {
	return time.Now().Format(dateFormat)
}

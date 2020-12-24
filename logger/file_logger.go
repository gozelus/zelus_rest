package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"syscall"
	"time"

	"github.com/gozelus/zelus_rest/timez"
)

const (
	dateFormat      = "2006-01-02"
	timeFormat = "2006-01-02T15:04:05.000Z07"
	hoursPerDay     = 24
	bufferSize      = 100
	defaultDirMode  = 0755
	defaultFileMode = 0600
)

type fileLogger struct {
	keepDays int
	path     string

	// prefix 日志文件前缀
	prefix string
	// 当前正在写入的文件， fp -> 文件描述符
	fp *os.File
	// 当前正在写入的文件名
	filename string
	// 预先创建的文件名
	backup string

	// 用于接收 write 方法写进来的字节
	bytesChannel chan []byte
	// 用于标志当前logger已经不接受写入
	done chan int
}

func newFileLogger(path, prefix string) (*fileLogger, error) {
	f := &fileLogger{
		prefix:       prefix,
		keepDays:     7,
		path:         path,
		fp:           nil,
		filename:     fmt.Sprintf("%s/%s-%s.log", path, prefix, getNowDate()),
		backup:       "",
		bytesChannel: make(chan []byte),
	}
	if err := f.init(); err != nil {
		return nil, err
	}
	f.startWorker()
	return f, nil
}

func (f *fileLogger) Write(data []byte) (int, error) {
	select {
	case f.bytesChannel <- data:
		return len(data), nil
	case <-f.done:
		log.Println(string(data))
		return 0, errors.New("error: log file closed")
	}
}

// 内部启动一个goroutine，接收channel中的字符并写入文件描述符内
func (f *fileLogger) startWorker() {
	go func() {
		for {
			select {
			case event := <-f.bytesChannel:
				f.fp.Write(event)
			case <-f.done:
				return
			}
		}
	}()
}

// init 根据配置创建文件夹及文件
func (f *fileLogger) init() error {
	if _, err := os.Stat(f.filename); err != nil {
		basePath := path.Dir(f.filename)
		if _, err = os.Stat(basePath); err != nil {
			if err = os.MkdirAll(basePath, defaultDirMode); err != nil {
				return err
			}
		}
		if f.fp, err = os.Create(f.filename); err != nil {
			return err
		}
	} else if f.fp, err = os.OpenFile(f.filename, os.O_APPEND|os.O_WRONLY, defaultFileMode); err != nil {
		return err
	}

	// 进程退出时关闭文件描述符?
	syscall.CloseOnExec(int(f.fp.Fd()))
	return nil
}

func getTimestamp() string {
	return timez.Time().Format(timeFormat)
}
func getNowDate() string {
	return time.Now().Format(dateFormat)
}

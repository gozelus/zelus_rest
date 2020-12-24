package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/fatih/color"
)

var (
	console = os.Stdout
	file    io.Writer

	logLevel LogLevel
	mode     LogOutputMode
	colorful bool

	once sync.Once
)

type LogOutputMode uint32

const (
	FileMode    LogOutputMode = 1 //0b001
	ConsoleMode LogOutputMode = 2 //0b010
	AllMode     LogOutputMode = 3 //0b011
)

type LogLevel uint32

const (
	DebugLogLevel LogLevel = iota + 1
	InfoLogLevel
	WarnLogLevel
	ErrorLogLevel
)

type (
	Logger interface {
		Debugf(string, ...interface{})
		Debug(...interface{})

		Infof(string, ...interface{})
		Info(...interface{})

		Warnf(string, ...interface{})
		Warn(...interface{})

		Errorf(str string, args ...interface{})
		Error(string, ...interface{})
	}
)

type LogConf struct {
	// ServiceName 用于日志文件的前缀
	ServiceName string

	// Mode 指定输出的位置
	Mode LogOutputMode

	// 若 mode 包含文件输出，则需要指定日志文件的位置，默认为 /tmp/logs
	FilePath string

	// 日志文件默认保存多少天，默认10天
	KeepDays int

	// 最低输出级别，默认 Debug
	Level LogLevel

	// Colorful 是否需要控制台的彩色输出
	Colorful bool
}

func MustSetup(c LogConf) {
	once.Do(func() {
		setupLogLevel(&c)
		setFilePath(&c)
		setupMode(&c)
		setupWithFiles(&c)
		setupColorful(&c)
	})
}

func setupColorful(c *LogConf) {
	colorful = c.Colorful
}

func setFilePath(c *LogConf) {
	if c.FilePath == "" {
		c.FilePath = "/tmp/logs"
	}
}
func setupMode(c *LogConf) {
	switch c.Mode {
	case c.Mode:
	case FileMode:
		mode = FileMode
	case ConsoleMode:
		mode = ConsoleMode
	default:
		mode = AllMode
	}
}
func setupLogLevel(c *LogConf) {
	switch c.Level {
	case DebugLogLevel:
		fallthrough
	case InfoLogLevel:
		fallthrough
	case WarnLogLevel:
		fallthrough
	case ErrorLogLevel:
		logLevel = c.Level
	default:
		logLevel = DebugLogLevel
	}
}

func setupWithFiles(c *LogConf) {
	var err error
	if file, err = newFileLogger(c.FilePath, c.ServiceName); err != nil {
		log.Fatal(err)
	}
}

func Info(v ...interface{}) {
	info(fmt.Sprint(v))
}
func Debug(v ...interface{}) {
	debug(fmt.Sprint(v))
}
func Warn(v ...interface{}) {
	warn(fmt.Sprint(v))
}
func Error(v ...interface{}) {
	errorz(fmt.Sprint(v))
}
func Infof(str string, args ...interface{}) {
	info(fmt.Sprintf(str, args...))
}
func Debugf(str string, args ...interface{}) {
	debug(fmt.Sprintf(str, args...))
}
func Warnf(str string, args ...interface{}) {
	warn(fmt.Sprintf(str, args...))
}
func Errorf(str string, args ...interface{}) {
	errorz(fmt.Sprintf(str, args...))
}

type logEntry struct {
	LevelStr  string          `json:"@level"`
	Color     color.Attribute `json:"-"`
	Timestamp string          `json:"@timestamp"`
	Content   string          `json:"content"`
	Level     LogLevel        `json:"-"`
}

func warn(msg string) {
	output(WarnLogLevel, msg)
}
func errorz(msg string) {
	// TODO need add call stack
	output(ErrorLogLevel, msg)
}
func debug(msg string) {
	output(DebugLogLevel, msg)
}
func info(msg string) {
	output(InfoLogLevel, msg)
}

func getWriters() []io.Writer {
	var writers []io.Writer
	if (mode | FileMode) > 0 {
		writers = append(writers, file)
	}
	if (mode | ConsoleMode) > 0 {
		writers = append(writers, console)
	}
	return writers
}

func output(level LogLevel, msg string) {
	shouldLog(level, func() {
		entry := &logEntry{
			Timestamp: getTimestamp(),
			Content:   msg,
			Level:     level,
		}
		switch entry.Level {
		case DebugLogLevel:
			entry.LevelStr = "[DEBUG]"
			entry.Color = color.FgHiMagenta
		case InfoLogLevel:
			entry.LevelStr = "[INFO]"
			entry.Color = color.FgHiBlue
		case WarnLogLevel:
			entry.LevelStr = "[WARN]"
			entry.Color = color.FgHiYellow
		case ErrorLogLevel:
			entry.LevelStr = "[ERROR]"
			entry.Color = color.FgHiRed
		}
		shouldLogToFile(entry)
		shouldLogToConsole(entry)
	})
}
func shouldLog(level LogLevel, f func()) {
	if level >= logLevel {
		f()
	}
}
func shouldLogToConsole(entry *logEntry) {
	if (mode | ConsoleMode) > 0 {
		outputColorString(entry, console)
	}
}
func shouldLogToFile(entry *logEntry) {
	if (mode | FileMode) > 0 {
		outputJson(entry, file)
	}
}
func outputColorString(entry *logEntry, writer io.Writer) {
	color.New(entry.Color).Fprintln(writer, fmt.Sprintf("%s %s %s", entry.LevelStr, entry.Timestamp, entry.Content))
}

func outputJson(entry *logEntry, writer io.Writer) {
	var bytes []byte
	var err error
	if bytes, err = json.Marshal(entry); err != nil {
		log.Fatal(err)
	}
	writer.Write(append(bytes, '\n'))
}

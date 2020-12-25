package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/fatih/color"
)

const (
	skipStack = 4
)
var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)
var (
	// 控制台输出
	console = os.Stdout
	// 文件输出
	file io.Writer

	// 输出级别
	logLevel LogLevel

	// 输出方式
	mode LogOutputMode

	// 控制台输出颜色控制，日志采集后带上颜色会导致读取错误，生产环境最好关掉
	colorful bool

	// 控制初始化次数
	once sync.Once

	// 是否初始化过的标志位，用原子操作读写
	initialized uint32
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

func getDefaultConf() LogConf {
	return LogConf{
		ServiceName: "api",
		Mode:        AllMode,
		FilePath:    "/tmp/logs",
		KeepDays:    10,
		Level:       DebugLogLevel,
		Colorful:    false,
	}
}

func MustSetup(c LogConf) {
	once.Do(func() {
		atomic.StoreUint32(&initialized, 1)
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

func InfoWithStack(v ...interface{}) {
	infoWithStack(fmt.Sprint(v))
}
func DebugWithStack(v ...interface{}) {
	debugWithStack(fmt.Sprint(v))
}
func WarnWithStack(v ...interface{}) {
	warnWithStack(fmt.Sprint(v))
}
func ErrorWithStack(v ...interface{}) {
	errorzWithStack(fmt.Sprint(v))
}
func InfofWithStack(str string, args ...interface{}) {
	infoWithStack(fmt.Sprintf(str, args...))
}
func DebugfWithStack(str string, args ...interface{}) {
	debugWithStack(fmt.Sprintf(str, args...))
}
func WarnfWithStack(str string, args ...interface{}) {
	warnWithStack(fmt.Sprintf(str, args...))
}
func ErrorfWithStack(str string, args ...interface{}) {
	errorzWithStack(fmt.Sprintf(str, args...))
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

func warnWithStack(msg string) {
	output(WarnLogLevel, formatWithCaller(msg, skipStack))
}
func errorzWithStack(msg string) {
	output(ErrorLogLevel, formatWithCaller(msg, skipStack))
}
func debugWithStack(msg string) {
	output(DebugLogLevel, formatWithCaller(msg, skipStack))
}
func infoWithStack(msg string) {
	output(InfoLogLevel, formatWithCaller(msg, skipStack))
}
func warn(msg string) {
	output(WarnLogLevel, msg)
}
func errorz(msg string) {
	output(ErrorLogLevel, msg)
}
func debug(msg string) {
	output(DebugLogLevel, msg)
}
func info(msg string) {
	output(InfoLogLevel, msg)
}

func output(level LogLevel, msg string) {
	if atomic.LoadUint32(&initialized) == 0 {
		MustSetup(getDefaultConf())
	}
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

func formatWithCaller(msg string, skip int) string{
	return fmt.Sprintf("%s\n\n%s", msg, stack(skip))
}

func stack(skip int) []byte {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contains dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastSlash := bytes.LastIndex(name, slash); lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

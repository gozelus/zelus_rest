package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
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
	Context interface {
		GetRequestID() string
	}
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
func InfoWithStackWithContext(ctx context.Context, v ...interface{}) {
	infoWithStackWithContext(ctx, fmt.Sprint(v))
}
func DebugWithStackWithContext(ctx context.Context, v ...interface{}) {
	debugWithStackWithContext(ctx, fmt.Sprint(v))
}
func WarnWithStackWithContext(ctx context.Context, v ...interface{}) {
	warnWithStackWithContext(ctx, fmt.Sprint(v))
}
func ErrorWithStackWithContext(ctx context.Context, v ...interface{}) {
	errorzWithStackWithContext(ctx, fmt.Sprint(v))
}
func InfofWithStackWithContext(ctx context.Context, str string, args ...interface{}) {
	infoWithStackWithContext(ctx, fmt.Sprintf(str, args...))
}
func DebugfWithStackWithContext(ctx context.Context, str string, args ...interface{}) {
	debugWithStackWithContext(ctx, fmt.Sprintf(str, args...))
}
func WarnfWithStackWithContext(ctx context.Context, str string, args ...interface{}) {
	warnWithStackWithContext(ctx, fmt.Sprintf(str, args...))
}
func ErrorfWithStackWithContext(ctx context.Context, str string, args ...interface{}) {
	errorzWithStackWithContext(ctx, fmt.Sprintf(str, args...))
}
func InfoWithContext(ctx context.Context, v ...interface{}) {
	infoWithContext(ctx, fmt.Sprint(v))
}
func DebugWithContext(ctx context.Context, v ...interface{}) {
	debugWithContext(ctx, fmt.Sprint(v))
}
func WarnWithContext(ctx context.Context, v ...interface{}) {
	warnWithContext(ctx, fmt.Sprint(v))
}
func ErrorWithContext(ctx context.Context, v ...interface{}) {
	errorzWithContext(ctx, fmt.Sprint(v))
}
func InfofWithContext(ctx context.Context, str string, args ...interface{}) {
	infoWithContext(ctx, fmt.Sprintf(str, args...))
}
func DebugfWithContext(ctx context.Context, str string, args ...interface{}) {
	debugWithContext(ctx, fmt.Sprintf(str, args...))
}
func WarnfWithContext(ctx context.Context, str string, args ...interface{}) {
	warnWithContext(ctx, fmt.Sprintf(str, args...))
}
func ErrorfWithContext(ctx context.Context, str string, args ...interface{}) {
	errorzWithContext(ctx, fmt.Sprintf(str, args...))
}
func warnWithStack(msg string) {
	outputWithContext(nil, WarnLogLevel, formatWithCaller(msg, skipStack))
}
func errorzWithStack(msg string) {
	outputWithContext(nil, ErrorLogLevel, formatWithCaller(msg, skipStack))
}
func debugWithStack(msg string) {
	outputWithContext(nil, DebugLogLevel, formatWithCaller(msg, skipStack))
}
func infoWithStack(msg string) {
	outputWithContext(nil, InfoLogLevel, formatWithCaller(msg, skipStack))
}
func warn(msg string) {
	outputWithContext(nil, WarnLogLevel, msg)
}
func errorz(msg string) {
	outputWithContext(nil, ErrorLogLevel, msg)
}
func debug(msg string) {
	outputWithContext(nil, DebugLogLevel, msg)
}
func info(msg string) {
	outputWithContext(nil, InfoLogLevel, msg)
}
func warnWithStackWithContext(ctx context.Context, msg string) {
	outputWithContext(ctx, WarnLogLevel, formatWithCaller(msg, skipStack))
}
func errorzWithStackWithContext(ctx context.Context, msg string) {
	outputWithContext(ctx, ErrorLogLevel, formatWithCaller(msg, skipStack))
}
func debugWithStackWithContext(ctx context.Context, msg string) {
	outputWithContext(ctx, DebugLogLevel, formatWithCaller(msg, skipStack))
}
func infoWithStackWithContext(ctx context.Context, msg string) {
	outputWithContext(ctx, InfoLogLevel, formatWithCaller(msg, skipStack))
}
func warnWithContext(ctx context.Context, msg string) {
	outputWithContext(ctx, WarnLogLevel, msg)
}
func errorzWithContext(ctx context.Context, msg string) {
	outputWithContext(ctx, ErrorLogLevel, msg)
}
func debugWithContext(ctx context.Context, msg string) {
	outputWithContext(ctx, DebugLogLevel, msg)
}
func infoWithContext(ctx context.Context, msg string) {
	outputWithContext(ctx, InfoLogLevel, msg)
}

type logEntry struct {
	LevelStr  string          `json:"level"`
	Color     color.Attribute `json:"-"`
	Timestamp string          `json:"timestamp"`
	ContextID string          `json:"context_id"`
	UserID    int64           `json:"user_id"`
	Caller    string          `json:"caller"`
	Message   string          `json:"message"`
	Level     LogLevel        `json:"-"`
}

func outputWithContext(ctx context.Context, level LogLevel, msg string) {
	if atomic.LoadUint32(&initialized) == 0 {
		MustSetup(getDefaultConf())
	}
	shouldLog(level, func() {
		entry := &logEntry{
			Timestamp: getTimestamp(),
			Message:   msg,
			Level:     level,
		}
		if ctx != nil {
			id, ok := ctx.Value("rest-request-id").(string)
			if ok {
				entry.ContextID = id
			}
			userid, ok := ctx.Value("jwt-user-id").(int64)
			if ok {
				entry.UserID = userid
			}
		}
		switch entry.Level {
		case DebugLogLevel:
			entry.LevelStr = "DEBUG"
			entry.Color = color.FgHiMagenta
		case InfoLogLevel:
			entry.LevelStr = "INFO"
			entry.Color = color.FgHiBlue
		case WarnLogLevel:
			entry.LevelStr = "WARN"
			entry.Color = color.FgHiYellow
		case ErrorLogLevel:
			entry.LevelStr = "ERROR"
			entry.Color = color.FgHiRed
		}
		entry.Caller = getCaller(skipStack + 2)
		//shouldLogToFile(entry)
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
	outputJson(entry, writer)
}

func outputJson(entry *logEntry, writer io.Writer) {
	var bytes []byte
	var err error
	if bytes, err = json.Marshal(entry); err != nil {
		log.Fatal(err)
	}
	writer.Write(append(bytes, '\n'))
}

func getCaller(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}
	// 不打印所有的文件路径
	paths := strings.Split(file, "/")
	var files []string
	for i := len(paths) - 1; i >= 0; i-- {
		files = append(files, paths[i])
		if len(files) > 3 {
			break
		}
	}
	for i, j := 0, len(files)-1; i < j; i, j = i+1, j-1 {
		files[i], files[j] = files[j], files[i]
	}
	return fmt.Sprintf("%s:%d", strings.Join(files, "/"), line)
}
func formatWithCaller(msg string, skip int) string {
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

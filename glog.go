package glog

import "strings"

// Level Debug Info Warning Error Fatal, DEBUG，INFO，WARNING，ERROR，FATAL
//自定义日志级别
type Level uint16

const (
	levelDebug Level = iota
	levelInfo
	levelWarning
	levelError
	levelFatal
)

// Logger 定义一个logger接口
type Logger interface {
	Debug(format string, args ...any)
	Info(format string, args ...any)
	Warning(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
	Close()
}

func getLevelStr(level Level) string {
	switch level {
	case levelDebug:
		return "DEBUG"
	case levelInfo:
		return "INFO"
	case levelWarning:
		return "WARNING"
	case levelError:
		return "ERROR"
	case levelFatal:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

//更具用户传入的字符串类型的日志级别，解析出对应的level
func parseLogLevel(leverStr string) Level {
	leverStr = strings.ToLower(leverStr) //字符串转换成小写
	switch leverStr {
	case "debug":
		return levelDebug
	case "info":
		return levelInfo
	case "warning":
		return levelWarning
	case "error":
		return levelError
	case "fatal":
		return levelFatal
	default:
		return levelDebug

	}
}

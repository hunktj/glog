package glog

import "strings"

// Level Debug Info Warning Error Fatal, DEBUG，INFO，WARNING，ERROR，FATAL
//自定义日志级别
type Level uint16

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
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
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarning:
		return "WARNING"
	case LevelError:
		return "ERROR"
	case LevelFatal:
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
		return LevelDebug
	case "info":
		return LevelInfo
	case "warning":
		return LevelWarning
	case "error":
		return LevelError
	case "fatal":
		return LevelFatal
	default:
		return LevelDebug

	}
}

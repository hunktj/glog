package glog

import (
	"fmt"
	"os"
	"time"
)

// ConsoleLogger 往终端打印日志
type ConsoleLogger struct {
	level Level
}

// NewConsoleLog 是一个生成文件日志结构体的构造函数 初始化日志文件
func NewConsoleLog(leverStr string) *ConsoleLogger {
	logLevel := parseLogLevel(leverStr)
	cl := &ConsoleLogger{
		level: logLevel,
	}
	return cl
}

//将公共的记录日志的方法封装成一个单独的方法
func (c *ConsoleLogger) log(level Level, format string, args ...any) {
	if c.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...)
	_, fileName, lineNum := getCallerInfo(3)
	logLevelStr := getLevelStr(level) //日志级别
	nowTimeStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	//logMsg := fmt.Sprintf("%s Level:[%s] func:[%s：%s] num:%d, %s", nowTimeStr, logLevelStr, funcName, fileName, lineNum, msg)
	logMsg := fmt.Sprintf("%s Level:[%s] [%s]--> %d , %s", nowTimeStr, logLevelStr, fileName, lineNum, msg)
	fmt.Fprintln(os.Stdout, logMsg)

}

//Debug 日志
func (c *ConsoleLogger) Debug(format string, args ...any) {
	c.log(levelDebug, format, args...)

}
func (c *ConsoleLogger) Info(format string, args ...any) {
	c.log(levelInfo, format, args...)

}
func (c *ConsoleLogger) Warning(format string, args ...any) {
	//往文件写日志
	c.log(levelWarning, format, args...)

}
func (c *ConsoleLogger) Error(format string, args ...any) {
	//往文件写日志
	c.log(levelError, format, args...)

}
func (c *ConsoleLogger) Fatal(format string, args ...any) {
	//往文件写日志
	c.log(levelFatal, format, args...)

}

//Close 终端标准输出不需要关闭
func (c *ConsoleLogger) Close() {

}

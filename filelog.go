package glog

import (
	"fmt"
	"os"
	"path"
	"time"
)

// Filelogger 日志写入文件的结构体
type Filelogger struct {
	level    Level
	filePath string
	fileName string
	file     *os.File //os包中的file类型指针
	errFile  *os.File
	maxSize  int64
}

//NewFileLog 是一个生成文件日志结构体的构造函数 初始化日志文件
func NewFileLog(leverStr, logFilePath, logFileName string) *Filelogger {
	logLevel := parseLogLevel(leverStr)
	fl := &Filelogger{
		level:    logLevel,
		filePath: logFilePath,
		fileName: logFileName,
		maxSize:  10 * 1024 * 1024,
	}
	fl.initFile() // 根据上面的文件路径和文件名打开文件日志
	return fl
}

//初始化文件日志文件句柄
func (f *Filelogger) initFile() {
	logName := path.Join(f.filePath, f.fileName)
	//打开文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		panic(fmt.Errorf("open file:%s,%v failed", logName, err))
	}
	f.file = fileObj
	errLogName := fmt.Sprintf("%s.err", logName)
	//打开错误日志文件
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("open file:%s,%v failed", errLogName, err))
	}
	f.file = errFileObj
}

//检查是否要拆分日志文件
func (f *Filelogger) checkSplit(file *os.File) bool {
	finfo, _ := file.Stat()
	fsize := finfo.Size()
	return fsize >= f.maxSize //当传进来的日志文件大小超过限额 返回true
}

//分装一个切分的方法
func (f *Filelogger) splitLogFile(file *os.File) *os.File {
	//切分文件
	filename := file.Name() //拿到文件的完整路径
	backupName := fmt.Sprintf("%s_%v.bak", filename, time.Now().Unix())
	//关闭原来的文件
	file.Close()
	//备份原来的文件
	os.Rename(filename, backupName)
	//新建一个文件
	fileObj, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		panic(fmt.Errorf("打开日志文件错误"))
	}
	return fileObj
}

//将公共的记录日志的方法封装成一个单独的方法
func (f *Filelogger) log(level Level, format string, args ...any) {
	if f.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...)
	_, fileName, lineNum := getCallerInfo(3)
	logLevelStr := getLevelStr(level) //日志级别
	nowTimeStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	//logMsg := fmt.Sprintf("%s Level:[%s] func:[%s：%s] num:%d, %s", nowTimeStr, logLevelStr, funcName, fileName, lineNum, msg)
	logMsg := fmt.Sprintf("%s Level:[%s] [%s]--> %d , %s", nowTimeStr, logLevelStr, fileName, lineNum, msg)
	if f.checkSplit(f.file) {
		f.file = f.splitLogFile(f.file)
	}
	fmt.Fprintln(f.file, logMsg)
	//如果是err或者是fatal级别的日志还要记录到发f.errFile
	if level >= LevelError {
		if f.checkSplit(f.errFile) {
			f.errFile = f.splitLogFile(f.errFile)
		}
		fmt.Fprintln(f.errFile, logMsg)
	}
}

//Debug 日志
func (f *Filelogger) Debug(format string, args ...any) {
	f.log(LevelDebug, format, args...)

}
func (f *Filelogger) Info(format string, args ...any) {
	f.log(LevelInfo, format, args...)

}
func (f *Filelogger) Warning(format string, args ...any) {
	//往文件写日志
	f.log(LevelWarning, format, args...)

}
func (f *Filelogger) Error(format string, args ...any) {
	//往文件写日志
	f.log(LevelError, format, args...)

}
func (f *Filelogger) Fatal(format string, args ...any) {
	//往文件写日志
	f.log(LevelFatal, format, args...)

}
func (f *Filelogger) Close() {
	//往文件写日志
	f.file.Close()
	f.errFile.Close()

}

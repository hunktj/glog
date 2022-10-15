package glog

import (
	"fmt"
	"path"
	"runtime"
)

func getCallerInfo(skip int) (string, string, int) {
	pc, fileName, line, ok := runtime.Caller(skip) //0 找到方法本身的路径，1 在那调用的路径 2第三层的方法调用路径
	//根据pc拿到当前执行的函数名
	if !ok {
		fmt.Println("runtime.Caller 执行出错")
		return "", "", 0
	}
	funcName := path.Base(runtime.FuncForPC(pc).Name())
	fileName = path.Base(fileName) //取最后的名字
	return funcName, fileName, line
}

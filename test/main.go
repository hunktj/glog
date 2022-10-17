package main

import "github.com/vpertj/glog"

var log glog.Logger

func main() {
	log = glog.NewFileLog("debug", "./", "test.log")
	userid := 10
	//for i := 0; i < 200000; i++ {
	//	log.Debug("这是一条测试日志 %d....", userid)
	//
	//}
	log.Debug("这是一条测试日志 %d....", userid)
	log.Info("这是一条测试日志 %d....", userid)
	log.Error("error log-------------")

}

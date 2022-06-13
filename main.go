package main

import (
	"gsafety/global"

	"gsafety/core"

	"go.uber.org/zap"
)

// @title Swagger Gsafety demo API
// @version 0.0.1
// @description 智慧园区demo后端服务api
// @in header
// @BasePath /
func main() {

	global.GS_VP = core.Viper() // 初始化Viper
	global.GS_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GS_LOG)

	// 启动服务
	core.RunServer()
}

package main

import (
	"fmt"
	"go_web/web_app/logger"
	"go_web/web_app/settings"

	"go.uber.org/zap"
)

// tip Go Web开发较通用的脚手架模板

func main() {
	// tip 1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("Init settins failed, err: %v\n", err)
	}

	// tip 2.初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("Init logger failed, err: %v\n", err)
	}
	zap.L().Debug("Logger init success ...")

	// tip 3.初始化MySQL连接

	// 4.初始化Redis连接
	// 5.注册路由
	// 5.启动服务（优雅关机）
}

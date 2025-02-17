package main

import (
	"go.uber.org/zap"
	"net/http"
)

// 定义全局的logger实例
var logger *zap.Logger

var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	// 程序退出之前，将缓冲区内的日志刷盘到磁盘上
	defer logger.Sync()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}

func InitLogger() {
	logger, _ = zap.NewDevelopment()
	sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		sugarLogger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

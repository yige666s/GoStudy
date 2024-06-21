package main

import (
	"day14/config"
	"day14/logger"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerTest() {
	// 1. 从config.json中加载配置
	if len(os.Args) < 1 {
		return
	}
	if err := config.Init(os.Args[1]); err != nil {
		panic(err)
	}

	// 2. 初始化logger
	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed,err :%v\n", err)
		return
	}
	gin.SetMode(config.Conf.Mode)

	// 3. 注册中间件
	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/hello", func(ctx *gin.Context) {
		var ( // 数据需要记录到日志
			name = "jack"
			age  = 19
		)
		// 4. 记录自定义日志信息
		zap.L().Debug("this is hello func", zap.String("user", name), zap.Int("age", age))
		ctx.String(http.StatusOK, "hello world")
	})

	addr := fmt.Sprintf(":%v", config.Conf.Port)
	r.Run(addr)
}

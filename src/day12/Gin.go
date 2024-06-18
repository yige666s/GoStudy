package main

import "github.com/gin-gonic/gin"

func GinDemo() {
	// 创建默认路由
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		// ctx.JSON: 返回JSON格式数据
		ctx.JSON(200, gin.H{
			"message": "hello world!",
		})
	})
	// 启动http服务
	r.Run()
}

// TODO RESTful API

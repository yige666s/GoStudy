package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func pprofTest() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080") // 监听并在 8080 端口上启动服务
}

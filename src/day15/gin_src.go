package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin源码解析
// 见gin_src.md
func gin_src() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) { ctx.String(http.StatusOK, "hello world!") })
	r.Run(":8080")
}

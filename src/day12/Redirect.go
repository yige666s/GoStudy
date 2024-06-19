package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect() {
	r := gin.Default()
	// HTTP 重定向很容易。 内部、外部重定向均支持。
	r.GET("/test", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://www.google.com")
	})

	// 路由重定向，使用HandleContext：
	r.GET("/test2", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/test"
		r.HandleContext(ctx)
	})

	r.Run(":8080")
}

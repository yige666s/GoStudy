package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// SwaggerTest()
	// LeakyBukketLimit()
	// TokenLimit()
	// TokenLimit2()
	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello", time.Now())
	})
	r.Run(":8080")
}

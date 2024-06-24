package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// ViperDemo()
	// ViperWithStruct()
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world!")
	})
	r.Run(":8080")
}

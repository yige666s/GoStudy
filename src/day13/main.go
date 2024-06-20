package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		fmt.Println("hello,world")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	r.Run(":9090")
}

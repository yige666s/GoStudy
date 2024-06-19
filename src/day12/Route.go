package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route() {
	r := gin.Default()

	// 匹配所有方法的路由
	r.Any("/test", func(ctx *gin.Context) { fmt.Println("Any route") })

	// 为没有匹配到路由的请求都返回views/404.html页面。
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "view/404.html", nil)
	})

	// 可以将拥有共同URL前缀的路由划分为一个路由组。习惯性一对{}包裹同组的路由
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(ctx *gin.Context) {})
		userGroup.GET("/login", func(ctx *gin.Context) {})
		userGroup.POST("/singup", func(ctx *gin.Context) {})
	}
	showGroup := r.Group("/shop")
	{
		showGroup.GET("/index", func(ctx *gin.Context) {})
		showGroup.GET("/cart", func(ctx *gin.Context) {})
		showGroup.POST("/checkout", func(ctx *gin.Context) {})
	}

	// 路由组也是支持嵌套的
	shopGroup := r.Group("/shop")
	{
		goods := shopGroup.Group("/goods")
		goods.GET("/info", func(ctx *gin.Context) {})
	}
	// 通常我们将路由分组用在划分业务逻辑或划分API版本时。
	// Gin框架中的路由使用的是httprouter这个库。其基本原理就是构造一个路由地址的前缀树。

	r.Run(":8080")
}

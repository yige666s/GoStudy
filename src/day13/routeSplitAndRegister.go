package main

import (
	"github.com/gin-gonic/gin"
)

// 一种路由拆分项目目录
// gin_demo
// ├── app
// │   ├── blog
// │   │   ├── handler.go
// │   │   └── router.go
// │   └── shop
// │       ├── handler.go
// │       └── router.go
// ├── go.mod
// ├── go.sum
// ├── main.go
// └── routers
//
//	└── routers.go
//
// routers/routers.go中根据需要定义Include函数用来注册子app中定义的路由
// Init函数用来进行路由的初始化操作：
type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Inlcude(opts ...Option) {
	options = append(options, opts...)
}

// 初始化路由
func Init() *gin.Engine {
	r := gin.Default()
	for _, opt := range options {
		opt(r)
	}
	return r
}

// func main1() {
// 	// // 加载多个路由配置
// 	// r := gin.Default()
// 	// // Inlcude(shop.Routers, goods.Routers)
// 	// // 初始化路由
// 	// r := Init()
// 	// if err := r.run(); err != nil {
// 	// 	fmt.Println("startup service failed,err:%v\n", err)
// 	// }
// }

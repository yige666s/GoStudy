package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// REST的含义就是客户端与Web服务器之间进行交互的时候，使用HTTP协议中的4个请求方法代表不同的动作
// 只要API程序遵循了REST风格，那就可以称其为RESTful API

func JsonRendering() {
	r := gin.Default()
	r.GET("/someJson", func(ctx *gin.Context) {
		// 1. 自己拼接Json
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello world!"}) // gin.H 是map[string]interface{}的缩写
	})

	r.GET("/moreJson", func(ctx *gin.Context) {
		// 2. 使用结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "Jack"
		msg.Message = "Hello world!"
		msg.Age = 18
		ctx.JSON(http.StatusOK, msg)
	})
	r.Run(":8080")
}

func XmlRendering() {
	r := gin.Default()
	r.GET("/someXml", func(ctx *gin.Context) {
		// 1. 自己拼接Json
		ctx.XML(http.StatusOK, gin.H{"message": "Hello world!"}) // gin.H 是map[string]interface{}的缩写
	})

	r.GET("/moreXml", func(ctx *gin.Context) {
		// 2. 使用结构体
		// 注意需要使用具名的结构体类型。
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "Jack"
		msg.Message = "Hello world!"
		msg.Age = 18
		ctx.XML(http.StatusOK, msg)
	})
	r.Run(":8080")
}

func YamlRendering() {
	r := gin.Default()
	r.GET("/someYaml", func(ctx *gin.Context) {
		ctx.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	})
	r.Run(":8080")
}

// func ProtobufRending() {
// 	r := gin.Default()
// 	r.GET("/someProtobuf", func(ctx *gin.Context) {
// 		reps := []int64{int64(1), int64(2)}
// 		label := "test"
// 		// protobuf 的具体定义写在 testdata/protoexample 文件中。
// 		data := &protoexample.Test{
// 			Label: &label,
// 			Reps:  resps,
// 		}
// 		// 请注意，数据在响应中变为二进制数据,将输出被 protoexample.Test protobuf序列化了的数据
// 		ctx.ProtoBuf(http.StatusOK, data)
// 	})
// }

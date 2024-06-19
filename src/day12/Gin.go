package main

import (
	"encoding/json"
	"fmt"
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

// querystring指的是URL中?后面携带的参数
func QueryString() {
	r := gin.Default()

	r.GET("/user/search", func(ctx *gin.Context) {
		username := ctx.DefaultQuery("username", "lili")
		// username := ctx.Query("username")
		address := ctx.Query("address")
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run(":8080")
}

// 前端POST表单请求，通过form获取请求数据
func PostRequest() {
	r := gin.Default()
	r.POST("/user/search", func(ctx *gin.Context) {
		// username := ctx.DefaultPostForm("username", "jack") // DefaultPostForm取不到值时会返回指定的默认值
		username := ctx.PostForm("username")
		address := ctx.PostForm("address")
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run(":8080")
}

// 获取Json请求数据内容
func PostWithJson() {
	r := gin.Default()
	r.POST("/json", func(ctx *gin.Context) {
		b, _ := ctx.GetRawData() // 从ctx.request.body获取请求数据
		// 定义map/struct
		var m map[string]interface{}
		// 反序列化获取数据
		_ = json.Unmarshal(b, &m)
		ctx.JSON(http.StatusOK, m)
	})
	r.Run(":8080")
}

// 通过URL获取请求参数
func PathRequest() {
	r := gin.Default()
	r.POST("/user/search/:username/:address", func(ctx *gin.Context) {
		username := ctx.Param("username")
		address := ctx.Param("address")
		ctx.JSON(http.StatusOK, gin.H{"message": "ok", "username": username, "address": address})
	})
	r.Run(":8080")
}

// 可以基于请求的Content-Type识别请求数据类型
// 并利用反射机制自动提取请求中QueryString、form表单、JSON、XML等参数到结构体中
type Login struct {
	User   string `form:"user" json:"user" binding:"required"`
	Passwd string `form:"passwd" json:"passwd" binding:"required"`
}

func ParamBinding() {
	r := gin.Default()

	// 绑定Form-data示例 user = jack, passwd = 12345
	r.POST("/loginForm", func(ctx *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-type自动选择绑定器
		err := ctx.ShouldBind(&login)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"user":   login.User,
			"passwd": login.Passwd,
		})
	})

	// 绑定Json示例 {"user":"jack","passwd":"123456"}
	r.POST("/loginJson", func(ctx *gin.Context) {
		var login Login
		err := ctx.ShouldBind(&login)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		fmt.Printf("Login info:%#v\n", login)
		ctx.JSON(http.StatusOK, gin.H{
			"user":   login.User,
			"passwd": login.Passwd,
		})
	})

	// 绑定Querystring示例 /loginQuery?user=jack&passwd=123456
	r.GET("/loginQuery", func(ctx *gin.Context) {
		var login Login
		err := ctx.ShouldBind(&login)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"user":   login.User,
			"passwd": login.Passwd,
		})
	})

	r.Run(":8080")
	// 	ShouldBind会按照下面的顺序解析请求中的数据完成绑定：
	// 如果是 GET 请求，只使用 Form 绑定引擎（query）。
	// 如果是 POST 请求，首先检查 content-type 是否为 JSON 或 XML，然后再使用 Form（form-data）。
}

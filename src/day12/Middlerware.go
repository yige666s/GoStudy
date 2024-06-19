package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。这个钩子函数就叫中间件，
// 中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等。
// Gin中的中间件必须是一个gin.HandlerFunc类型。

// 记录接口耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Set("name", "jack") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		ctx.Next()              // 调用该请求的剩余处理程序
		// ctx.Abort()             // 不调用该程序的剩余处理程序
		cost := time.Since(start) //计算耗时
		log.Println(cost)
	}
}

// 记录响应体的中间件
// 可能会想要记录下某些情况下返回给客户端的响应数据，这个时候就可以编写一个中间件来搞定。
type BodyLogWriter struct {
	gin.ResponseWriter               // 嵌入gin框架ResponseWriter
	body               *bytes.Buffer //记录使用的response
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)                  //记录一份
	return w.ResponseWriter.Write(b) //正真写入响应
}

// ginBodyLogMiddleware 一个记录返回给客户端响应体的中间件
// https://stackoverflow.com/questions/38501325/how-to-log-response-body-in-gin
func ginBodyLogMiddleware(c *gin.Context) {
	blw := &BodyLogWriter{
		body:           bytes.NewBuffer([]byte{}),
		ResponseWriter: c.Writer,
	}
	c.Writer = blw // // 使用我们自定义的类型替换默认的
	c.Next()       //执行业务逻辑
	fmt.Println("response body" + blw.body.String())
}

// 跨域中间件cors,该中间件需要注册在业务处理函数前面。
// 推荐使用社区的https://github.com/gin-contrib/cors 库，一行代码解决前后端分离架构下的跨域问题。
func CorsMiddleware() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},                                        //允许跨域发来请求的URL
		AllowMethods:     []string{"GET", "PUT", "POST", "OPTIONS"},                          //允许请求的方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},                //允许的头部字段
		ExposeHeaders:    []string{"Content-Length"},                                         // 暴漏的头部字段
		AllowCredentials: true,                                                               // 配置允许浏览器在跨域请求中发送凭证信息，从而使得跨域请求能够携带用户的身份验证信息
		AllowOriginFunc:  func(origin string) bool { return origin == "https://github.com" }, //自定义过滤源站的方法
		MaxAge:           12 * time.Hour,                                                     // 预检请求缓存CORS许可最大保存时间
	}))
	r.Run(":8080")
}

func CorsSimple() {
	r := gin.Default()
	r.Use(cors.Default())
	// 等同于
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	r.Run()
}

// 在gin框架中，我们可以为每个路由添加任意数量的中间件

func GlobalRouteMiddleware() {
	r := gin.Default()

	// 为全局路由注册
	// r.Use(StatCost()) // 注册一个全局中间件
	// r.GET("/test", func(ctx *gin.Context) {
	// 	name := ctx.MustGet("name").(string) // 从上下文获取值
	// 	log.Println(name)
	// 	ctx.JSON(http.StatusOK, gin.H{"message": "hello " + name})
	// })

	//为某个路由单独注册
	r.GET("/test2", StatCost(), func(ctx *gin.Context) {
		name := ctx.MustGet("name").(string)
		log.Println(name)
		ctx.JSON(http.StatusOK, gin.H{"message": "hello " + name})
	})

	//为路由组注册中间件
	goods := r.Group("/goods", StatCost())
	{
		goods.GET("/info", func(ctx *gin.Context) {
			name := ctx.Request.FormValue("name")
			price := ctx.Request.FormValue("price")
			ctx.JSON(http.StatusOK, gin.H{"name": name, "price": price})
			// 当在中间件或handler中启动新的goroutine时，
			// 不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）
			// 它是非线程安全的，意味着在多个goroutine中并发访问和修改Context可能会导致数据竞争和不可预测的行为。
			// go func(*gin.Context) {
			// 	ctx.Request.FormValue("name")
			// 	log.Println(name)
			// }(ctx)
		})
	}

	r.Run(":8080")
}

//gin.Default()默认使用了Logger和Recovery中间件，其中：
// Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
// Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
// 如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由

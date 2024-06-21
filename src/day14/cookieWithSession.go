package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// 状态可以理解为客户端和服务器在某次会话中产生的数据，那无状态的就以为这些数据不会被保留。
// 会话中产生的数据又是我们需要保存的，也就是说要“保持状态”。因此Cookie就是在这样一个场景下诞生。

// 一般网络用户习惯用其复数形式 Cookies，指某些网站为了辨别用户身份、
// 进行 Session 跟踪而存储在用户本地终端上的数据，而这些数据通常会经过加密处理。
// 特征：
// 浏览器发送请求的时候，自动把携带该站点之前存储的Cookie信息。
// 服务端可以设置Cookie数据。
// Cookie是针对单个域名的，不同域名之间的Cookie是独立的。
// Cookie数据可以配置过期时间，过期的Cookie数据会被系统清除。

// // net/http库，cookie会出现在请求header的http cookie，响应header的set-cookies
// func net_http_cookie() {
// 	// 在response Header中的set-cookie设置为指定的cookie
// 	http.SetCookie(http.ResponseWriter{}, &http.Cookie{})
// }

// // 解析并返回该Request的Cookie头设置的所有cookie
// func (r *Request) Cookies() []*Cookie

// // 返回请求中名为name的cookie，如果未找到该cookie会返回nil, ErrNoCookie。
// func (r *Request) Cookie(name string) (*Cookie, error)

// // 向请求中添加一个cookie
// func (r *Request) AddCookie(c *Cookie)

// Gin操作Cookies
func CookiesWithGin() {
	r := gin.Default()
	r.GET("/cookie", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("gin_cookie") //获取cookie
		if err != nil {
			cookie = "NotSet"
			ctx.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("cookie value is %v\n", cookie)
		ctx.JSON(http.StatusOK, gin.H{
			"cookie": cookie,
		})
	})
	r.Run(":8080")
}

// Session支持更多的字节，并且他保存在服务器，有较高的安全性
// 用户登陆成功之后，我们在服务端为每个用户创建一个特定的session和一个唯一的标识，它们一一对应。其中：
// Session是在服务端保存的一个数据结构，用来跟踪用户的状态，这个数据可以保存在集群、数据库、文件中；
// 唯一标识通常称为Session ID会写入用户的Cookie中。
// 这样该用户后续再次访问时，请求会自动携带Cookie数据（其中包含了Session ID），服务器通过该Session ID就能找到与之对应的Session数据，也就知道来的人是“谁”。

var (
	key   = []byte("secert_key")
	store = sessions.NewCookieStore(key)
)

// SessionMiddleware is a Gin middleware that manages sessions using Gorilla sessions.
func SessionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, _ := store.Get(ctx.Request, "sessoin_1")
		ctx.Set("session", session)
		ctx.Next()
	}
}

// GetSession retrieves the session from the Gin context.
func GetSession(ctx *gin.Context) *sessions.Session {
	session, exist := ctx.Get("session")
	if !exist {
		fmt.Println("seesion is empty")
		return nil
	}
	return session.(*sessions.Session)
}

func SessionMiddlewareTest() {
	r := gin.Default()
	r.Use(SessionMiddleware())
	// 获取session并设置
	r.GET("/", func(ctx *gin.Context) {
		session := GetSession(ctx)
		if session == nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get session"})
			return
		}
		// 通过一个counter观察session记住了之前的请求
		var counter int
		v := session.Values["counter"]
		if v == nil {
			counter = 0
		} else {
			counter = v.(int)
		}
		counter++
		session.Values["counter"] = v
		session.Save(ctx.Request, ctx.Writer)

		ctx.JSON(http.StatusOK, gin.H{"count": counter})
	})

	// TODO 清除session
	r.GET("/clear", func(ctx *gin.Context) {

	})
}

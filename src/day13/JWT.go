package main

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWT全称JSON Web Token是一种跨域认证解决方案，属于一个开放的标准，
// 它规定了一种Token实现方式，目前多用于前后端分离项目和OAuth2.0业务场景下。

// 用于签名的字符串
var mySigningKey = []byte("helloworld!")

// GenRegisteredClaims 生成默认Jwt
func GenRegisteredClaims() (string, error) {
	// 创建Claims声明
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)), // Token过期时间
		Issuer:    "jack",                                            //签发人
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 生成token对象
	return token.SignedString(mySigningKey)                    //生成签名字符串
}

// ParseRegisteredClaims 解析jwt
func ValidateRegisterClaims(tokenString string) bool {
	// 解析token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil { // 解析token失败
		return false
	}
	return token.Valid
}

type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
	// 可根据需要自行添加字段
}

const TokenExpireDuration = time.Hour * 1                 // token有效期为1小时
var CustomSecret = []byte("php is the best dev language") // 用于加盐的字符串
func GenToekn(username string) (string, error) { // 自定义Token生成函数
	claims := CustomClaims{
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "my-project",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims) // 使用指定的签名方法创建签名对象
	return token.SignedString(CustomSecret)                    // 使用指定的secret签名并获得完整的编码后的字符串token
}

func ParseToken(tokenString string) (*CustomClaims, error) { //自定义Token解析函数
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式错误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		// 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		c.Next()
	}
}

func JWTTest() {
	r := gin.Default()
	r.GET("/home", JWTAuthMiddleware(), func(c *gin.Context) { //加了拦截器校验的接口
		username := c.MustGet("username").(string)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"Msg":  "success",
			"data": gin.H{
				"username": username,
			},
		})
	})
	r.Run(":8080")
}

// 校验用户名,密码->发放token->访问其他接口,拦截器校验->成功访问

// 以下为令牌刷新流程
//   +--------+                                           +---------------+
// |        |--(A)------- Authorization Grant --------->|               |
// |        |                                           |               |
// |        |<-(B)----------- Access Token -------------|               |
// |        |               & Refresh Token             |               |
// |        |                                           |               |
// |        |                            +----------+   |               |
// |        |--(C)---- Access Token ---->|          |   |               |
// |        |                            |          |   |               |
// |        |<-(D)- Protected Resource --| Resource |   | Authorization |
// | Client |                            |  Server  |   |     Server    |
// |        |--(E)---- Access Token ---->|          |   |               |
// |        |                            |          |   |               |
// |        |<-(F)- Invalid Token Error -|          |   |               |
// |        |                            +----------+   |               |
// |        |                                           |               |
// |        |--(G)----------- Refresh Token ----------->|               |
// |        |                                           |               |
// |        |<-(H)----------- Access Token -------------|               |
// +--------+           & Optional Refresh Token        +---------------+

// 						 Figure 2: Refreshing an Expired Access Token

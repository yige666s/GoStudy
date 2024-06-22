package main

// Swagger本质上是一种用于描述使用JSON表示的RESTful API的接口描述语言
// 想要使用gin-swagger为你的代码自动生成接口文档，一般需要下面三个步骤：
// 按照swagger要求给接口代码添加声明式注释，具体参照声明式注释格式。
// 使用swag工具扫描代码自动生成API接口文档数据
// 使用gin-swagger渲染在线接口文档页面

import (
	_ "day15/docs" // 千万不要忘了导入把你上一步生成的docs
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityID int64 `json:"community_id" form:"community_id"` // 可以为空
	Page        int64 `json:"page" form:"page" example:"1"`     // 页码
	Size        int64 `json:"size" form:"size" example:"10"`    // 每页数据量
	// Order       string `json:"order" form:"order" example:"score"` // 排序依据
}

// _ResponsePostList 帖子列表接口响应数据
// type _ResponsePostList struct {
// 	Code    ResCode                 `json:"code"`    // 业务响应状态码
// 	Message string                  `json:"message"` // 提示信息
// 	Data    []*models.ApiPostDetail `json:"data"`    // 数据
// }

// Ping example
// @Summary Show a ping message
// @Description get a ping message
// @Tags example
// @Accept  json
// @Produce  json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func pingHandler(c *gin.Context) {
	c.String(200, "pong")
}

// GetPostListHandler2 升级版帖子列表接口
// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
// @title Swagger Example API
func GetPostListHandler2(c *gin.Context) {
	// GET请求参数(query string)：/api/v1/posts2?page=1&size=10&order=time
	// 初始化结构体时指定初始参数
	p := &ParamPostList{
		Page: 1,
		Size: 10,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		fmt.Println("GetPostListHandler2 with invalid params, err: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "invalid params"})
		return
	}

	//获取数据
	// data, err := logic.GetPostListNew(p)
	// if err != nil {
	// 	fmt.Println("GetPostListNew failed", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "GetPostListNew failed"})
	// 	return
	// }

	// 返回相应
	// c.JSON(http.StatusOK, data)
}

// @version 1.0
// @description This is a sample server for a pet store.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
func SwaggerTest() {
	r := gin.Default()
	r.GET("/ping", pingHandler)
	r.POST("/post2", GetPostListHandler2)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}

// gin-swagger同时还提供了DisablingWrapHandler函数，方便我们通过设置某些环境变量来禁用Swagger。例如：

// r.GET("/swagger/*any", gs.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))
// 此时如果将环境变量NAME_OF_ENV_VARIABLE设置为任意值，则/swagger/*any将返回404响应，就像未指定路由时一样。

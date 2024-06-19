package fileupload

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func FileUpload() {
	r := gin.Default()
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以修改 router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(ctx *gin.Context) {
		// 单个文件
		file, err := ctx.FormFile("f1")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		log.Println(file.Filename)

		// 指定上传位置
		dst := fmt.Sprintf("/tmp/%s", file.Filename)
		err1 := os.MkdirAll(dst, 0755)
		if err1 != nil {
			fmt.Println("Error creating directory :", err1)
			return
		}
		ctx.SaveUploadedFile(file, dst)
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})
	r.Run(":8080")
}

func MutiFileUpload() {
	r := gin.Default()
	// r.MaxMultipartMemory = 8 << 20 // 8MB
	r.POST("/mutilUpload", func(ctx *gin.Context) {
		// Mutilpart form
		form, _ := ctx.MultipartForm()
		files := form.File["file"]
		// 遍历上传
		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("/tmp/%s_%d", file.Filename, index)
			err1 := os.MkdirAll(dst, 0755)
			if err1 != nil {
				fmt.Println("Error creating directory :", err1)
				return
			}

			ctx.SaveUploadedFile(file, dst)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded", len(files)),
		})
	})
	r.Run(":8080")
}

package user

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// uploadFile 用于处理上传文件的请求
func uploadFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		log.Printf("上传文件失败,错误原因: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "上传文件失败",
		})
		return
	}

	files := form.File["files"]

	//遍历所有文件
	for index, file := range files {
		log.Printf("index: %v, file: %v", index, file.Filename)
		dst := fmt.Sprintf("./uploads/houses/%v/%v_%v", c.MustGet("house_id"), file.Filename, index)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			log.Printf("保存文件失败,错误原因: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "保存文件失败",
			})
			return
		}
		log.Printf("文件保存成功,保存路径: %v", dst)
	}

	log.Printf("共计上传文件数量: %v", len(files))
}

// deleteFile 用于删除文件
func deleteFile(c *gin.Context, filename string) {
	err := os.Remove(filename)
	if err != nil {
		log.Printf("删除文件失败,错误原因: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "删除文件失败",
		})
		return
	}
	log.Printf("文件删除成功,文件路径: %v", filename)
}

// ReleaseGet 用于处理用户发布信息界面的GET请求
func ReleaseGet(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

// ReleasePost 用于处理用户发布信息界面的POST请求,用于发布新的房源
func ReleasePost(c *gin.Context) {
	uploadFile(c)

	c.JSON(http.StatusOK, gin.H{
		"url":     "/user/release",
		"method":  "POST",
		"message": "发布成功",
	})
}

// ReleasePut 用于处理用户发布信息界面的PUT请求,用于更新用户发布的房源信息
func ReleasePut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/user/release",
		"method":  "PUT",
		"user_id": c.Param("user_id"),
	})
}

// ReleaseDelete 用于删除用户发布的房源信息
func ReleaseDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/user/release",
		"method":  "Delete",
		"user_id": c.Param("user_id"),
	})
}

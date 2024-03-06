package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReleaseGet 用于处理用户发布信息界面的GET请求
func ReleaseGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/user/release",
		"method":  "GET",
	})
}

// ReleasePost 用于处理用户发布信息界面的POST请求,用于发布新的房源
func ReleasePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/user/release",
		"method":  "POST",
	})
}

// ReleaseDelete 用于删除用户发布的房源信息
func ReleaseDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/user/release",
		"method":  "Delete",
	})
}

package auth

import (
	"github.com/gin-gonic/gin"
)

// RegisterGet 用于处理用户的注册界面的GET请求
func RegisterGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/auth/register",
		"method":  "GET",
	})
}

// RegisterPost 用于处理用户的注册界面的POST请求
func RegisterPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/auth/register",
		"method":  "POST",
	})
}

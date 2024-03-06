package auth

import (
	"github.com/gin-gonic/gin"
)

// LoginGet 用于处理用户的登录界面的GET请求
func LoginGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/auth/login",
		"method":  "GET",
	})
}

// LoginPost 用于处理用户的登录界面的POST请求
func LoginPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/auth/login",
		"method":  "POST",
	})
}

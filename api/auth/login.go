package auth

import (
	"github.com/gin-gonic/gin"
	"log"
	"online-house-trading-platform/pkg/model"
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
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Printf("error: %v", err)
	}
	c.JSON(200, gin.H{
		"message":  "/auth/login",
		"method":   "POST",
		"username": user.Username,
		"password": user.Password,
	})
}

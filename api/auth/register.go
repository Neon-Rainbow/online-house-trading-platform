package auth

import (
	"github.com/gin-gonic/gin"
	"log"
	"online-house-trading-platform/pkg/model"
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
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Printf("error: %v", err)
	}
	c.JSON(200, gin.H{
		"message":  "/auth/register",
		"method":   "POST",
		"username": user.Username,
		"email":    user.Email,
	})
	//TODO: 这里需要实现保存用户信息到数据库
}

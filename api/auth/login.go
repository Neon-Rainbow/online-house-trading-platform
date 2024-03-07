package auth

import (
	"log"
	"net/http"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
)

// LoginGet 用于处理用户的登录界面的GET请求
func LoginGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
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
	c.JSON(http.StatusOK, gin.H{
		"message":  "/auth/login",
		"method":   "POST",
		"username": user.Username,
		"password": user.Password,
	})
}

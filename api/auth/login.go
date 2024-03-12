package auth

import (
	"log"
	"net/http"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
)

// LoginGet 用于处理用户的登录界面的GET请求
// 返回状态码200和登录界面的信息
func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "/auth/login",
	//	"method":  "GET",
	//})
}

// LoginPost 用于处理用户的登录界面的POST请求
// 返回状态码200和登录成功的信息
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

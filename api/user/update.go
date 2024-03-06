package user

import (
	"github.com/gin-gonic/gin"
)

// ProfileUpdateGet 用于获取更新用户信息的界面,需要登陆后访问
func ProfileUpdateGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/user/profile/update",
		"method":  "GET",
	})
}

// ProfileUpdatePost 用于处理更新用户信息的请求,需要登陆后访问
func ProfileUpdatePost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/user/profile/update",
		"method":  "POST",
	})
}

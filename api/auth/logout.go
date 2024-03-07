package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LogoutPost 用于处理用户的登出界面的POST请求
func LogoutPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/auth/logout",
		"method":  "POST",
	})
}

package middleware

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/controller"

	"github.com/gin-gonic/gin"
)

// AdminMiddleWare 用于检查用户是否为管理员
func AdminMiddleWare() func(c *gin.Context) {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" {
			controller.ResponseErrorWithCode(c, codes.NoPermission)
			c.Abort()
			return
		}
		c.Next()
	}
}

package middleware

import (
	"net/http"
	"online-house-trading-platform/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "无效的token",
			})
			c.Abort()
			return
		}
		c.Set("user_id", mc.UserID)
		c.Set("username", mc.Username)
		c.Next()
	}
}

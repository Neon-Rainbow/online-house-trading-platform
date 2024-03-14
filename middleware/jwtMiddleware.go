package middleware

import (
	"log"
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

		var tkn string
		// 如果请求头部的token格式不正确，则返回错误信息
		// 正确的格式为：Bearer token 或者 token
		if len(parts) == 1 {
			tkn = parts[0]
		} else if len(parts) == 2 && parts[0] == "Bearer" {
			tkn = parts[1]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "请求携带的token格式错误，无权限访问",
			})
			c.Abort()
			return
		}

		log.Printf("authHeader: %v", authHeader)
		log.Printf("parts: %v", parts)

		mc, err := jwt.ParseToken(tkn)
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

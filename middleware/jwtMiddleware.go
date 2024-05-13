package middleware

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/controller"
	"online-house-trading-platform/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseErrorWithCode(c, codes.RequestWithoutTokenError)
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
			controller.ResponseErrorWithCode(c, codes.InvalidTokenFormatError)
			c.Abort()
			return
		}

		//log.Printf("authHeader: %v", authHeader)
		//log.Printf("parts: %v", parts)

		mc, err := jwt.ParseToken(tkn)
		if err != nil {
			controller.ResponseErrorWithCode(c, codes.InvalidTokenError)
			c.Abort()
			return
		}
		//fmt.Println(mc)
		c.Set("user_id", mc.UserID)
		c.Set("username", mc.Username)
		c.Next()
	}
}

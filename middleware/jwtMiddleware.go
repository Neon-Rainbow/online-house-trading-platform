package middleware

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/controller"
	"online-house-trading-platform/pkg/jwt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			zap.L().Error("JWTAuthMiddleware:请求头部没有token",
				zap.String("错误码", strconv.FormatInt(int64(codes.RequestWithoutTokenError), 10)),
				zap.String("authHeader", authHeader),
			)
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
			zap.L().Error("JWTAuthMiddleware: jwt token格式错误",
				zap.String("错误码", strconv.FormatInt(int64(codes.InvalidTokenFormatError), 10)),
				zap.String("authHeader", authHeader),
				zap.Strings("parts", parts),
			)
			controller.ResponseErrorWithCode(c, codes.InvalidTokenFormatError)
			c.Abort()
			return
		}

		myClaims, err := jwt.ParseToken(tkn)
		if err != nil {
			zap.L().Error("JWTAuthMiddleware: jwt token解析错误",
				zap.String("错误码", strconv.FormatInt(int64(codes.InvalidTokenError), 10)),
				zap.String("token", tkn),
			)
			controller.ResponseErrorWithCode(c, codes.InvalidTokenError)
			c.Abort()
			return
		}

		if myClaims.TokenType != "access_token" {
			zap.L().Error("JWTAuthMiddleware: 非访问令牌尝试访问资源",
				zap.String("tokenType", myClaims.TokenType),
			)
			controller.ResponseErrorWithCode(c, codes.UnauthorizedAccessError)
			c.Abort()
			return
		}

		c.Set("user_id", myClaims.UserID)
		c.Set("username", myClaims.Username)
		c.Set("role", myClaims.UserRole)
		c.Next()
	}
}

package middleware

import (
	"online-house-trading-platform/internal/controller"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UserIDMatchMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Gin上下文中获取user_id（假设是从JWT中间件注入的）
		contextUserID, exists := c.Get("user_id")
		if !exists {
			// 如果在上下文中找不到user_id，返回错误响应
			zap.L().Error("UserIDMatchMiddleware: c.Get(\"user_id\") failed",
				zap.String("错误码", strconv.FormatInt(int64(1017), 10)),
			)
			controller.ResponseErrorWithCode(c, 1017)
			c.Abort()
			return
		}

		// 从请求的URL参数中获取user_id
		paramUserIDStr := c.Param("user_id")
		paramUserID, err := strconv.Atoi(paramUserIDStr)
		paramUserIDUint := uint(paramUserID)
		if err != nil {
			// 如果URL参数中的user_id不是有效的整数，返回错误响应
			zap.L().Error("UserIDMatchMiddleware: strconv.Atoi(paramUserIDStr) failed",
				zap.String("错误码", strconv.FormatInt(int64(1018), 10)),
			)
			controller.ResponseErrorWithCode(c, 1018)
			c.Abort()
			return
		}

		// 比较上下文中的user_id和URL参数中的user_id
		if contextUserID != paramUserIDUint {
			// 如果不匹配，返回错误响应
			zap.L().Error("UserIDMatchMiddleware: user_id不匹配",
				zap.String("错误码", strconv.FormatInt(int64(1032), 10)),
				zap.String("contextUserID", strconv.Itoa(int(contextUserID.(uint)))),
				zap.String("paramUserID", paramUserIDStr),
			)
			controller.ResponseErrorWithCode(c, 1032)
			c.Abort()
			return
		}

		// 如果匹配，继续后续的处理器
		c.Next()
	}
}

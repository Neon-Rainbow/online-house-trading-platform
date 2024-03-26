package middleware

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserIDMatchMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Gin上下文中获取user_id（假设是从JWT中间件注入的）
		contextUserID, exists := c.Get("user_id")
		if !exists {
			// 如果在上下文中找不到user_id，返回错误响应
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权访问"})
			log.Printf("未授权访问,context中未能找到JWT中间件中注入的user_id\n请求IP: %v\n请求url: %v", c.ClientIP(), c.Request.URL)
			c.Abort()
			return
		}

		// 从请求的URL参数中获取user_id
		paramUserIDStr := c.Param("user_id")
		paramUserID, err := strconv.Atoi(paramUserIDStr)
		paramUserIDUint := uint(paramUserID)
		if err != nil {
			// 如果URL参数中的user_id不是有效的整数，返回错误响应
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
			log.Printf("无效的用户ID,URL参数中的user_id不是有效的整数\n请求IP: %v\n请求url: %v", c.ClientIP(), c.Request.URL)
			c.Abort()
			return
		}

		// 比较上下文中的user_id和URL参数中的user_id
		if contextUserID != paramUserIDUint {
			// 如果不匹配，返回错误响应
			c.JSON(http.StatusForbidden, gin.H{"error": "无权访问此资源,因为token中的user_id与URL参数中的user_id不匹配"})
			log.Printf("无权访问此资源,因为token中的user_id与URL参数中的user_id不匹配\ntoken中的user_id: %v\nURL参数中的user_id: %v\n请求IP: %v\n请求url: %v", contextUserID, paramUserID, c.ClientIP(), c.Request.URL)
			c.Abort()
			return
		}

		// 如果匹配，继续后续的处理器
		c.Next()
	}
}

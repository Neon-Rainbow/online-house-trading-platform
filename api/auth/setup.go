// Package auth 用于处理用户的登录和注册请求
package auth

import (
	"online-house-trading-platform/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetUpAuthAPI 建设了一个用户API,用于处理用户的登录和注册,url为/auth
// 该API包含了以下功能:
// 1. 用户登录,其中GET请求用于获取登录界面,POST请求用于处理登录请求
// 2. 用户注册,其中GET请求用于获取注册界面,POST请求用于处理注册请求
// 3. 用户登出,其中POST请求用于处理登出请求
// 该API使用了数据库中间件,用于获取数据库连接
func SetUpAuthAPI(r *gin.Engine, db *gorm.DB) {
	userGroup := r.Group("/auth").Use(middleware.DBMiddleware(db))
	{
		userGroup.GET("/login", LoginGet)
		userGroup.POST("/login", LoginPost)

		userGroup.GET("/register", RegisterGet)
		userGroup.POST("/register", RegisterPost)

		userGroup.POST("/logout", LogoutPost)
	}
}

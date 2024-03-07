// Package auth 用于处理用户的登录和注册请求
package auth

import (
	"online-house-trading-platform/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetUpAuthAPI 建设了一个用户API,用于处理用户的登录和注册,url为/auth
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

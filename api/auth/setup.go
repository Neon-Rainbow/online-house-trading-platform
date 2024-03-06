// Package auth 用于处理用户的登录和注册请求
package auth

import "github.com/gin-gonic/gin"

// SetUpAuthAPI 建设了一个用户API,用于处理用户的登录和注册,url为/auth
func SetUpAuthAPI(r *gin.Engine) {
	userGroup := r.Group("/auth")
	{
		userGroup.GET("/login", LoginGet)
		userGroup.POST("/login", LoginPost)

		userGroup.GET("/register", RegisterGet)
		userGroup.POST("/register", RegisterPost)

		userGroup.POST("/logout", LogoutPost)
	}
}

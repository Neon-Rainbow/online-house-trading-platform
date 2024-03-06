package user

import (
	"github.com/gin-gonic/gin"
)

// SetUpUserAPI 用于设置用户个人信息相关的路由
func SetUpUserAPI(router *gin.Engine) {
	userGroup := router.Group("/user/profile")
	{
		userGroup.GET("/", ProfileGet)
		userGroup.POST("/", ProfilePost)
		userGroup.GET("/update", ProfileUpdateGet)
		userGroup.POST("/update", ProfileUpdatePost)
	}
}

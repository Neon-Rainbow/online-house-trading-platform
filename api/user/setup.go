// Package user 提供了用户个人信息相关的API,包括用户个人资料,用户发表的信息,用户收藏的信息等
package user

import (
	"online-house-trading-platform/api/user/profile"

	"github.com/gin-gonic/gin"
)

// SetUpUserAPI 用于设置用户个人信息相关的路由
func SetUpUserAPI(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userProfileGroup := userGroup.Group("/profile")
		{
			userProfileGroup.GET("/", profile.ProfileGet)
			//userGroup.POST("/", ProfilePost)
			userProfileGroup.GET("/update", profile.ProfileUpdateGet)
			userProfileGroup.POST("/update", profile.ProfileUpdatePost)
		}
		userGroup.GET("/favourites", ViewFavouritesGet)
		userGroup.DELETE("/favourites", DeleteFavouritesDelete)

		userGroup.GET("/release", ReleaseGet)
		userGroup.POST("/release", ReleasePost)
		userGroup.DELETE("/release", ReleaseDelete)

		userGroup.GET("/appointment", UserAppointmentGet)
		userGroup.POST("/appointment", UserAppointmentPost)
		userGroup.DELETE("/appointment", UserAppointmentDelete)
	}
}

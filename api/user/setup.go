// Package user 提供了用户个人信息相关的API,包括用户个人资料,用户发表的信息,用户收藏的信息等
package user

import (
	"online-house-trading-platform/api/user/profile"
	"online-house-trading-platform/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetUpUserAPI 用于设置用户个人信息相关的路由
func SetUpUserAPI(router *gin.Engine, db *gorm.DB) {
	userGroup := router.Group("/user/:user_id")
	{
		userProfileGroup := userGroup.Group("/profile").Use(middleware.DBMiddleware(db))
		{
			userProfileGroup.GET("/", profile.ProfileGet)
			userProfileGroup.GET("/update", profile.ProfileUpdateGet)
			userProfileGroup.POST("/update", profile.ProfileUpdatePost)
		}

		userOtherGroup := userGroup.Group("/").Use(middleware.DBMiddleware(db))
		{
			userOtherGroup.GET("/favourites", ViewFavouritesGet)
			userOtherGroup.DELETE("/favourites", DeleteFavouritesDelete)

			userOtherGroup.GET("/release", ReleaseGet)
			userOtherGroup.PUT("/release", ReleasePut)
			userOtherGroup.POST("/release", ReleasePost)
			userOtherGroup.DELETE("/release", ReleaseDelete)

			userOtherGroup.GET("/appointment", UserAppointmentGet)
			userOtherGroup.POST("/appointment", UserAppointmentPost)
			userOtherGroup.DELETE("/appointment", UserAppointmentDelete)
		}

	}
}

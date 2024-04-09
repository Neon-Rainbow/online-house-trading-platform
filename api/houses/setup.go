// Package houses 用来处理房屋相关的请求,包括添加房屋,删除房屋,收藏房屋等
package houses

import (
	"online-house-trading-platform/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetUpHousesAPI 建设了一个房屋API,用于处理用户的添加房屋和删除房屋,url为/houses
func SetUpHousesAPI(r *gin.Engine, db *gorm.DB) {
	housesGroup := r.Group("/houses").Use(middleware.DBMiddleware(db))
	{
		housesGroup.GET("/", GetHouseList)

		housesGroup.GET("/:id", HouseByIDGet)

		//下面的四个接口由于业务逻辑的修改,现在已经弃用
		//housesGroup.GET("/add", AddGet)
		//housesGroup.POST("/add", AddPost)
		//housesGroup.GET("/delete", DeleteGet)
		//housesGroup.POST("/delete", DeletePost)

		//预约和收藏需要使用JWT中间件,即需要用户登录后才能进行操作
		housesGroup.POST("/appointment", middleware.JWTAuthMiddleware(), HousesAppointmentPost)

		housesGroup.POST("/collect", middleware.JWTAuthMiddleware(), CollectHousesPost)
	}
}

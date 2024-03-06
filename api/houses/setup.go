// Package houses 用来处理房屋相关的请求,包括添加房屋,删除房屋,收藏房屋等
package houses

import (
	"github.com/gin-gonic/gin"
)

// SetUpHousesAPI 建设了一个房屋API,用于处理用户的添加房屋和删除房屋,url为/houses
func SetUpHousesAPI(r *gin.Engine) {
	housesGroup := r.Group("/houses")
	{
		housesGroup.GET("/", HouseListGet)

		housesGroup.GET("/:id", HouseByIDGet)

		//下面四个请求均重定向到了/user/release中
		housesGroup.GET("/add", AddGet)
		housesGroup.POST("/add", AddPost)

		housesGroup.GET("/delete", DeleteGet)
		housesGroup.POST("/delete", DeletePost)

		housesGroup.POST("/appointment", HousesAppointmentPost)

		housesGroup.POST("/collect", CollectHousesPost)
	}
}

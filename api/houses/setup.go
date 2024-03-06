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

		housesGroup.GET("/add", AddGet)
		housesGroup.POST("/add", AddPost)

		housesGroup.GET("/delete", DeleteGet)
		housesGroup.POST("/delete/:id", DeletePost)
	}
}

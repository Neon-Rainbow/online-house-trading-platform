package houses

import (
	"github.com/gin-gonic/gin"
)

// HouseListGet 用于处理用户的添加房屋界面的GET请求
func HouseListGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/houses/",
		"method":  "GET",
	})
}

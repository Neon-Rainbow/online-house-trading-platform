package houses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HouseListGet 用于处理用户的添加房屋界面的GET请求
func HouseListGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/houses/",
		"method":  "GET",
	})
}

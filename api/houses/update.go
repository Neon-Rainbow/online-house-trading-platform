package houses

import (
	"github.com/gin-gonic/gin"
)

// UpdateGet 用于处理用户的添加房屋界面的GET请求
func UpdateGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/houses/update",
		"method":  "GET",
	})
}

// UpdatePost 用于处理用户的添加房屋界面的POST请求
func UpdatePost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/houses/update",
		"method":  "POST",
	})
}

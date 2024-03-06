package houses

import (
	"github.com/gin-gonic/gin"
)

// DeleteGet 用于处理用户的添加房屋界面的GET请求
func DeleteGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/houses/delete",
		"method":  "GET",
	})
}

func DeletePost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/houses/delete",
		"method":  "POST",
		"id":      c.Param("id"),
	})
}

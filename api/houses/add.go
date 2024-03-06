package houses

import (
	"github.com/gin-gonic/gin"
)

// AddGet 用于处理用户的添加房屋界面的GET请求
func AddGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/houses/add",
		"method":  "GET",
	})
}

// AddPost 用于处理用户的添加房屋界面的POST请求
func AddPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/houses/add",
		"method":  "POST",
	})
}

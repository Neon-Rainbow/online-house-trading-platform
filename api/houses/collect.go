package houses

import (
	"github.com/gin-gonic/gin"
)

// CollectHousesPost 用于在/houses界面,即房屋列表界面收藏房屋
func CollectHousesPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/houses/collect",
		"method":  "POST",
	})
}

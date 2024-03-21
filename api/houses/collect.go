package houses

import (
	"net/http"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CollectHousesPost 用于在/houses界面,即房屋列表界面收藏房屋
func CollectHousesPost(c *gin.Context) {
	db, exists := c.MustGet("db").(*gorm.DB)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法获取数据库连接",
		})
		return
	}
	var favourite model.Favourite
	err := c.ShouldBind(&favourite)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}
	err = db.Create(&favourite).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "收藏失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "收藏成功",
		"url":     "/houses/collect",
	})
}

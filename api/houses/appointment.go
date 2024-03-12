package houses

import (
	"net/http"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HousesAppointmentPost 用于实现用户预约房屋
func HousesAppointmentPost(c *gin.Context) {
	db, exists := c.MustGet("db").(*gorm.DB)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法获取数据库连接",
		})
		return
	}

	var reserve model.Reserve
	err := c.ShouldBind(&reserve)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	err = db.Create(&reserve).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "预约失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "预约成功",
		"url":     "/houses/appointment",
	})
}

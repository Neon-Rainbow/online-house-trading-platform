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

	// 从请求中获取预约信息,将预约信息绑定到model.Reserve结构体中
	var reserve model.Reserve
	err := c.ShouldBind(&reserve)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	// 将预约信息存入数据库
	err = db.Create(&reserve).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "预约失败",
		})
		return
	}

	// 如果预约成功,则返回预约成功的信息
	c.JSON(http.StatusOK, gin.H{
		"message": "预约成功",
		"url":     "/houses/appointment",
	})
}

package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserAppointmentGet 用于查看用户预约的房屋信息
func UserAppointmentGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/user/appointment",
		"method":  "GET",
		"user_id": c.Param("user_id"),
	})
}

// UserAppointmentPost 用于预约房屋
// 该接口似乎可以废弃,因为预约房屋的接口已经在houses包中实现
func UserAppointmentPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/user/appointment",
		"method":  "POST",
		"user_id": c.Param("user_id"),
	})
}

// UserAppointmentDelete 用于删除用户预约的房屋信息
func UserAppointmentDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/user/appointment",
		"method":  "DELETE",
		"user_id": c.Param("user_id"),
	})
}

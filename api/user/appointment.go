package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAppointmentGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/user/appointment",
		"method":  "GET",
	})
}

func UserAppointmentPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/user/appointment",
		"method":  "POST",
	})
}

func UserAppointmentDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/user/appointment",
		"method":  "DELETE",
	})
}

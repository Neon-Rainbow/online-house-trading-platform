package houses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HousesAppointmentPost 用于实现用户预约房屋
func HousesAppointmentPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/houses/appointment",
		"method":  "POST",
	})
}

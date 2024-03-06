package houses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HousesAppointmentGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/houses/appointment",
		"method":  "GET",
	})
}

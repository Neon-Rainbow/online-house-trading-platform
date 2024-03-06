package houses

import (
	"github.com/gin-gonic/gin"
)

func HouseByIDGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/houses/:id/",
		"method":  "GET",
		"id":      c.Param("id"),
	})
}

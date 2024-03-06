package houses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HouseByIDGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/houses/:id/",
		"method":  "GET",
		"id":      c.Param("id"),
	})
}

package controller

import (
	"os"

	"github.com/gin-gonic/gin"
)

func DeleteLogFile(c *gin.Context) {
	_ = os.Remove("./application.log")
	_ = os.Remove("./formatted_application.log")
	ResponseSuccess(c, nil)
	return
}

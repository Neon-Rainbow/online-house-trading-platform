package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LearnMoreGet(c *gin.Context){
	c.HTML(http.StatusOK, "learn_more.html", nil)
}
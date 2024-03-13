package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomePageGet 用来处理首页的Get请求
func HomePageGet(c *gin.Context) {
	c.HTML(http.StatusOK, "frontpage.html", nil)
}

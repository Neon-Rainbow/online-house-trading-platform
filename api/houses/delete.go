package houses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteGet 用于处理用户的删除房屋界面的GET请求 现已重定向至 /user/release
func DeleteGet(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "/houses/delete",
	//	"method":  "GET",
	//})
	c.Redirect(http.StatusMovedPermanently, "/user/release")
}

// DeletePost 用于处理用户的删除房屋界面的POST请求 现已重定向至 /user/release
func DeletePost(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "/houses/delete",
	//	"method":  "POST",
	//	"id":      c.Param("id"),
	//})
	c.Redirect(http.StatusMovedPermanently, "/user/release")
}

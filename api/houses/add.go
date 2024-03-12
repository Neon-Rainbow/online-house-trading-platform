package houses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddGet 用于处理用户的添加房屋界面的GET请求
// 业务逻辑修改,认为用户添加房屋应该在/user/release界面进行
// 接口现在已经弃用
func AddGet(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "/houses/add",
	//	"method":  "GET",
	//})
	//c.Redirect(http.StatusMovedPermanently, "/user/release")
}

// AddPost 用于处理用户的添加房屋界面的POST请求
// 业务逻辑修改,认为用户添加房屋应该在/user/release界面进行
// 接口现在已经弃用
func AddPost(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "/houses/add",
	//	"method":  "POST",
	//})
	c.Redirect(http.StatusMovedPermanently, "/user/release")
}

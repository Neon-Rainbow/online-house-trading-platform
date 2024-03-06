// Package profile 提供了用户个人信息,例如名字,年龄等的API
package profile

import (
	"github.com/gin-gonic/gin"
)

// ProfileGet 用于处理用户的个人信息界面的GET请求, 该界面需要登录后才能访问, 未登录用户将被重定向到登录界面,
// 该界面获取用户的个人信息, 包括用户名, 邮箱, 电话号码等
func ProfileGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/user/profile",
		"method":  "GET",
		"user_id": c.Param("user_id"),
	})
}

// ProfilePost 用于处理用户的个人信息界面的POST请求, 该界面需要登录后才能访问, 未登录用户将被重定向到登录界面,
// 发送该post请求后应该需要被redirect到/user/profile/update界面,该部分应该是前端来完成
func ProfilePost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/user/profile",
		"method":  "POST",
	})
}

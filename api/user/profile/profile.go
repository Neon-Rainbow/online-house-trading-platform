// Package profile 提供了用户个人信息,例如名字,年龄等的API
package profile

import (
	"log"
	"net/http"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProfileGet 用于处理用户的个人信息界面的GET请求, 该界面需要登录后才能访问, 未登录用户将被重定向到登录界面,
// 该界面获取用户的个人信息, 包括用户名, 邮箱, 电话号码等
func ProfileGet(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		c.JSON(500, gin.H{
			"error": "无法获取数据库连接",
		})
		log.Printf("无法获取数据库连接")
		return
	}

	var user model.User
	result := db.Preload("UserAvatar").First(&user) // 从数据库中查询用户信息

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "数据库查询出错",
		})
		log.Printf("查询用户信息时数据库查询出错, 错误原因为 %v", result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
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

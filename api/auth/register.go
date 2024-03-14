package auth

import (
	"log"
	"net/http"
	"online-house-trading-platform/pkg/model"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterGet 用于处理用户的注册界面的GET请求
// 返回状态码200和注册界面的信息
func RegisterGet(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "/auth/register",
	//	"method":  "GET",
	//})
	c.HTML(http.StatusOK, "register.html", nil)
}

// RegisterPost 用于处理用户的注册界面的POST请求
// 注册成功后返回状态码200和注册成功的信息
// 注册失败后返回如果时因为用户名或邮箱已存在则返回状态码400和错误信息,若是其他原因则返回状态码500和错误信息
// 格式为json,如{"error":"用户名已存在"}
func RegisterPost(c *gin.Context) {
	db, exists := c.MustGet("db").(*gorm.DB)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取数据库连接"})
		return
	}

	var user model.User
	err := c.ShouldBind(&user)

	user.Password = encryptPassword(user.Password) // 对密码进行加密

	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误"})
		return
	}
	if user.Username == "" || user.Password == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户名、密码和邮箱不能为空",
		})
		return
	}

	err = db.Create(&user).Error

	//判断注册信息是否合法,若非法则将错误信息返回给前端
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			if strings.Contains(err.Error(), "idx_users_username") {
				// 用户名重名,返回状态码400和错误信息
				c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
				return
			} else if strings.Contains(err.Error(), "idx_users_email") {
				// 邮箱已经被注册,返回状态码400和错误信息
				c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱已被注册"})
				return
			}
		}
		log.Printf("Error creating user: %v", err)
		// 其他错误,返回状态码500和错误信息
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}

	// 注册成功,返回状态码200和注册成功的信息
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"url":     "/auth/register",
	})
}

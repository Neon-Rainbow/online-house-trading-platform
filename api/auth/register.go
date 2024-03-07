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
func RegisterGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/auth/register",
		"method":  "GET",
	})
}

// RegisterPost 用于处理用户的注册界面的POST请求
func RegisterPost(c *gin.Context) {
	db, exists := c.MustGet("db").(*gorm.DB)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取数据库连接"})
		return
	}

	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Printf("Error binding data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}
	if user.Username == "" || user.Password == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户名、密码和邮箱不能为空",
		})
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			if strings.Contains(err.Error(), "idx_users_username") {
				c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
				return
			} else if strings.Contains(err.Error(), "idx_users_email") {
				c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱已被注册"})
				return
			}
		}
		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}

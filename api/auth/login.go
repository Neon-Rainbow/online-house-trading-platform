package auth

import (
	"errors"
	"log"
	"net/http"
	"online-house-trading-platform/pkg/jwt"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LoginGet 用于处理用户的登录界面的GET请求
// 返回状态码200和登录界面的信息
func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "/auth/login",
	//	"method":  "GET",
	//})
}

// LoginPost 用于处理用户的登录界面的POST请求
// 返回状态码200和登录成功的信息
func LoginPost(c *gin.Context) {
	db, exists := c.MustGet("db").(*gorm.DB)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取数据库连接"})
		return
	}

	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Printf("error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	var dbUser model.User
	result := db.Where("username = ?", user.Username).First(&dbUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "找不到用户",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "数据库查询出错",
			})
		}
		return
	}
	if encryptPassword(user.Password) != dbUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "密码错误",
		})
		return
	}
	token, err := jwt.GenerateToken(user.Username, dbUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法生成token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"data": gin.H{
			"token":    token,
			"user_id":  dbUser.ID,
			"username": dbUser.Username,
		},
	})
	log.Printf("用户登录成功: %v", user.Username)

}

package dao

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetDB 用于获取数据库连接
func GetDB(c *gin.Context) (*gorm.DB, error) {
	db, exists := c.MustGet("db").(*gorm.DB)
	if !exists {
		return nil, errors.New("无法获取数据库连接")
	}
	return db, nil
}

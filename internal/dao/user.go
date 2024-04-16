package dao

import (
	"online-house-trading-platform/pkg/model"

	"gorm.io/gorm"
)

// GetUserByUsername 用于根据用户名获取用户信息
func GetUserByUsername(db *gorm.DB, username string) (*model.User, error) {
	var user model.User
	result := db.Where("username = ?", username).First(&user)
	return &user, result.Error
}

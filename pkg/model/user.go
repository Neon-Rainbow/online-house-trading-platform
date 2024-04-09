package model

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex;not null;type:varchar(255)" form:"username" binding:"required"`
	Password string `json:"password" gorm:"not null;type:varchar(255)" form:"password" binding:"required"`
	Email    string `json:"email" gorm:"uniqueIndex;not null;type:varchar(255)" form:"email"`
	Role     string `json:"role" form:"role"`
}

// UserAvatar 用户头像
type UserAvatar struct {
	gorm.Model
	UserID uint   `gorm:"not null" json:"user_id"`
	URL    string `gorm:"not null" json:"url"`
}

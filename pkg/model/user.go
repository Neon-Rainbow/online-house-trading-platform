package model

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username string `json:"username", gorm:"uniqueIndex;not null"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

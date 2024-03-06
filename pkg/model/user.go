package model

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex;not null;type:varchar(255)"`
	Password string `json:"password" gorm:"not null;type:varchar(255)"`
	Email    string `json:"email" gorm:"uniqueIndex;not null;type:varchar(255)"`
	Role     string `json:"role"`
}

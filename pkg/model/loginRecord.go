package model

import (
	"gorm.io/gorm"
	"time"
)

type LoginRecord struct {
	gorm.Model
	UserId      uint      `json:"user_id" gorm:"not null"`
	LoginTime   time.Time `json:"login_time" gorm:"not null"`
	LoginIp     string    `json:"login_ip" gorm:"not null"`
	LoginMethod string    `json:"login_method" gorm:"not null"`
}

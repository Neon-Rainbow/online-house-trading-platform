package model

import (
	"gorm.io/gorm"
)

// Reserve 预定模型,用于存储用户预约看房的信息
// 需要存储预约看房的时间
type Reserve struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	HouseID uint   `gorm:"not null"`
	Time    string `gorm:"not null"`
}
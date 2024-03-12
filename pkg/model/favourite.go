package model

import "gorm.io/gorm"

// Favourite 收藏模型,用于存储用户收藏的房屋
type Favourite struct {
	gorm.Model
	UserID  uint `gorm:"not null" json:"user_id"`
	HouseID uint `gorm:"not null" json:"house_id"`
}

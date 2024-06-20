package model

import (
	"gorm.io/gorm"
)

// Favourite 收藏模型，用于存储用户收藏的房屋
type Favourite struct {
	gorm.Model
	UserID  uint `gorm:"not null;index:idx_user_house,unique" json:"user_id"`  // 作为唯一索引的一部分
	HouseID uint `gorm:"not null;index:idx_user_house,unique" json:"house_id"` // 作为唯一索引的一部分
}

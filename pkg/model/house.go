package model

import (
	"gorm.io/gorm"
)

// House 房屋模型
type House struct {
	gorm.Model
	Owner       string  `json:"name"`                     // 房屋所有者
	OwnerID     uint    `gorm:"not null" json:"owner_id"` // 房屋所有者ID
	Title       string  `gorm:"not null" json:"title"`    // 房屋标题
	Description string  `json:"description"`              // 房屋描述
	Price       float64 `gorm:"not null" json:"price"`    // 房屋价格
	Address     string  `gorm:"not null" json:"address"`  // 房屋地址
}

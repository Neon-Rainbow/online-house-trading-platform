package model

import (
	"gorm.io/gorm"
)

// House 定义了房屋的模型。
type House struct {
	gorm.Model
	Owner            string       `json:"owner"`                    // Owner 表示房屋的所有者。
	OwnerID          uint         `gorm:"not null" json:"owner_id"` // OwnerID 表示房屋所有者的ID，不能为空。
	Title            string       `gorm:"not null" json:"title"`    // Title 是房屋的标题，不能为空。
	Description      string       `json:"description"`              // Description 描述了房屋的详细信息。
	Price            float64      `gorm:"not null" json:"price"`    // Price 表示房屋的价格，不能为空。
	Address          string       `gorm:"not null" json:"address"`  // Address 是房屋的地址，不能为空。
	Images           []HouseImage `gorm:"foreignKey:HouseID"`       // Images 存储与此房屋相关联的图片。
	HouseOrientation string       `json:"house_orientation"`        // HouseOrientation 房屋朝向
	Layout           string       `json:"layout"`                   // Layout 房屋户型
	Area             float64      `json:"area"`                     // Area 房屋面积
	Floor            string       `json:"floor"`                    // Floor 房屋楼层
	RentPrice        float64      `json:"rent_price"`               // RentPrice 房屋租金
	Type             string       `json:"type"`                     // Type 房屋类型
	PostCode         string       `json:"post_code"`                // PostCode 房屋邮编
}

// HouseImage 定义了房屋图片的模型。
type HouseImage struct {
	gorm.Model
	HouseID uint   `gorm:"not null"` // HouseID 是关联的房屋ID。
	URL     string `gorm:"not null"` // URL 是图片的存储位置。
}

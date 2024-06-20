package model

import "gorm.io/gorm"

type ViewingRecords struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	HouseID uint `json:"house_id"`
}

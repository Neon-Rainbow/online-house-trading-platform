package dao

import (
	"online-house-trading-platform/pkg/model"

	"gorm.io/gorm"
)

func DeleteAccount(db *gorm.DB, id uint) error {
	result := db.Where("id = ?", id).Preload("Avatar").Delete(&model.User{})
	return result.Error
}

package dao

import (
	"gorm.io/gorm"
	"online-house-trading-platform/pkg/model"
)

func DeleteAccount(db *gorm.DB, id uint) error {
	result := db.Delete(&model.User{}).Where("id = ?", id)
	return result.Error
}

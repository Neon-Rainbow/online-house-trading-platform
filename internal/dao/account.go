package dao

import (
	"gorm.io/gorm"
	"online-house-trading-platform/pkg/model"
)

func DeleteAccount(db *gorm.DB, id uint) error {
	return db.Delete(&model.User{}, id).Error
}

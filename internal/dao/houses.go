package dao

import (
	"online-house-trading-platform/pkg/model"

	"gorm.io/gorm"
)

// CreateAppointment 用于实现用户预约房屋
func CreateAppointment(db *gorm.DB, reserve *model.Reserve) error {
	return db.Create(reserve).Error
}

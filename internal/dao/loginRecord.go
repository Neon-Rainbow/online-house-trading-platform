package dao

import (
	"gorm.io/gorm"
	"online-house-trading-platform/pkg/model"
)

func CreateLoginRecord(db *gorm.DB, record *model.LoginRecord) error {
	result := db.Create(record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetLoginRecord(db *gorm.DB, id string) (*[]model.LoginRecord, error) {
	var records []model.LoginRecord
	result := db.Where("user_id = ?", id).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	return &records, nil
}

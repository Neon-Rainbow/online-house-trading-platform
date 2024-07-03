package dao

import (
	"gorm.io/gorm"
	"online-house-trading-platform/pkg/database"
	"online-house-trading-platform/pkg/model"
)

func CreateLoginRecord(record *model.LoginRecord) error {
	db := database.Database
	result := db.Create(record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetLoginRecord(id string, pageSize int, pageNum int) ([]model.LoginRecord, error) {
	db := database.Database
	var records []model.LoginRecord
	var result *gorm.DB
	if pageSize == 0 && pageNum == 0 {
		result = db.Where("user_id = ?", id).Find(&records)
	} else {
		result = db.Where("user_id = ?", id).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&records)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}

func GetAllLoginRecords(pageSize int, pageNum int) ([]model.LoginRecord, error) {
	db := database.Database
	var records []model.LoginRecord
	var result *gorm.DB
	if pageSize == 0 && pageNum == 0 {
		result = db.Find(&records)
	} else {
		result = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&records)

	}
	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}

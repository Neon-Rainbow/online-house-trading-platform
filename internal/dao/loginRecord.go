package dao

import (
	"online-house-trading-platform/pkg/database"
	"online-house-trading-platform/pkg/model"

	"gorm.io/gorm"
)

func CreateLoginRecord(record *model.LoginRecord) error {
	db := database.Database
	result := db.Create(record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetLoginRecord 用于获取用户的登录记录
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

func GetUserRecentlyLoginRecord(id string) (*model.LoginRecord, error) {
	db := database.Database
	var record model.LoginRecord
	result := db.Where("user_id = ?", id).Limit(1).Order("update_at desc").First(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return &record, nil
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

package dao

import (
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

func GetLoginRecord(id string) ([]model.LoginRecord, error) {
	db := database.Database
	var records []model.LoginRecord
	result := db.Where("user_id = ?", id).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}

func GetAllLoginRecords() ([]model.LoginRecord, error) {
	db := database.Database
	var records []model.LoginRecord
	result := db.Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}

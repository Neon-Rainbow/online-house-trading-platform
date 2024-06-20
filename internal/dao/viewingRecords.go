package dao

import (
	"gorm.io/gorm"
	"online-house-trading-platform/pkg/model"
)

// GetViewingRecordsByUserID 根据用户ID获取用户看房记录
func GetViewingRecordsByUserID(db *gorm.DB, idUint uint, pageSize int, pageNum int) ([]model.ViewingRecords, error) {
	var viewingRecords []model.ViewingRecords
	err := db.Where("user_id = ?", idUint).Order("updated_at DESC").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&viewingRecords).Error
	if err != nil {
		return nil, err
	}
	return viewingRecords, nil
}

// AddViewingRecords 添加用户看房记录
func AddViewingRecords(db *gorm.DB, viewingRecords *model.ViewingRecords) error {
	err := db.Create(viewingRecords).Error
	if err != nil {
		return err
	}
	return nil
}

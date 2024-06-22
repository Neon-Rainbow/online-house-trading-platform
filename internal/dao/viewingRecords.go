package dao

import (
	"errors"
	"gorm.io/gorm"
	"online-house-trading-platform/pkg/model"
)

// GetViewingRecordsByUserID 根据用户ID获取用户看房记录
func GetViewingRecordsByUserID(db *gorm.DB, idUint uint, pageSize int, pageNum int) ([]model.ViewingRecords, int64, error) {
	var viewingRecords []model.ViewingRecords
	var totalRecords int64
	err := db.Model(&model.ViewingRecords{}).Where("user_id = ?", idUint).Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Where("user_id = ?", idUint).Order("updated_at DESC").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&viewingRecords).Error
	if err != nil {
		return nil, totalRecords, err
	}
	return viewingRecords, totalRecords, nil
}

// AddViewingRecords 添加用户看房记录
func AddViewingRecords(db *gorm.DB, viewingRecords *model.ViewingRecords) error {
	// 检查是否存在相同的看房记录
	var existingRecord model.ViewingRecords
	err := db.Where("user_id = ? AND house_id = ?", viewingRecords.UserID, viewingRecords.HouseID).First(&existingRecord).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果记录不存在，则插入新记录
		if err := db.Create(viewingRecords).Error; err != nil {
			return err
		}
	} else {
		// 如果记录存在，则更新记录的时间戳
		existingRecord.UpdatedAt = viewingRecords.UpdatedAt
		if err := db.Save(&existingRecord).Error; err != nil {
			return err
		}
	}

	return nil
}

package dao

import (
	"online-house-trading-platform/pkg/model"

	"gorm.io/gorm"
)

// CreateAppointment 用于实现用户预约房屋
func CreateAppointment(db *gorm.DB, reserve *model.Reserve) error {
	return db.Create(reserve).Error
}

// CreateFavorite 用于实现用户收藏房屋
func CreateFavorite(db *gorm.DB, favorite *model.Favourite) error {
	return db.Create(favorite).Error
}

// GetAllHouseInformation 用于获取数据库中的所有房屋信息
func GetAllHouseInformation(db *gorm.DB) ([]model.House, error) {
	var houses []model.House
	result := db.Preload("Images").Find(&houses)
	if result.Error != nil {
		return nil, result.Error
	}
	return houses, nil
}

// GetHouseInformationByID 用于获取数据库中指定ID的房屋信息
func GetHouseInformationByID(db *gorm.DB, houseID uint) (*model.House, error) {
	var house *model.House
	result := db.Preload("Images").First(&house, houseID)
	if result.Error != nil {
		return house, result.Error
	}
	return house, nil
}

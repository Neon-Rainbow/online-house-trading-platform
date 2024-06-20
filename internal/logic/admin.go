package logic

import (
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"

	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) (*[]model.User, error) {
	users, err := dao.GetAllUsers(db)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetAllHouses(db *gorm.DB) (*[]model.House, error) {
	houses, err := dao.GetAllHouses(db)
	if err != nil {
		return nil, err
	}
	return houses, nil
}

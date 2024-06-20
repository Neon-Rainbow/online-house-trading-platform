package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"

	"gorm.io/gorm"
)

// FetchAllHouses 用于获取所有房屋信息
func FetchAllHouses(db *gorm.DB) ([]model.House, *model.Error) {
	houses, err := dao.GetAllHouseInformation(db)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetHouseListError}
	}
	return houses, nil
}

// FetchCertainHouseInformationByID 用于获取指定ID的房屋信息
func FetchCertainHouseInformationByID(db *gorm.DB, houseID uint, userID uint) (*model.House, *model.Error) {
	house, err := dao.GetHouseInformationByID(db, houseID)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetHouseInfoError}
	}
	record := &model.ViewingRecords{
		HouseID: houseID,
		UserID:  userID,
	}
	err = dao.AddViewingRecords(db, record)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.AddViewingRecordsError}
	}
	return house, nil
}

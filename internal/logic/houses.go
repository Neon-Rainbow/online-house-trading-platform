package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
)

// FetchAllHouses 用于获取所有房屋信息
func FetchAllHouses() ([]model.House, *model.Error) {
	houses, err := dao.GetAllHouseInformation()
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetHouseListError}
	}
	return houses, nil
}

// FetchCertainHouseInformationByID 用于获取指定ID的房屋信息
func FetchCertainHouseInformationByID(houseID uint, userID uint) (*model.House, *model.Error) {
	house, err := dao.GetHouseInformationByID(houseID)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetHouseInfoError}
	}
	record := &model.ViewingRecords{
		HouseID: houseID,
		UserID:  userID,
	}
	err = dao.AddViewingRecords(record)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.AddViewingRecordsError}
	}
	return house, nil
}

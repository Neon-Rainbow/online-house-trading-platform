package logic

import (
	"online-house-trading-platform/internal/controller"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
)

// FetchAllHouses 用于获取所有房屋信息
func FetchAllHouses(c *gin.Context) ([]model.House, *model.Error) {
	db, err := dao.GetDB(c)
	if err != nil {
		return nil, &model.Error{StatusCode: controller.GetDBError}
	}

	houses, err := dao.GetAllHouseInformation(db)
	if err != nil {
		return nil, &model.Error{StatusCode: controller.GetHouseListError}
	}
	return houses, nil
}

// FetchCertainHouseInformationByID 用于获取指定ID的房屋信息
func FetchCertainHouseInformationByID(c *gin.Context, houseID uint) (*model.House, *model.Error) {
	db, err := dao.GetDB(c)
	if err != nil {
		return nil, &model.Error{StatusCode: controller.GetDBError}
	}

	house, err := dao.GetHouseInformationByID(db, houseID)
	if err != nil {
		return nil, &model.Error{StatusCode: controller.GetHouseInfoError}
	}
	return house, nil

}

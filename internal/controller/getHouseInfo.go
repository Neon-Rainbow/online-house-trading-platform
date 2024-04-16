package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllHouses 用于获取所有房屋的信息
func GetAllHouses(c *gin.Context) {
	houses, err := logic.FetchAllHouses(c)
	if err != nil {
		ResponseError(c, *err)
	}
	ResponseSuccess(c, houses)
}

// GetHouseInfoByID 用于获取某个房屋的详细信息
func GetHouseInfoByID(c *gin.Context) {
	houseID := c.Param("house_id")
	houseIDUint, err := strconv.ParseUint(houseID, 10, 64)
	if err != nil {
		ResponseErrorWithCode(c, codes.HouseIDInvalid)
	}
	house, apiError := logic.FetchCertainHouseInformationByID(c, uint(houseIDUint))
	if apiError != nil {
		ResponseError(c, *apiError)
	}
	ResponseSuccess(c, house)
}

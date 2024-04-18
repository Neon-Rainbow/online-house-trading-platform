package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllHouses 用于获取所有房屋的信息
// @Summary 获取所有房屋信息
// @Description 获取所有房屋信息
// @Tags 房屋
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} controller.ResponseData "获取成功"
// @Failure 400 {object} controller.ResponseData "预约失败,具体原因查看json中的message字段和code字段"
// @Router /house [get]
func GetAllHouses(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	houses, err := logic.FetchAllHouses(db)
	if err != nil {
		ResponseError(c, *err)
		return
	}
	ResponseSuccess(c, houses)
	return
}

// GetHouseInfoByID 用于获取某个房屋的详细信息
// @Summary 获取某个房屋的详细信息
// @Description 获取某个房屋的详细信息
// @Tags 房屋
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param house_id path string true "房屋ID"
// @Success 200 {object} controller.ResponseData "获取成功"
// @failure 200 {object} controller.ResponseData "获取失败"
// @Router /house/{house_id} [get]
func GetHouseInfoByID(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		ResponseErrorWithCode(c, codes.GetDBError)
	}

	houseID := c.Param("house_id")
	houseIDUint, err := strconv.ParseUint(houseID, 10, 64)
	if err != nil {
		ResponseErrorWithCode(c, codes.HouseIDInvalid)
	}
	house, apiError := logic.FetchCertainHouseInformationByID(db, uint(houseIDUint))
	if apiError != nil {
		ResponseError(c, *apiError)
	}
	ResponseSuccess(c, house)
}

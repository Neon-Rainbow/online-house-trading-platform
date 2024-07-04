package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
// @Router /houses [get]
func GetAllHouses(c *gin.Context) {
	houses, err := logic.FetchAllHouses()
	if err != nil {
		zap.L().Error("GetAllHouses: logic.FetchAllHouses failed",
			zap.String("错误码", strconv.FormatInt(int64(err.StatusCode), 10)),
		)
		ResponseError(c, *err)
		return
	}
	ResponseSuccess(c, houses)
	return
}

// GetHouseInformationByHouseID 用于获取某个房屋的详细信息
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
func GetHouseInformationByHouseID(c *gin.Context) {
	houseID := c.Param("house_id")
	houseIDUint, err := strconv.ParseUint(houseID, 10, 64)
	if err != nil {
		zap.L().Error("GetHouseInformationByHouseID: strconv.ParseUint failed",
			zap.String("错误码", strconv.FormatInt(int64(codes.HouseIDInvalid), 10)),
			zap.String("house_id", houseID),
		)
		ResponseErrorWithCode(c, codes.HouseIDInvalid)
		return
	}
	_userid, _ := c.Get("user_id")
	userID := _userid.(uint)

	house, apiError := logic.FetchCertainHouseInformationByID(uint(houseIDUint), userID)
	if apiError != nil {
		zap.L().Error("GetHouseInformationByHouseID: logic.FetchCertainHouseInformationByID failed",
			zap.String("错误码", strconv.FormatInt(int64(apiError.StatusCode), 10)),
		)
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, house)
}

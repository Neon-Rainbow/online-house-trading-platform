package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HousesAppointmentPost 用于处理用户预约房屋的POST请求
// @Summary 预约房屋
// @Description 用户预约房屋
// @Tags 预约
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param house_id body uint true "房屋ID"
// @Param time body time true "预约时间"
// @Success 200 {object} controller.ResponseData "预约成功"
// @Failure 400 {object} controller.ResponseData "预约失败,具体原因查看json中的message字段和code字段"
// @Router /houses/appointment [post]
func HousesAppointmentPost(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		ResponseErrorWithCode(c, codes.GetUserIDError)
		return
	}
	userIDUint, ok := userID.(uint)
	if !ok {
		ResponseErrorWithCode(c, codes.UserIDTypeError)
		return
	}

	var reserve model.Reserve
	err := c.ShouldBind(&reserve)
	if err != nil {
		ResponseErrorWithCode(c, codes.ReserveInvalidParam)
		return
	}
	apiError := logic.AppointmentHandle(db, &reserve, userIDUint)

	if apiError != nil {
		ResponseError(c, *apiError)
		return
	}

	ResponseSuccess(c, nil)
	return
}

// HousesAppointmentGet 用于获取用户预约的房屋
// @Summary 获取用户预约的房屋
// @Description 获取用户预约的房屋
// @Tags 预约
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param user_id path uint true "用户ID"
// @Success 200 {object} controller.ResponseData "获取成功"
// @Failure 400 {object} controller.ResponseData "获取失败"
// @Router /user/{user_id}/appointment [get]
func HousesAppointmentGet(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		ResponseErrorWithCode(c, codes.GetUserIDError)
		return
	}
	userIDUint, ok := userID.(uint)
	if !ok {
		ResponseErrorWithCode(c, codes.UserIDTypeError)
		return
	}
	reserve, apiError := logic.GetReserve(db, userIDUint)
	if apiError != nil {
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, reserve)
	return
}

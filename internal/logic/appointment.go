package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
)

// AppointmentHandle 用于处理用户预约房屋的请求\
// @title AppointmentHandle
// @description 处理用户预约房屋的请求
// @param reserve *model.Reserve 预约信息
// @param userID uint 用户ID
// @return apiError *model.Error 错误信息
func AppointmentHandle(reserve *model.Reserve, userID uint) (apiError *model.Error) {
	if reserve.Time.IsZero() || reserve.HouseID == 0 {
		return &model.Error{StatusCode: codes.ReserveInvalidParam}
	}
	reserve.UserID = userID
	err := dao.CreateAppointment(reserve)
	if err != nil {
		return &model.Error{StatusCode: codes.ReserveError}
	}
	return nil
}

// GetReserve 用于获取用户的预约信息
// @title GetReserve
// @description 获取用户的预约信息
// @param userID uint 用户ID
// @return reserve []model.Reserve 预约信息
func GetReserve(userId uint) ([]model.Reserve, *model.Error) {
	reserve, err := dao.GetReserve(userId)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetReserveInformationError}
	}
	return reserve, nil
}

// GetAllReserve 用于获取所有用户的预约信息
// @title GetAllReserve
// @description 获取所有用户的预约信息
// @return reserve []model.Reserve 预约信息
func GetAllReserve() ([]model.Reserve, *model.Error) {
	reserve, err := dao.GetAllReserve()
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetReserveInformationError}
	}
	return reserve, nil
}

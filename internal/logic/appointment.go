package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
)

// AppointmentHandle 用于处理用户预约房屋的请求
func AppointmentHandle(reserve *model.Reserve, userID uint) *model.Error {
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
func GetReserve(idUint uint) ([]model.Reserve, *model.Error) {
	reserve, err := dao.GetReserve(idUint)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetReserveInformationError}
	}
	return reserve, nil
}

// GetAllReserve 用于获取所有用户的预约信息
func GetAllReserve() (*[]model.Reserve, *model.Error) {
	reserve, err := dao.GetAllReserve()
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetReserveInformationError}
	}
	return reserve, nil
}

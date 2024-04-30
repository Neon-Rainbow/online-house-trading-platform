package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"

	"gorm.io/gorm"
)

// AppointmentHandle 用于处理用户预约房屋的请求
func AppointmentHandle(db *gorm.DB, reserve *model.Reserve, userID uint) *model.Error {
	if reserve.Time.IsZero() || reserve.HouseID == 0 {
		return &model.Error{StatusCode: codes.ReserveInvalidParam}
	}
	reserve.UserID = userID
	err := dao.CreateAppointment(db, reserve)
	if err != nil {
		return &model.Error{StatusCode: codes.ReserveError}
	}
	return nil
}

// GetReserve 用于获取用户的预约信息
func GetReserve(db *gorm.DB, idUint uint) ([]model.Reserve, *model.Error) {
	reserve, err := dao.GetReserve(db, idUint)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetReserveInformetionError}
	}
	return reserve, nil
}

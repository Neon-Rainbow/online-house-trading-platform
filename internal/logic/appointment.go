package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"

	"gorm.io/gorm"
)

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

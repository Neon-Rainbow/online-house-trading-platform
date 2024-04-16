package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
)

func AppointmentHandle(c *gin.Context, reserve *model.Reserve, userID uint) *model.Error {
	db, err := dao.GetDB(c)
	if err != nil {
		return &model.Error{StatusCode: codes.GetDBError}
	}

	if reserve.Time.IsZero() || reserve.HouseID == 0 {
		return &model.Error{StatusCode: codes.ReserveInvalidParam}
	}
	reserve.UserID = userID
	err = dao.CreateAppointment(db, reserve)
	if err != nil {
		return &model.Error{StatusCode: codes.ReserveError}
	}
	return nil
}

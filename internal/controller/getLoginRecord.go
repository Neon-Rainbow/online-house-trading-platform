package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetUserLoginRecordByUserID(c *gin.Context) {
	userID := c.Param("user_id")

	loginRecords, err := dao.GetLoginRecord(userID)
	if err != nil {
		zap.L().Error("GetUserLoginRecordByUserID: dao.GetUserLoginRecordByUserID failed",
			zap.Error(err),
			zap.String("user_id", userID),
		)
		ResponseErrorWithCode(c, codes.LoginServerBusy)
		return
	}

	ResponseSuccess(c, loginRecords)
}

func GetAllLoginRecords(c *gin.Context) {
	loginRecords, err := dao.GetAllLoginRecords()
	if err != nil {
		zap.L().Error("GetAllLoginRecords: dao.GetAllLoginRecords failed",
			zap.Error(err),
		)
		ResponseErrorWithCode(c, codes.LoginServerBusy)
		return
	}

	ResponseSuccess(c, loginRecords)
}

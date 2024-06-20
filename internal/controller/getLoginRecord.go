package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetUserLoginRecordByUserID(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		zap.L().Error("LoginPost: c.MustGet(\"db\").(*gorm.DB) failed")
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	userID := c.Param("user_id")

	loginRecords, err := dao.GetLoginRecord(db, userID)
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
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		zap.L().Error("LoginPost: c.MustGet(\"db\").(*gorm.DB) failed")
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	loginRecords, err := dao.GetAllLoginRecords(db)
	if err != nil {
		zap.L().Error("GetAllLoginRecords: dao.GetAllLoginRecords failed",
			zap.Error(err),
		)
		ResponseErrorWithCode(c, codes.LoginServerBusy)
		return
	}

	ResponseSuccess(c, loginRecords)
}

package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
)

func GetLoginRecord(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		zap.L().Error("LoginPost: c.MustGet(\"db\").(*gorm.DB) failed")
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	userID := c.Param("user_id")

	loginRecords, err := dao.GetLoginRecord(db, userID)
	if err != nil {
		zap.L().Error("GetLoginRecord: dao.GetLoginRecord failed",
			zap.Error(err),
			zap.String("user_id", userID),
		)
		ResponseErrorWithCode(c, codes.LoginServerBusy)
		return
	}

	ResponseSuccess(c, loginRecords)
}

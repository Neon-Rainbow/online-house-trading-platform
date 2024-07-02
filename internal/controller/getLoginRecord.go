package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"strconv"
)

// GetUserLoginRecordByUserID 用于获取用户登录记录
// @Summary 获取用户登录记录
// Router /user/{user_id}/login_records [get]
func GetUserLoginRecordByUserID(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	pageNum, _ := strconv.Atoi(c.Query("page_num"))

	userID := c.Param("user_id")

	loginRecords, err := dao.GetLoginRecord(userID, pageSize, pageNum)
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

// GetAllLoginRecords 用于获取所有登录记录
// @Summary 获取所有登录记录
func GetAllLoginRecords(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	pageNum, _ := strconv.Atoi(c.Query("page_num"))

	loginRecords, err := dao.GetAllLoginRecords(pageSize, pageNum)
	if err != nil {
		zap.L().Error("GetAllLoginRecords: dao.GetAllLoginRecords failed",
			zap.Error(err),
		)
		ResponseErrorWithCode(c, codes.LoginServerBusy)
		return
	}

	ResponseSuccess(c, loginRecords)
}

package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"strconv"
)

func GetUserViewingRecordsByUserID(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		zap.L().Error("HousesAppointmentPost: c.MustGet(\"db\").(*gorm.DB) failed",
			zap.Int("错误码", codes.GetDBError.Int()),
		)
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	userID := c.Param("user_id")
	userIDUint64, _ := strconv.ParseUint(userID, 10, 32)
	userIDUint := uint(userIDUint64)

	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	pageNum, _ := strconv.Atoi(c.Query("page_num"))

	viewingRecords, apiError := logic.GetViewingRecords(db, userIDUint, pageSize, pageNum)
	if apiError != nil {
		ResponseErrorWithCode(c, apiError.StatusCode)
		return
	}
	ResponseSuccess(c, viewingRecords)

}

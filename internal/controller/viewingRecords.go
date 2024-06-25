package controller

import (
	"online-house-trading-platform/internal/logic"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserViewingRecordsByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	userIDUint64, _ := strconv.ParseUint(userID, 10, 32)
	userIDUint := uint(userIDUint64)

	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	pageNum, _ := strconv.Atoi(c.Query("page_num"))

	viewingRecords, totalCounts, apiError := logic.GetViewingRecords(userIDUint, pageSize, pageNum)
	if apiError != nil {
		ResponseErrorWithCode(c, apiError.StatusCode)
		return
	}

	response := gin.H{
		"viewing_records": viewingRecords,
		"total_size":      totalCounts,
	}

	ResponseSuccess(c, response)

}

package controller

import (
	"context"
	"errors"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resultChannel := make(chan []model.LoginRecord, 1)
	errorChannel := make(chan error, 1)

	go func() {
		loginRecords, err := dao.GetAllLoginRecords(pageSize, pageNum)
		if err != nil {
			errorChannel <- err
			return
		}
		resultChannel <- loginRecords
		return
	}()

	//loginRecords, err := dao.GetAllLoginRecords(pageSize, pageNum)
	//if err != nil {
	//	zap.L().Error("GetAllLoginRecords: dao.GetAllLoginRecords failed",
	//		zap.Error(err),
	//	)
	//	ResponseErrorWithCode(c, codes.LoginServerBusy)
	//	return
	//}

	select {
	case loginRecords := <-resultChannel:
		ResponseSuccess(c, loginRecords)
		return
	case err := <-errorChannel:
		zap.L().Error("GetAllLoginRecords: dao.GetAllLoginRecords failed",
			zap.Error(err))
		ResponseErrorWithCode(c, codes.LoginServerBusy)
		return
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			zap.L().Error("GetAllLoginRecords: ctx deadline exceeded")
			ResponseTimeout(c)
			return
		} else {
			zap.L().Error("GetAllLoginRecords: ctx error", zap.Error(ctx.Err()))
			ResponseErrorWithCode(c, codes.LoginServerBusy)
			return
		}
	}

	return
}

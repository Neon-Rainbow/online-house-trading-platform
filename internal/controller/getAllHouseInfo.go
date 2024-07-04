package controller

import (
	"context"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetAllHousesInformation 用于处理管理员获取所有房屋信息的Get请求
func GetAllHousesInformation(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resultChannel := make(chan []model.House, 1)
	errorChannel := make(chan *model.Error, 1)

	go func() {
		defer close(resultChannel)
		defer close(errorChannel)
		houses, err := logic.GetAllHouses()
		if err != nil {
			errorChannel <- err
			return
		}
		resultChannel <- houses
		return
	}()

	select {
	case houses := <-resultChannel:
		ResponseSuccess(c, houses)
		return
	case err := <-errorChannel:
		zap.L().Error("GetAllHousesInformation: logic.GetAllHouses failed",
			zap.Int("错误码", err.StatusCode.Int()),
		)
		ResponseError(c, *err)
		return
	case <-ctx.Done():
		zap.L().Error("GetAllHousesInformation: logic.GetAllHouses timeout",
			zap.Int("错误码", codes.RequestTimeOut.Int()))
		ResponseTimeout(c)
		return
	default:
		zap.L().Error("GetAllHousesInformation: logic.GetAllHouses failed",
			zap.Int("错误码", codes.LoginServerBusy.Int()),
		)
		ResponseErrorWithCode(c, codes.LoginServerBusy)
		return
	}

	//houses, err := logic.GetAllHouses()
	//if err != nil {
	//	zap.L().Error("houses, err := logic.GetAllHouses(db) failed")
	//}
	//ResponseSuccess(c, houses)
}

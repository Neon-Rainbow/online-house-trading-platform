package controller

import (
	"online-house-trading-platform/internal/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetAllHousesInformation 用于处理管理员获取所有房屋信息的Get请求
func GetAllHousesInformation(c *gin.Context) {
	houses, err := logic.GetAllHouses()
	if err != nil {
		zap.L().Error("houses, err := logic.GetAllHouses(db) failed")
	}
	ResponseSuccess(c, houses)
}

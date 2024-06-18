package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GetAllHousesInformation 用于处理管理员获取所有房屋信息的Get请求
func GetAllHousesInformation(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		zap.L().Error("RegisterPost: c.MustGet(\"db\").(*gorm.DB) failed",
			zap.String("错误码", strconv.FormatInt(int64(codes.GetDBError), 10)),
		)
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}
	houses, err := logic.GetAllHouses(db)
	if err != nil {
		zap.L().Error("houses, err := logic.GetAllHouses(db) failed")
	}
	ResponseSuccess(c, houses)
}

package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"strconv"
)

type deleteRequest struct {
	Id uint `json:"id"`
}

func DeleteAccount(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		zap.L().Error("RegisterPost: c.MustGet(\"db\").(*gorm.DB) failed",
			zap.String("错误码", strconv.FormatInt(int64(codes.GetDBError), 10)),
		)
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	var deleteReq deleteRequest
	err := c.ShouldBind(&deleteReq)
	if err != nil {
		zap.L().Error("RegisterPost: c.ShouldBind(&deleteReq) failed",
			zap.Int("错误码", codes.RegisterInvalidParam.Int()),
		)
		ResponseErrorWithCode(c, codes.RegisterInvalidParam)
		return
	}

	apiError := logic.DeleteAccountHandle(db, deleteReq.Id)
	if apiError != nil {
		zap.L().Error("RegisterPost: logic.RegisterHandle failed",
			zap.Int("错误码", apiError.StatusCode.Int()),
		)
		ResponseErrorWithCode(c, apiError.StatusCode)
		return
	}
}

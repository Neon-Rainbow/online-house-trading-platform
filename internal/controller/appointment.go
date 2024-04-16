package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HousesAppointmentPost 用于处理用户预约房屋的GET请求
func HousesAppointmentPost(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		ResponseErrorWithCode(c, codes.GetDBError)
	}

	userID, exists := c.Get("user_id")
	if !exists {
		ResponseErrorWithCode(c, codes.GetUserIDError)
	}
	userIDUint, ok := userID.(uint)
	if !ok {
		ResponseErrorWithCode(c, codes.UserIDTypeError)
	}

	var reserve model.Reserve
	err := c.ShouldBind(&reserve)
	if err != nil {
		ResponseErrorWithCode(c, codes.ReserveInvalidParam)
	}
	apiError := logic.AppointmentHandle(db, &reserve, userIDUint)

	if apiError != nil {
		ResponseError(c, *apiError)
	}

	ResponseSuccess(c, nil)
}

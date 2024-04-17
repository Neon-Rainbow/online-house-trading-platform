package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HousesAppointmentPost 用于处理用户预约房屋的GET请求
// @Summary 预约房屋
// @Description 用户预约房屋
// @Tags 预约
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param house_id body uint true "房屋ID"
// @Param time body time true "预约时间"
// @Success 200 {object} controller.ResponseData "预约成功"
// @Failure 400 {object} controller.ResponseData "预约失败,具体原因查看json中的message字段和code字段"
// @Router /appointment [post]
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

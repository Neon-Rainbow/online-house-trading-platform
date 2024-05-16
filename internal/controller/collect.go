package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// CollectPost 用于处理用户收藏房屋的Post请求
// @Summary 收藏房屋
// @Description 用户收藏房屋
// @Tags 收藏
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param house_id body uint true "房屋ID"
// @Success 200 {object} controller.ResponseData "收藏成功"
// @Failure 400 {object} controller.ResponseData "预约失败,具体原因查看json中的message字段和code字段"
// @Router /houses/collect [post]
func CollectPost(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		ResponseErrorWithCode(c, codes.GetUserIDError)
		return
	}
	userIDUint, ok := userID.(uint)
	if !ok {
		ResponseErrorWithCode(c, codes.UserIDTypeError)
		return
	}
	var favourite model.Favourite
	err := c.ShouldBind(&favourite)
	if err != nil {
		ResponseErrorWithCode(c, codes.ReserveInvalidParam)
		return
	}
	apiError := logic.CollectHandle(db, &favourite, userIDUint)
	if apiError != nil {
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, nil)
	return
}

// GetUserFavourites 用于获取用户收藏的房屋
// @Summary 获取用户收藏的房屋
// @Description 获取用户收藏的房屋
// @Tags 收藏
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param user_id path uint true "用户ID"
// @Success 200 {object} controller.ResponseData "获取成功"
// @Failure 400 {object} controller.ResponseData "预约失败,具体原因查看json中的message字段和code字段"
// @Router /user/{user_id}/favourites [post]
func GetUserFavourites(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		zap.L().Error("GetUserFavourites: c.MustGet(\"db\").(*gorm.DB) failed",
			zap.Int("错误码", codes.GetDBError.Int()),
		)
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		zap.L().Error("GetUserFavourites: c.Get(\"user_id\") failed",
			zap.Int("错误码", codes.GetUserIDError.Int()),
		)
		ResponseErrorWithCode(c, codes.GetUserIDError)
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		zap.L().Error("GetUserFavourites: userID.(uint) failed",
			zap.Int("错误码", codes.UserIDTypeError.Int()),
			zap.Any("用户ID", userID),
		)
		ResponseErrorWithCode(c, codes.UserIDTypeError)
		return
	}

	favourites, apiError := logic.GetUserFavourites(db, userIDUint)
	if apiError != nil {
		zap.L().Error("GetUserFavourites: logic.GetUserFavourites failed",
			zap.Int("错误码", apiError.StatusCode.Int()),
			zap.Int("用户ID", int(userIDUint)),
		)
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, favourites)
	return
}

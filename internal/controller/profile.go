package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProfileGet 用于处理用户获取个人信息的Get请求
// @Summary 获取个人信息
// @Description 获取个人信息
// @Tags 个人信息
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param user_id path string true "用户ID"
// @Success 200 {object} controller.ResponseData "获取成功"
// @Failure 400 {object} controller.ResponseData "获取失败"
// @Router /profile/{user_id} [get]
func ProfileGet(c *gin.Context) {
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
	userProfile, apiError := logic.GetUserProfile(db, userIDUint)
	if apiError != nil {
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, userProfile)
}

// ProfilePut 用于处理用户修改个人信息的Put请求
// @Summary 修改个人信息
// @Description 修改个人信息
// @Tags 个人信息
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object body model.RegisterRequest true "用户信息"
// @Success 200 {object} controller.ResponseData "修改成功"
// @Failure 400 {object} controller.ResponseData "修改失败"
// @Router /profile/profile [put]
func ProfilePut(c *gin.Context) {
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

	var userProfile model.User
	err := c.ShouldBind(&userProfile)
	if err != nil {
		ResponseErrorWithCode(c, codes.BindDataError)
		return
	}

	apiError := logic.ModifyUserProfile(db, &userProfile, userIDUint)
	if apiError != nil {
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, nil)
	return
}

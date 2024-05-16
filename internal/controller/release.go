package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ReleaseGet 用于处理获取发布房屋信息页面的请求
// @Summary 获取发布房屋信息页面
// @Description 获取发布房屋信息页面
// @Tags 发布
// @Accept json
// @Produce html
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {string} html "发布房屋信息页面"
// @Router /release [get]
//func ReleaseGet(c *gin.Context) {
//	//c.HTML(http.StatusOK, "release.html", nil)
//	ResponseSuccess(c, nil)
//	return
//}

// ReleasePost 用于处理发布房屋信息页面的POST请求,用于发布新的房源
// @Summary 发布房屋信息
// @Description 发布房屋信息
// @Tags 发布
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param user_id path string true "用户ID"
// @Param req body model.HouseRequest true "发布房屋信息请求"
// @Success 200 {object} controller.ResponseData "发布成功"
// @Failure 400 {object} controller.ResponseData "发布失败"
// @Router /user/:user_id/release [post]
func ReleasePost(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	var req model.HouseRequest
	err := c.ShouldBind(&req)
	if err != nil {
		zap.L().Debug(
			"ReleasePost: c.ShouldBind(&req) failed",
			zap.Error(err),
			zap.String("path", c.Request.URL.Path),
			zap.String("function", "ReleasePost"))
		ResponseErrorWithCode(c, codes.ReleaseBindDataError)
		return
	}

	apiError := logic.ProcessHouseAndImages(db, &req, c)
	if apiError != nil {
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, nil)
	return
}

// ReleasePut 用于处理更新房屋信息的请求
// @Summary 更新房屋信息
// @Description 更新房屋信息
// @Tags 发布
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param user_id path string true "用户ID"
// @Param req body model.HouseUpdateRequest true "更新房屋信息请求"
// @Success 200 {object} controller.ResponseData "更新成功"
// @Failure 400 {object} controller.ResponseData "更新失败"
// @Router /user/:user_id/release [put]
func ReleasePut(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	var req model.HouseUpdateRequest
	err := c.ShouldBind(&req)
	if err != nil {
		ResponseErrorWithCode(c, codes.ReleaseBindDataError)
		return
	}

	apiError := logic.UpdateHouseAndImages(db, &req, c)
	if apiError != nil {
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, nil)
	return
}

// ReleaseDeleteWholeHouse 用于处理删除整个房屋信息的请求
// @Summary 删除房屋信息
// @Description 删除房屋信息
// @Tags 发布
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param user_id path string true "用户ID"
// @Param house_id body string true "房屋ID"
// @Success 200 {object} controller.ResponseData "删除成功"
// @Failure 400 {object} controller.ResponseData "删除失败"
// @Router /user/:user_id/release [delete]
func ReleaseDeleteWholeHouse(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	var req struct {
		HouseID uint `json:"house_id"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		ResponseErrorWithCode(c, codes.ReleaseBindDataError)
		return
	}

	apiError := logic.DeleteHouse(db, req.HouseID)
	if apiError != nil {
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, nil)
}

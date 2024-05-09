package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
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
func ReleaseGet(c *gin.Context) {
	// c.HTML(http.StatusOK, "release.html", nil)
	ResponseSuccess(c, nil)
	return
}

// ReleasePost 用于处理发布房屋信息页面的POST请求,用于发布新的房源
// @Summary 发布房屋信息
// @Description 发布房屋信息
// @Tags 发布
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param req body model.HouseRequest true "发布房屋信息请求"
// @Success 200 {object} controller.ResponseData "发布成功"
// @Failure 400 {object} controller.ResponseData "发布失败"
// @Router /release [post]
func ReleasePost(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	var req model.HouseRequest
	err := c.ShouldBind(&req)
	if err != nil {
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

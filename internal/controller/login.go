package controller

import (
	"net/http"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LoginGet 用于处理用户的登录界面的GET请求
// @Summary 登录界面
// @Description 显示用户登录界面
// @Tags 登录
// @Accept json
// @Produce json
// @Success 200 {string} html "登录界面"
// @Router /auth/login [get]
func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// LoginPost 用于处理用户的登录界面的POST请求
// @Summary 登录接口
// @Description 用户登录接口
// @Tags 登录
// @Accept json
// @Produce json
// @Param object query model.LoginRequest false "查询参数"
// @Success 200 {object} controller.ResponseData "登录成功"
// @Failure 400 {object} controller.ResponseData "预约失败,具体原因查看json中的message字段和code字段"
// @Router /auth/login [post]
func LoginPost(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		ResponseErrorWithCode(c, codes.GetDBError)
	}

	var loginReq model.LoginRequest

	err := c.ShouldBind(&loginReq)
	if err != nil {
		ResponseErrorWithCode(c, codes.LoginInvalidParam)
	}

	loginResp, apiError := logic.LoginHandle(db, loginReq)
	if apiError != nil {
		ResponseError(c, *apiError)
	}

	ResponseSuccess(c, loginResp)
}

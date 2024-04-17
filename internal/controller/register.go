package controller

import (
	"net/http"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterGet 用于处理用户的注册界面的GET请求
// @Summary 注册界面
// @Description 显示用户注册界面
// @Tags 注册
// @Accept json
// @Produce json
// @Success 200 {string} html "注册界面"
// @Router /auth/register [get]
func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

// RegisterPost 用于处理用户的注册界面的POST请求
// @Summary 注册接口
// @Description 用户注册接口
// @Tags 注册
// @Accept json
// @Produce json
// @Param object query model.RegisterRequest false "查询参数"
// @Success 200 {object} controller.ResponseData "注册成功"
// @failure 200 {object} controller.ResponseData "注册失败"
// @Router /auth/register [post]
func RegisterPost(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		ResponseErrorWithCode(c, codes.GetDBError)
	}

	var registerReq model.RegisterRequest
	err := c.ShouldBind(&registerReq)
	if err != nil {
		ResponseErrorWithCode(c, codes.RegisterInvalidParam)
	}

	apiError := logic.RegisterHandle(db, registerReq)
	if apiError != nil {
		ResponseError(c, *apiError)
	}

	ResponseSuccess(c, nil)
}

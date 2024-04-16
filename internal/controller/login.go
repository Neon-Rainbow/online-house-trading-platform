package controller

import (
	"net/http"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
)

// LoginGet 用于处理用户的登录界面的GET请求
func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// LoginPost 用于处理用户的登录界面的POST请求
func LoginPost(c *gin.Context) {
	var loginReq model.LoginRequest

	err := c.ShouldBind(&loginReq)
	if err != nil {
		ResponseErrorWithCode(c, LoginInvalidParam)
	}

	loginResp, apiError := logic.LoginHandle(c, loginReq)
	if apiError != nil {
		ResponseError(c, *apiError)
	}

	ResponseSuccess(c, loginResp)
}

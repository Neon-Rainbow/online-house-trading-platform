package controller

import (
	"net/http"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
)

// RegisterGet 用于处理用户的注册界面的GET请求
func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func RegisterPost(c *gin.Context) {
	var registerReq model.RegisterRequest
	err := c.ShouldBind(&registerReq)
	if err != nil {
		ResponseErrorWithCode(c, RegisterInvalidParam)
	}

	apiError := logic.RegisterHandle(c, registerReq)
	if apiError != nil {
		ResponseError(c, *apiError)
	}

	ResponseSuccess(c, nil)
}

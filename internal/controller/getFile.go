package controller

import (
	"fmt"
	"online-house-trading-platform/codes"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type fileUrl struct {
	Url string `json:"url" form:"url"`
}

// GetFileByURL 处理根据URL获取文件的请求
// @Summary 根据URL获取文件
// @Description 通过提供的URL获取文件
// @Tags 文件
// @Accept json
// @Param url query string true "文件URL"
// @Success 200 {string} string "请求成功"
// @Failure 400 {object} object "请求参数错误"
// @Router /getFile [get]
func GetFileByURL(c *gin.Context) {
	var url fileUrl
	err := c.ShouldBind(&url)
	if err != nil {
		zap.L().Error("GetFileByURL", zap.Error(err))
		ResponseErrorWithCode(c, codes.LoginInvalidParam)
		return
	}
	filePath := url.Url
	if filePath == "" {
		zap.L().Error("filePath为空")
		ResponseErrorWithCode(c, codes.LoginInvalidParam)
	}

	fmt.Println("filepath:" + filePath)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filePath))
	c.File(filePath)
	ResponseSuccess(c, filePath)
	return
}

func GetLogFile(c *gin.Context) {
	filePath := "./application.log"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filePath))
	c.File(filePath)
	ResponseSuccess(c, filePath)
	return
}

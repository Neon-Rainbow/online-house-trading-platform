package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomePageGet 用来处理首页的Get请求
func HomePageGet(c *gin.Context) {
	c.HTML(http.StatusOK, "frontpage.html", nil)
}

// LearnMoreGet 用来处理LearnMore页面的Get请求
func LearnMoreGet(c *gin.Context) {
	c.HTML(http.StatusOK, "learn_more.html", nil)
}

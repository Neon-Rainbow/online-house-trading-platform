package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ViewFavouritesGet 用于处理用户查看收藏界面的GET请求
func ViewFavouritesGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/user/favourites",
		"method":  "GET",
	})
}

// DeleteFavouritesDelete 用于处理用户删除收藏的信息
func DeleteFavouritesDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/user/favourites",
		"method":  "Delete",
	})
}

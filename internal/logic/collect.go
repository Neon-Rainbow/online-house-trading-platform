package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
)

// CollectHandle 用来处理用户收藏房屋的请求
func CollectHandle(c *gin.Context, favourite *model.Favourite, userID uint) *model.Error {
	db, err := dao.GetDB(c)
	if err != nil {
		return &model.Error{StatusCode: codes.GetDBError}
	}

	if favourite.HouseID == 0 || favourite.UserID == 0 {
		return &model.Error{StatusCode: codes.ReserveInvalidParam}
	}
	favourite.UserID = userID
	err = dao.CreateFavorite(db, favourite)
	if err != nil {
		return &model.Error{StatusCode: codes.ReserveError}
	}
	return nil
}

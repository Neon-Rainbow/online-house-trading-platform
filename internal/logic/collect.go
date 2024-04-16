package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"

	"gorm.io/gorm"
)

// CollectHandle 用来处理用户收藏房屋的请求
func CollectHandle(db *gorm.DB, favourite *model.Favourite, userID uint) *model.Error {
	if favourite.HouseID == 0 || favourite.UserID == 0 {
		return &model.Error{StatusCode: codes.ReserveInvalidParam}
	}
	favourite.UserID = userID
	err := dao.CreateFavorite(db, favourite)
	if err != nil {
		return &model.Error{StatusCode: codes.ReserveError}
	}
	return nil
}

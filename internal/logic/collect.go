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

	exist, _ := dao.CheckCombinationUserIDAndHouseIDInFavouriteExists(db, favourite.UserID, favourite.HouseID)
	if exist {
		return &model.Error{StatusCode: codes.RecordExists}
	}

	err := dao.CreateFavorite(db, favourite)
	if err != nil {
		return &model.Error{StatusCode: codes.ReserveError}
	}
	return nil
}

func GetUserFavourites(db *gorm.DB, userID uint) ([]model.Favourite, *model.Error) {
	favourites, err := dao.GetUserFavourites(db, userID)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetUserFavouritesError}
	}
	return favourites, nil
}

func GetAllFavourites(db *gorm.DB) (*[]model.Favourite, *model.Error) {
	favourites, err := dao.GetAllFavourites(db)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetUserFavouritesError}
	}
	return favourites, nil
}

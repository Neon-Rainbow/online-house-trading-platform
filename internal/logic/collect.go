package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
)

// CollectHandle 用来处理用户收藏房屋的请求
// @title CollectHandle
// @description 处理用户收藏房屋的请求
// @param favourite *model.Favourite 收藏信息
// @param userID uint 用户ID
// @return *model.Error 错误信息
func CollectHandle(favourite *model.Favourite, userID uint) *model.Error {
	if favourite.HouseID == 0 || favourite.UserID == 0 {
		return &model.Error{StatusCode: codes.ReserveInvalidParam}
	}
	favourite.UserID = userID

	exist, _ := dao.CheckCombinationUserIDAndHouseIDInFavouriteExists(favourite.UserID, favourite.HouseID)
	if exist {
		return &model.Error{StatusCode: codes.RecordExists}
	}

	err := dao.CreateFavorite(favourite)
	if err != nil {
		return &model.Error{StatusCode: codes.ReserveError}
	}
	return nil
}

func GetUserFavourites(userID uint) ([]model.Favourite, *model.Error) {
	favourites, err := dao.GetUserFavourites(userID)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetUserFavouritesError}
	}
	return favourites, nil
}

func GetAllFavourites() ([]model.Favourite, *model.Error) {
	favourites, err := dao.GetAllFavourites()
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetUserFavouritesError}
	}
	return favourites, nil
}

func DeleteFavourite(userID uint, houseID uint) *model.Error {
	err := dao.DeleteFavouriteByUserID(userID, houseID)
	if err != nil {
		return &model.Error{StatusCode: codes.LoginServerBusy}
	}
	return nil
}

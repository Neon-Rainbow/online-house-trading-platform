package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
)

// GetAllUsers 用于获取所有用户
// @title GetAllUsers
// @description 获取所有用户
// @param includeDeleted string 是否包含已删除用户
// @return users *[]model.User 用户列表
// @return apiError *model.Error 错误信息
func GetAllUsers(includeDeleted bool) (users []model.User, apiError *model.Error) {
	users, err := dao.GetAllUsers(includeDeleted)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetAllUsersError}
	}
	return users, nil
}

// GetAllHouses 用于获取所有房屋
// @title GetAllHouses
// @description 获取所有房屋
// @return houses []model.House 房屋列表
// @return apiError *model.Error 错误信息
func GetAllHouses() (houses []model.House, apiError *model.Error) {
	houses, err := dao.GetAllHouses()
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetAllHousesError}
	}
	return houses, nil
}

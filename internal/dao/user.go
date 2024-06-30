package dao

import (
	"errors"
	"online-house-trading-platform/pkg/database"
	"online-house-trading-platform/pkg/model"

	"gorm.io/gorm"
)

// GetUserByUsername 用于根据用户名获取用户信息
func GetUserByUsername(username string) (*model.User, error) {
	db := database.Database
	var user model.User
	result := db.Preload("Avatar").Where("username = ?", username).First(&user)
	return &user, result.Error
}

// CreateUser 用于创建用户
func CreateUser(user *model.User) error {
	db := database.Database
	return db.Create(user).Error
}

// CheckUserExists 用于检查用户名和邮箱是否已存在,返回两个bool值,第一个bool值表示用户名是否存在,第二个bool值表示邮箱是否存在
// @param username string 用户名
// @param email string 邮箱
// @return isUsernameExists bool 用户名是否存在
// @return isEmailExists bool 邮箱是否存在
// @return err error 错误信息
func CheckUserExists(username string, email string) (isUsernameExists bool, isEmailExists bool, err error) {
	db := database.Database
	var count int64
	if err := db.Model(&model.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, false, err
	}
	usernameExists := count > 0

	if err := db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, false, err
	}
	emailExists := count > 0

	return usernameExists, emailExists, nil
}

// GetUserFavourites 用于获取用户的收藏
func GetUserFavourites(userID uint) ([]model.Favourite, error) {
	db := database.Database
	var favourites []model.Favourite
	result := db.Where("user_id = ?", userID).Find(&favourites)
	return favourites, result.Error
}

// GetUserProfile 用于获取用户的个人信息
func GetUserProfile(userID uint) (*model.User, error) {
	db := database.Database
	var userProfile *model.User
	result := db.Preload("Avatar").First(&userProfile, userID)
	return userProfile, result.Error
}

// ModifyUserProfile 用于修改用户的个人信息
func ModifyUserProfile(requestModel *model.UserReq, userID uint) error {
	db := database.Database
	return db.Model(model.User{}).Where("id = ?", userID).Updates(requestModel).Error
}

// GetReserve 用于获取用户的预约信息
func GetReserve(idUint uint) ([]model.Reserve, error) {
	db := database.Database
	var reserve []model.Reserve
	result := db.Where("user_id = ?", idUint).Find(&reserve)
	return reserve, result.Error
}

// CreateUserAvatar 用于创建用户的头像
func CreateUserAvatar(avatar *model.UserAvatar) error {
	db := database.Database
	return db.Create(avatar).Error
}

// ModifyUserAvatar 用于修改用户的头像
func ModifyUserAvatar(avatar *model.UserAvatar) error {
	db := database.Database
	err := db.Save(avatar).Error
	return err
}

// GetUserRelease 获取某个用户发布的房屋信息
func GetUserRelease(userID uint) ([]model.House, error) {
	db := database.Database
	var houses []model.House
	result := db.Preload("Images").Where("owner_id = ?", userID).Find(&houses)
	if result.Error != nil {
		return nil, result.Error
	}
	return houses, nil
}

// IsUserAdmin 用于判断用户是否为管理员
func IsUserAdmin(id uint) (isAdmin bool, err error) {
	db := database.Database
	var user model.User
	result := db.First(&user, id)
	if result.Error != nil {
		return false, result.Error
	}
	return user.Role == "admin", nil
}

// GetAllUsers 用于获取所有用户
// @param includeDeleted bool 是否包含已删除用户 默认不包含
// @return []model.User 用户列表
// @return error 错误信息
func GetAllUsers(includeDeleted ...bool) ([]model.User, error) {
	db := database.Database
	var user []model.User
	var result *gorm.DB
	if len(includeDeleted) > 0 && includeDeleted[0] {
		result = db.Unscoped().Preload("Avatar").Find(&user)
	} else {
		// 默认不包含已删除的用户
		result = db.Preload("Avatar").Find(&user)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func GetAllReserve() ([]model.Reserve, error) {
	db := database.Database
	var reserve []model.Reserve
	result := db.Find(&reserve)
	if result.Error != nil {
		return nil, result.Error
	}
	return reserve, nil
}

// GetAllFavourites 用于获取所有用户的收藏
func GetAllFavourites() ([]model.Favourite, error) {
	db := database.Database
	var favourites []model.Favourite
	result := db.Find(&favourites)
	if result.Error != nil {
		return nil, result.Error
	}
	return favourites, nil
}

// CheckCombinationUserIDAndHouseIDInFavouriteExists 检查用户ID和房屋ID的组合在favourite表中是否存在
func CheckCombinationUserIDAndHouseIDInFavouriteExists(userID uint, houseID uint) (bool, error) {
	db := database.Database
	var favourite model.Favourite
	result := db.Where("user_id = ? AND house_id = ?", userID, houseID).First(&favourite)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 没有找到记录，返回不存在
			return false, nil
		}
		// 数据库查询出错
		return false, result.Error
	}
	// 找到了记录，返回存在
	return true, nil
}

// DeleteFavouriteByUserID 用于删除用户的收藏
func DeleteFavouriteByUserID(userID uint, houseID uint) error {
	db := database.Database
	result := db.Delete(&model.Favourite{}, "user_id = ? AND house_id = ?", userID, houseID)
	return result.Error
}

// DeleteAppointment 用于删除用户的预约
func DeleteAppointment(userID uint, houseID uint) error {
	db := database.Database
	result := db.Delete(&model.Reserve{}, "user_id = ? AND house_id = ?", userID, houseID)
	return result.Error
}

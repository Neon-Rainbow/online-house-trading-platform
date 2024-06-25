package dao

import (
	"errors"
	"fmt"
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
func CheckUserExists(username, email string) (bool, bool, error) {
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
func GetUserFavourites(id uint) ([]model.Favourite, error) {
	db := database.Database
	var favourites []model.Favourite
	result := db.Where("user_id = ?", id).Find(&favourites)
	return favourites, result.Error
}

// GetUserProfile 用于获取用户的个人信息
func GetUserProfile(idUint uint) (*model.User, error) {
	db := database.Database
	var userProfile *model.User
	result := db.Preload("Avatar").First(&userProfile, idUint)
	return userProfile, result.Error
}

// ModifyUserProfile 用于修改用户的个人信息
func ModifyUserProfile(m *model.UserReq, idUint uint) error {
	db := database.Database
	return db.Model(model.User{}).Where("id = ?", idUint).Updates(m).Error
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
	fmt.Print(avatar)
	err := db.Raw("UPDATE user_avatars SET url = ? WHERE user_id = ?", avatar.URL, avatar.UserID).Error
	return err
}

// GetUserRelease 获取某个用户发布的房屋信息
func GetUserRelease(userID uint) (*[]model.House, error) {
	db := database.Database
	var houses *[]model.House
	result := db.Preload("Images").Where("owner_id = ?", userID).Find(&houses)
	if result.Error != nil {
		return nil, result.Error
	}
	return houses, nil
}

// IsUserAdmin 用于判断用户是否为管理员
func IsUserAdmin(id uint) (bool, error) {
	db := database.Database
	var user model.User
	result := db.First(&user, id)
	if result.Error != nil {
		return false, result.Error
	}
	return user.Role == "admin", nil
}

func GetAllUsers(includeDeleted string) (*[]model.User, error) {
	db := database.Database
	var user *[]model.User
	var result *gorm.DB
	//fmt.Println("includeDeleted: ", includeDeleted)
	if includeDeleted == "true" {
		//fmt.Println("includeDeleted")
		result = db.Unscoped().Find(&user)
	} else {
		//fmt.Println("not includeDeleted")
		result = db.Find(&user)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func GetAllReserve() (*[]model.Reserve, error) {
	db := database.Database
	var reserve *[]model.Reserve
	result := db.Find(&reserve)
	if result.Error != nil {
		return nil, result.Error
	}
	return reserve, nil
}

func GetAllFavourites() (*[]model.Favourite, error) {
	db := database.Database
	var favourites *[]model.Favourite
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

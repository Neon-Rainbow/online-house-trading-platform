package dao

import (
	"fmt"
	"online-house-trading-platform/pkg/model"

	"gorm.io/gorm"
)

// GetUserByUsername 用于根据用户名获取用户信息
func GetUserByUsername(db *gorm.DB, username string) (*model.User, error) {
	var user model.User
	result := db.Preload("Avatar").Where("username = ?", username).First(&user)
	return &user, result.Error
}

// CreateUser 用于创建用户
func CreateUser(db *gorm.DB, user *model.User) error {
	return db.Create(user).Error
}

// CheckUserExists 用于检查用户名和邮箱是否已存在,返回两个bool值,第一个bool值表示用户名是否存在,第二个bool值表示邮箱是否存在
func CheckUserExists(db *gorm.DB, username, email string) (bool, bool, error) {
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
func GetUserFavourites(db *gorm.DB, id uint) ([]model.Favourite, error) {
	var favourites []model.Favourite
	result := db.Where("user_id = ?", id).Find(&favourites)
	return favourites, result.Error
}

// GetUserProfile 用于获取用户的个人信息
func GetUserProfile(db *gorm.DB, idUint uint) (*model.User, error) {
	var userProfile *model.User
	result := db.Preload("Avatar").First(&userProfile, idUint)
	return userProfile, result.Error
}

// ModifyUserProfile 用于修改用户的个人信息
func ModifyUserProfile(db *gorm.DB, m *model.UserReq, idUint uint) error {
	return db.Model(model.User{}).Where("id = ?", idUint).Updates(m).Error
}

// GetReserve 用于获取用户的预约信息
func GetReserve(db *gorm.DB, idUint uint) ([]model.Reserve, error) {
	var reserve []model.Reserve
	result := db.Where("user_id = ?", idUint).Find(&reserve)
	return reserve, result.Error
}

// CreateUserAvatar 用于创建用户的头像
func CreateUserAvatar(db *gorm.DB, avatar *model.UserAvatar) error {
	return db.Create(avatar).Error
}

func ModifyUserAvatar(db *gorm.DB, avatar *model.UserAvatar) error {
	fmt.Print(avatar)
	err := db.Raw("UPDATE user_avatars SET url = ? WHERE user_id = ?", avatar.URL, avatar.UserID).Error
	return err
}
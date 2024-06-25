package dao

import (
	"online-house-trading-platform/pkg/database"
	"online-house-trading-platform/pkg/model"
)

// CreateAppointment 用于实现用户预约房屋
func CreateAppointment(reserve *model.Reserve) error {
	db := database.Database
	return db.Create(reserve).Error
}

// CreateFavorite 用于实现用户收藏房屋
func CreateFavorite(favorite *model.Favourite) error {
	db := database.Database
	return db.Create(favorite).Error
}

// GetAllHouseInformation 用于获取数据库中的所有房屋信息
func GetAllHouseInformation() ([]model.House, error) {
	db := database.Database
	var houses []model.House
	result := db.Preload("Images").Find(&houses)
	if result.Error != nil {
		return nil, result.Error
	}
	return houses, nil
}

// GetHouseInformationByID 用于获取数据库中指定ID的房屋信息
func GetHouseInformationByID(houseID uint) (*model.House, error) {
	db := database.Database
	var house *model.House
	result := db.Preload("Images").First(&house, houseID)
	if result.Error != nil {
		return house, result.Error
	}
	return house, nil
}

// CreateHouse 用于创建房屋记录
func CreateHouse(house *model.House) error {
	db := database.Database
	return db.Create(house).Error
}

// CreateHouseImages 用于在数据库中创建多个房屋图片记录
func CreateHouseImages(images []model.HouseImage) error {
	db := database.Database
	return db.Create(&images).Error
}

// DeleteHouse 用于删除房屋记录
// 该函数会返回被删除的房屋记录
// ./uploads/houses/文件夹下的图片文件不会被删除,dao层给logic层返回房屋信息后logic层来删除图片内容
func DeleteHouse(houseID uint) (*model.House, error) {
	db := database.Database
	var house model.House
	result := db.Preload("Images").First(&house, houseID)
	if result.Error != nil {
		return &house, result.Error
	}

	// 删除与房屋关联的图片记录
	err := db.Delete(&model.HouseImage{}, "house_id = ?", houseID).Error
	if err != nil {
		return &house, err
	}

	// 删除房屋记录
	err = db.Delete(&house).Error
	return &house, err
}

// UpdateHouse 更新房屋信息
func UpdateHouse(house *model.House, updateFields interface{}) error {
	db := database.Database
	return db.Model(&house).Updates(updateFields).Error
}

// DeleteHouseImages 删除房屋的旧图片记录
func DeleteHouseImages(houseID uint) error {
	db := database.Database
	return db.Where("house_id = ?", houseID).Delete(&model.HouseImage{}).Error
}

// CreateHouseImage 插入新的房屋图片记录
func CreateHouseImage(image *model.HouseImage) error {
	db := database.Database
	return db.Save(image).Error
}

// GetAllHouses 用于获取数据库中的所有房屋信息
func GetAllHouses() (*[]model.House, error) {
	db := database.Database
	var houses *[]model.House
	result := db.Find(&houses)
	if result.Error != nil {
		return nil, result.Error
	}
	return houses, nil
}

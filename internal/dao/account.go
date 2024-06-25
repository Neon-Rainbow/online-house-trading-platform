package dao

import (
	"online-house-trading-platform/pkg/database"
	"online-house-trading-platform/pkg/model"
)

func DeleteAccount(id uint) error {
	result := database.Database.Where("id = ?", id).Preload("Avatar").Delete(&model.User{})
	return result.Error
}

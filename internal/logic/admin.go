package logic

import (
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
)

func GetAllUsers(includeDeleted string) (users *[]model.User, err error) {
	users, err = dao.GetAllUsers(includeDeleted)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetAllHouses() (houses *[]model.House, err error) {
	houses, err = dao.GetAllHouses()
	if err != nil {
		return nil, err
	}
	return houses, nil
}

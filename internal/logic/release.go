package logic

import (
	"fmt"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

// ProcessHouseAndImages 用于处理房屋和图片
func ProcessHouseAndImages(req *model.HouseRequest) *model.Error {

	// 创建房屋记录
	house := model.House{
		Owner:            req.Owner,
		OwnerID:          req.OwnerID,
		Title:            req.Title,
		Description:      req.Description,
		Price:            req.Price,
		Address:          req.Address,
		HouseOrientation: req.HouseOrientation,
		Layout:           req.Layout,
		Area:             req.Area,
		Floor:            req.Floor,
		RentPrice:        req.RentPrice,
		Type:             req.Type,
		PostCode:         req.PostCode,
	}

	err := dao.CreateHouse(&house)
	if err != nil {
		return &model.Error{StatusCode: codes.CreateHouseError, Message: "创建房屋记录失败"}
	}

	//保存房屋图片
	for _, img := range req.Images {
		dst := fmt.Sprintf("./uploads/houses/%d/%s", house.ID, img.Filename)
		err := saveUploadedFile(img, dst)
		if err != nil {
			return &model.Error{StatusCode: codes.CreateHouseImageError, Message: "保存房屋图片到./uploads/houses/文件夹下失败"}
		}
		apiError := dao.CreateHouseImages([]model.HouseImage{{HouseID: house.ID, URL: dst}})
		if apiError != nil {
			return &model.Error{StatusCode: codes.CreateHouseImageError, Message: "保存房屋图片到数据库中失败"}
		}
	}
	return nil
}

// DeleteHouse 用于删除房屋记录
func DeleteHouse(houseID uint) *model.Error {
	house, err := dao.DeleteHouse(houseID)
	if err != nil {
		return &model.Error{StatusCode: codes.DeleteHouseError, Message: "删除房屋记录失败"}
	}
	for _, houseImage := range house.Images {
		filepath := fmt.Sprintf("./%s", houseImage.URL)
		_ = os.Remove(filepath)
	}
	return nil
}

// snakeCase 将字段名转换为蛇形命名法（snake_case）
//
//	func snakeCase(s string) string {
//		runes := []rune(s)
//		length := len(runes)
//
//		var result []rune
//		for i := 0; i < length; i++ {
//			if i > 0 && runes[i] >= 'A' && runes[i] <= 'Z' {
//				result = append(result, '_')
//			}
//			result = append(result, runes[i])
//		}
//
//		return string(result)
//	}
func snakeCase(s string) string {
	if s == "HouseID" {
		return "id"
	}
	// 正则表达式用于匹配大写字母
	reg := regexp.MustCompile("([a-z0-9])([A-Z])")
	// 替换匹配项并转换为小写
	snake := reg.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(snake)
}

// UpdateHouseAndImages 用于更新房屋信息和图片
func UpdateHouseAndImages(req *model.HouseUpdateRequest, c *gin.Context) *model.Error {
	// 更新房屋信息
	existingHouse, err := dao.GetHouseInformationByID(req.HouseID)
	if err != nil {
		return &model.Error{StatusCode: codes.UpdateHouseError, Message: "房屋信息获取失败"}
	}

	var house model.House
	// 更新房屋字段，仅使用非空值
	house.ID = existingHouse.ID
	house.CreatedAt = existingHouse.CreatedAt
	updateFields := make(map[string]interface{})

	reqValue := reflect.ValueOf(req).Elem()
	existingValue := reflect.ValueOf(existingHouse).Elem() // 确保 existingHouse 被解引用

	for i := 0; i < reqValue.NumField(); i++ {
		field := reqValue.Type().Field(i)
		fieldName := field.Name
		fieldValue := reqValue.Field(i)

		if !fieldValue.IsZero() {
			updateFields[snakeCase(fieldName)] = fieldValue.Interface()
		} else {
			updateFields[snakeCase(fieldName)] = existingValue.FieldByName(fieldName).Interface()
		}
	}

	// 删除 updateFields 中的 images 字段
	delete(updateFields, "images")

	// 更新房屋信息
	err = dao.UpdateHouse(&house, updateFields)
	if err != nil {
		return &model.Error{StatusCode: codes.UpdateHouseError, Message: "更新房屋信息失败"}
	}

	// 删除旧的图片记录
	if err := dao.DeleteHouseImages(req.HouseID); err != nil {
		return &model.Error{StatusCode: codes.UpdateHouseError, Message: "删除旧图片记录失败"}
	}

	// 保存新图片到文件夹并插入新的图片记录
	for _, imgFile := range req.Images {
		// 保存图片到 /upload 文件夹
		filename := generateRandomFileName()
		saveFilePath := fmt.Sprintf("./uploads/houses/%v/%s", house.ID, filename)
		if err := saveUploadedFile(imgFile, saveFilePath); err != nil {
			return &model.Error{StatusCode: codes.UpdateHouseError, Message: "保存图片文件失败"}
		}

		// 插入新的图片记录
		img := model.HouseImage{
			HouseID: req.HouseID,
			URL:     saveFilePath,
		}
		if err := dao.CreateHouseImage(&img); err != nil {
			return &model.Error{StatusCode: codes.UpdateHouseError, Message: "插入新图片记录失败"}
		}
	}

	return nil
}

// GetUserRelease 用于获取用户发布的房屋信息
func GetUserRelease(userID uint) ([]model.House, *model.Error) {
	houses, err := dao.GetUserRelease(userID)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetHouseInfoError, Message: "获取房屋信息失败"}
	}
	return houses, nil
}

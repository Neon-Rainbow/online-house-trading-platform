package logic

import (
	"fmt"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProcessHouseAndImages 用于处理房屋和图片
func ProcessHouseAndImages(db *gorm.DB, req *model.HouseRequest, c *gin.Context) *model.Error {

	// 创建房屋记录
	house := model.House{
		Owner:       req.Owner,
		OwnerID:     req.OwnerID,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Address:     req.Address,
	}

	err := dao.CreateHouse(db, &house)
	if err != nil {
		return &model.Error{StatusCode: codes.CreateHouseError, Message: "创建房屋记录失败"}
	}

	//保存房屋图片
	for _, img := range req.Images {
		dst := fmt.Sprintf("./uploads/houses/%d/%s", house.ID, img.Filename)
		err := c.SaveUploadedFile(img, dst)
		if err != nil {
			return &model.Error{StatusCode: codes.CreateHouseImageError, Message: "保存房屋图片到./uploads/houses/文件夹下失败"}
		}
		apiError := dao.CreateHouseImages(db, []model.HouseImage{{HouseID: house.ID, URL: dst}})
		if apiError != nil {
			return &model.Error{StatusCode: codes.CreateHouseImageError, Message: "保存房屋图片到数据库中失败"}
		}
	}
	return nil
}

// DeleteHouse 用于删除房屋记录
func DeleteHouse(db *gorm.DB, houseID uint) *model.Error {
	house, err := dao.DeleteHouse(db, houseID)
	if err != nil {
		return &model.Error{StatusCode: codes.DeleteHouseError, Message: "删除房屋记录失败"}
	}
	for _, houseImage := range house.Images {
		filepath := fmt.Sprintf("./%s", houseImage.URL)
		err := os.Remove(filepath)
		if err != nil {
			return &model.Error{StatusCode: codes.DeleteHouseError, Message: "删除房屋图片失败"}
		}
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

func UpdateHouseAndImages(db *gorm.DB, req *model.HouseUpdateRequest, c *gin.Context) *model.Error {
	// 更新房屋信息
	existingHouse, err := dao.GetHouseInformationByID(db, req.HouseID)
	if err != nil {
		return &model.Error{StatusCode: codes.UpdateHouseError, Message: "房屋信息获取失败"}

	}

	var house model.House
	// 更新房屋字段，仅使用非空值
	house.ID = existingHouse.ID
	house.CreatedAt = existingHouse.CreatedAt
	updateFields := make(map[string]interface{})

	reqValue := reflect.ValueOf(req).Elem()
	existingValue := reflect.ValueOf(existingHouse)

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
	//删除updateFields中的images字段
	delete(updateFields, "images")

	//if err := dao.UpdateHouse(db, &house); err != nil {
	//	return &model.Error{StatusCode: codes.UpdateHouseError, Message: "更新房屋信息失败"}
	//}
	err = dao.UpdateHouse(db, &house, updateFields)
	if err != nil {
		return &model.Error{StatusCode: codes.UpdateHouseError, Message: "更新房屋信息失败"}
	}

	// 删除旧的图片记录
	if err := dao.DeleteHouseImages(db, req.HouseID); err != nil {
		return &model.Error{StatusCode: codes.UpdateHouseError, Message: "删除旧图片记录失败"}
	}

	// 保存新图片到文件夹并插入新的图片记录
	for _, imgFile := range req.Images {
		// 保存图片到 /upload 文件夹
		filename := filepath.Base(imgFile.Filename)
		saveFilePath := fmt.Sprintf("./uploads/houses/%v/%s", house.ID, filename)
		if err := c.SaveUploadedFile(imgFile, saveFilePath); err != nil {
			return &model.Error{StatusCode: codes.UpdateHouseError, Message: "保存图片文件失败"}
		}

		// 插入新的图片记录
		img := model.HouseImage{
			HouseID: req.HouseID,
			URL:     saveFilePath,
		}
		if err := dao.CreateHouseImage(db, &img); err != nil {
			return &model.Error{StatusCode: codes.UpdateHouseError, Message: "插入新图片记录失败"}
		}
	}

	return nil

}

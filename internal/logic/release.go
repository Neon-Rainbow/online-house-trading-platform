package logic

import (
	"fmt"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"

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

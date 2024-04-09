package houses

import (
	"errors"
	"log"
	"net/http"
	"online-house-trading-platform/pkg/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetHouseInfo 用于给定ID,然后获取指定ID的房屋信息
func GetHouseInfo(db *gorm.DB, houseID uint) (*model.House, error) {
	var house model.House
	result := db.Preload("Images").First(&house, houseID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("house not found")
		} else {
			return nil, result.Error
		}
	}
	return &house, nil
}

// HouseByIDGet 用于获取URL中指定ID的房屋信息
// URL: GET /api/houses/:id
func HouseByIDGet(c *gin.Context) {
	db, exists := c.MustGet("db").(*gorm.DB)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法获取数据库连接",
		})
		return
	}

	houseIdStr := c.Param("id")
	houseID, err := strconv.ParseUint(houseIdStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "非法的ID格式",
		})
		log.Printf("非法的ID格式: %v", err)
		return
	}

	house, err := GetHouseInfo(db, uint(houseID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		log.Printf("房屋编号为:%v 查询房屋信息失败: %v ", houseID, err)
		return
	}
	//返回房屋信息
	c.JSON(http.StatusOK, house)
}

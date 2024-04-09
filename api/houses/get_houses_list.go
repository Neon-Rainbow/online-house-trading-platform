package houses

import (
	"log"
	"net/http"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetaCertainAmountOfHouseInformation 用于获取数据库中的一定数量的房屋信息
//func GetaCertainAmountOfHouseInformation(db *gorm.DB, begin int, end int) ([]model.House, error) {
//	var houses []model.House
//	// 从数据库中获取第begin到第end的房屋信息
//	result := db.Offset(begin).Limit(end - begin).Find(&houses)
//	return houses, result.Error
//}

// GetHouseList 用于处理用户的添加房屋界面的GET请求
func GetHouseList(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		c.JSON(500, gin.H{
			"error": "无法获取数据库连接",
		})
		return
	}

	var houses []model.House

	// 从数据库中查询所有的房屋信息,返回给前端
	result := db.Preload("Images").Find(&houses)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "数据库查询出错",
		})
		log.Printf("查询房屋信息时数据库查询出错, 错误原因为 %v", result.Error)
		return
	}
	// 将查询得到的结果返回给前端
	c.JSON(http.StatusOK, gin.H{
		"houses": houses,
		"length": result.RowsAffected,
	})
}

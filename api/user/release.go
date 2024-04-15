package user

import (
	"fmt"
	"log"
	"net/http"
	"online-house-trading-platform/pkg/model"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// uploadFile 用于处理上传文件的请求
func uploadFile(c *gin.Context) ([]model.HouseImage, error) {
	form, err := c.MultipartForm()
	if err != nil {
		log.Printf("上传文件失败,错误原因: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "上传文件失败",
		})
		return nil, err
	}

	files := form.File["files"]
	var images []model.HouseImage

	//遍历所有文件
	for index, file := range files {
		log.Printf("index: %v, file: %v", index, file.Filename)
		dst := fmt.Sprintf("./uploads/houses/%v/%v_%v", c.MustGet("house_id"), file.Filename, index)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			log.Printf("保存文件失败,错误原因: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "保存文件失败",
			})
			return nil, err
		}
		image := model.HouseImage{
			HouseID: c.MustGet("house_id").(uint),
			URL:     dst,
		}
		images = append(images, image)
		log.Printf("文件保存成功,保存路径: %v", dst)
	}
	log.Printf("共计上传文件数量: %v", len(files))
	return images, nil
}

// deleteFile 用于删除文件
func deleteFile(c *gin.Context, filename string) {
	err := os.Remove(filename)
	if err != nil {
		log.Printf("删除文件失败,错误原因: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "删除文件失败",
		})
		return
	}
	log.Printf("文件删除成功,文件路径: %v", filename)
}

// ReleaseGet 用于处理用户发布信息界面的GET请求
func ReleaseGet(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

// ReleasePost 用于处理用户发布信息界面的POST请求,用于发布新的房源
func ReleasePost(c *gin.Context) {
	// 从上下文中获取数据库连接
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法获取数据库连接",
		})
		return
	}

	images, err := uploadFile(c) // 上传文件
	// 保存图片
	for _, img := range images {
		err := db.Create(&img).Error
		if err != nil {
			log.Printf("保存图片失败,错误原因: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "保存图片失败",
			})
			return
		}
	}

	// 绑定数据
	var house model.House
	err = c.ShouldBind(&house)
	if err != nil {
		log.Printf("绑定数据失败,错误原因: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "绑定数据失败",
		})
		return
	}

	err = db.Create(&house).Error
	if err != nil {
		log.Printf("创建房屋失败,错误原因: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "创建房屋失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "发布成功",
	})
}

// ReleasePut 用于处理用户发布信息界面的PUT请求,用于更新用户发布的房源信息
func ReleasePut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/user/release",
		"method":  "PUT",
		"user_id": c.Param("user_id"),
	})
}

// ReleaseDelete 用于删除用户发布的房源信息
func ReleaseDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/user/release",
		"method":  "Delete",
		"user_id": c.Param("user_id"),
	})
}

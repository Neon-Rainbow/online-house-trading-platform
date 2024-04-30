package logic

import (
	"fmt"
	"io"
	"mime/multipart"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

// uploadFiles 用于上传文件
func uploadFiles(files []*multipart.FileHeader, houseID uint) ([]model.HouseImage, *model.Error) {
	var images []model.HouseImage
	dir := fmt.Sprintf("./uploads/houses/%d", houseID)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, &model.Error{StatusCode: codes.CreateDirError, Message: "创建文件夹失败"}
	}
	for index, file := range files {
		filename := filepath.Base(file.Filename)             // 获取文件名，并防止路径遍历安全问题
		dst := fmt.Sprintf("%s/%s_%d", dir, filename, index) // 创建文件的最终存储路径
		if err := saveUploadedFile(file, dst); err != nil {
			return nil, &model.Error{StatusCode: codes.SaveFileError, Message: "保存文件失败"}
		}

		image := model.HouseImage{
			HouseID: houseID, // 这里使用HouseID作为ownerID，实际应用中可能需要调整
			URL:     dst,
		}
		images = append(images, image)
	}
	return images, nil
}

// saveUploadedFile 用于保存上传的文件
func saveUploadedFile(file *multipart.FileHeader, dst string) error {
	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("打开文件失败: %v", err)
	}
	defer src.Close() // 确保文件最终关闭

	// 创建目标文件
	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer out.Close() // 确保文件最终关闭

	// 将数据从源文件复制到目标文件
	if _, err = io.Copy(out, src); err != nil {
		return fmt.Errorf("写文件失败: %v", err)
	}
	return nil
}

// ProcessHouseAndImages 用于处理房屋和图片
func ProcessHouseAndImages(db *gorm.DB, req *model.HouseRequest, ownerID uint) *model.Error {
	// 第一步: 上传文件并创建图片记录
	images, apiError := uploadFiles(req.Images, ownerID)
	if apiError != nil {
		return apiError // 返回从 uploadFiles 函数获取的错误
	}

	// 第二步: 创建房屋记录
	house := model.House{
		Owner:       req.Owner,
		OwnerID:     req.OwnerID,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Address:     req.Address,
		Images:      images,
	}
	err := dao.CreateHouse(db, &house)
	if err != nil {
		return &model.Error{StatusCode: codes.CreateHouseError, Message: "创建房屋记录失败"}
	}

	// 第三步: 将图片记录与房屋记录关联
	// 注意: 由于之前已经在创建房屋记录时关联了图片，此步骤可能不必要，除非需要额外的处理
	for _, img := range images {
		img.HouseID = house.ID // 确保每张图片的 HouseID 更新为新创建的房屋 ID
		apiError := dao.CreateHouseImages(db, []model.HouseImage{img})
		if apiError != nil {
			return &model.Error{StatusCode: codes.CreateHouseImageError, Message: "保存房屋图片失败"}
		}
	}
	return nil
}

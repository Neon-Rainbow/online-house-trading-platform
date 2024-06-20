package logic

import (
	"gorm.io/gorm"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
)

func GetViewingRecords(db *gorm.DB, idUint uint, pageSize int, pageNum int) ([]model.ViewingRecords, *model.Error) {
	viewingRecords, err := dao.GetViewingRecordsByUserID(db, idUint, pageSize, pageNum)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetViewingRecordsError}
	}
	return viewingRecords, nil
}

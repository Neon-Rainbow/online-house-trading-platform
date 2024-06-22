package logic

import (
	"gorm.io/gorm"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
)

func GetViewingRecords(db *gorm.DB, idUint uint, pageSize int, pageNum int) ([]model.ViewingRecords, int64, *model.Error) {
	viewingRecords, totalCounts, err := dao.GetViewingRecordsByUserID(db, idUint, pageSize, pageNum)
	if err != nil {
		return nil, 0, &model.Error{StatusCode: codes.GetViewingRecordsError}
	}
	return viewingRecords, totalCounts, nil
}

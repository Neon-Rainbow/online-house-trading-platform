package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
)

func GetViewingRecords(idUint uint, pageSize int, pageNum int) ([]model.ViewingRecords, int64, *model.Error) {
	viewingRecords, totalCounts, err := dao.GetViewingRecordsByUserID(idUint, pageSize, pageNum)
	if err != nil {
		return nil, 0, &model.Error{StatusCode: codes.GetViewingRecordsError}
	}
	return viewingRecords, totalCounts, nil
}

package logic

import (
	"fmt"
	"online-house-trading-platform/pkg/model"
	"os"
	"path/filepath"
	"time"

	"github.com/xuri/excelize/v2"
)

// ExportLoginRecordsToExcel 将登录记录导出到 Excel 文件
func ExportLoginRecordsToExcel(records []model.LoginRecord) string {
	// Create a new Excel file
	f := excelize.NewFile()

	// Create table headers
	headers := []string{"Login Time", "Login IP", "Login Method", "Address", "Operator"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue("Sheet1", cell, header)
	}

	// Fill data
	for i, record := range records {
		row := i + 2 // Data starts from the second row、
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), record.LoginTime.Format(time.RFC3339))
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), record.LoginIp)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), record.LoginMethod)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), record.Address)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), record.Operator)
	}

	dst := fmt.Sprintf("./download/%v/%s.xlsx", records[0].UserId, generateRandomFileName())
	dir := filepath.Dir(dst)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
		return ""
	}

	if err := f.SaveAs(dst); err != nil {
		fmt.Println("Error saving Excel file:", err)
		return ""
	} else {
		fmt.Println("Excel file saved successfully.")
	}
	return dst
}

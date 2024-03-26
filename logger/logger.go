package logger

import (
	"log"
	"os"
)

func InitLogger(logFilePath string) *os.File {
	// 打开日志文件
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("打开日志文件失败: %v", err)
	}
	log.SetOutput(logFile)
	log.Printf("日志文件打开成功,日志文件路径: %v", logFilePath)
	return logFile
}

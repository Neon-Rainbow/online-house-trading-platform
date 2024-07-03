package controller

import (
	"io/ioutil"
	"online-house-trading-platform/codes"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func DeleteLogFile(c *gin.Context) {
	errorChannel := make(chan error, 2)
	doneChannel := make(chan bool, 2)

	go func() {
		err := clearFileContent("./application.log")
		if err != nil {
			errorChannel <- err
		}
		doneChannel <- true
	}()

	go func() {
		err := clearFileContent("./formatted_application.log")
		if err != nil {
			errorChannel <- err
		}
		doneChannel <- true
	}()

	go func() {
		<-doneChannel
		<-doneChannel
		close(errorChannel)
	}()

	select {
	case err := <-errorChannel:
		if err != nil {
			zap.L().Error("Failed to clear log file", zap.Error(err))
			ResponseErrorWithCode(c, codes.LoginServerBusy)
			return
		}
	case <-time.After(10 * time.Second):
		zap.L().Error("DeleteLogFile exceeded time limit")
		ResponseTimeout(c)
		return
	}

	ResponseSuccess(c, nil)
	return
}

func clearFileContent(filePath string) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	zap.L().Info("Clearing log file", zap.String("file", filePath), zap.ByteString("content", content))

	err = ioutil.WriteFile(filePath, []byte{}, 0644)
	if err != nil {
		return err
	}

	return nil
}

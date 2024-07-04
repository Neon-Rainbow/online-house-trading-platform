package controller

import (
	"context"
	"errors"
	"io/ioutil"
	"online-house-trading-platform/codes"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func DeleteLogFile(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	errorChannel := make(chan error, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := clearFileContent(ctx, "./application.log")
		if err != nil {
			errorChannel <- err
			return
		}
		return
	}()

	go func() {
		defer wg.Done()
		err := clearFileContent(ctx, "./formatted_application.log")
		if err != nil {
			errorChannel <- err
			return
		}
		return
	}()

	go func() {
		wg.Wait()
		close(errorChannel)
	}()

	select {
	case err := <-errorChannel:
		if err != nil {
			zap.L().Error("Failed to clear log file", zap.Error(err))
			ResponseErrorWithCode(c, codes.LoginServerBusy)
			return
		}
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			zap.L().Error("DeleteLogFile exceeded time limit")
			ResponseTimeout(c)
			return
		} else {
			zap.L().Error("DeleteLogFile context error", zap.Error(ctx.Err()))
			ResponseErrorWithCode(c, codes.LoginServerBusy)
			return
		}
	}

	ResponseSuccess(c, nil)
	return
}

func clearFileContent(ctx context.Context, filePath string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

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

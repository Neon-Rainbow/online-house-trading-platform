APP_NAME := online_house_trading_platform
BUILD_DIR := build
VERSION := 1.0.0

all: clean build-macos build-windows package

clean:
	rm -rf $(BUILD_DIR)

build-macos:
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(APP_NAME)_macos main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)_windows.exe main.go

package: build-macos build-windows
	# Package macOS binary
	tar -cvzf $(BUILD_DIR)/$(APP_NAME)_$(VERSION)_darwin_arm64.tar.gz -C $(BUILD_DIR) $(APP_NAME)_macos
	# Package Windows binary
	zip $(BUILD_DIR)/$(APP_NAME)_$(VERSION)_windows_amd64.zip $(BUILD_DIR)/$(APP_NAME)_windows.exe

.PHONY: clean build-macos build-windows package
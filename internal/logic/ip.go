package logic

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetIPInformationWithContext 通过 IP 获取地址和运营商信息（带上下文支持）
func GetIPInformationWithContext(ctx context.Context, ip string) (address string, operator string) {
	url := fmt.Sprintf("http://cip.cc/%s", ip)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Printf("Failed to create request: %v", err)
		return "", ""
	}

	// 设置自定义的 User-Agent 头
	req.Header.Set("User-Agent", "curl/7.64.1")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to send GET request: %v", err)
		return "", ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("GET request failed with status: %s", resp.Status)
		return "", ""
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			fmt.Printf("Request cancelled or timed out")
			return "", ""
		default:
			line := scanner.Text()
			if strings.HasPrefix(line, "地址") {
				address = strings.TrimSpace(strings.Split(line, ":")[1])
			} else if strings.HasPrefix(line, "运营商") {
				operator = strings.TrimSpace(strings.Split(line, ":")[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read response body: %v", err)
		return "", ""
	}

	return address, operator
}

// GetIPInformation 通过 IP 获取地址和运营商信息
func GetIPInformation(ip string) (address string, operator string) {
	url := fmt.Sprintf("http://cip.cc/%s", ip)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// 设置自定义的 User-Agent 头
	req.Header.Set("User-Agent", "curl/7.64.1")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("GET request failed with status: %s", resp.Status)
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println("line:", line)
		if strings.HasPrefix(line, "地址") {
			address = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.HasPrefix(line, "运营商") {
			operator = strings.TrimSpace(strings.Split(line, ":")[1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	return
}

package main

import (
	"fmt"
	"log"
	"online-house-trading-platform/api"
	_ "online-house-trading-platform/pkg/model"
)

func main() {
	fmt.Println("Hello, World!")
	router := api.SetupRouter()
	
	err := router.Run(":8080")
	if err != nil {
		log.Printf("Server start failed: %v", err)
	}
}

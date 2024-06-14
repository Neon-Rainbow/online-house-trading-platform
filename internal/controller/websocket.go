package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
)

func WebsocketHandler(c *gin.Context) {
	id := c.Query("id")
	toId := c.Query("toId")
	connect, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(c.Writer, c.Request, nil) // 升级协议
	if err != nil {
		ResponseErrorWithCode(c, codes.LoginServerBusy)
		return
	}
	client := &logic.Client{
		ID:     logic.CreateID(id, toId), // 1->2
		SendID: logic.CreateID(toId, id), // 2->1
		Socket: connect,
		Send:   make(chan []byte),
	}
	// 用户注册到客户端管理器
	logic.Manager.Register <- client
	go client.Read()
	go client.Write()
}

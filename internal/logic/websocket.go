package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/redis"
)

type SendMsg struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type ReplyMsg struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}

type Client struct {
	ID     string          // 1->2 发送方->接收方
	SendID string          // 2->1 接收方->发送方
	Socket *websocket.Conn // websocket连接
	Send   chan []byte     // 发送消息
}

type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}

type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

var Manager = ClientManager{
	Clients:    make(map[string]*Client),
	Broadcast:  make(chan *Broadcast),
	Register:   make(chan *Client),
	Reply:      make(chan *Client),
	Unregister: make(chan *Client),
}

func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-Manager.Register:
			manager.handleRegister(conn)
		case conn := <-Manager.Unregister:
			manager.handleUnregister(conn)
		case broadcast := <-Manager.Broadcast:
			manager.handleBroadcast(broadcast)
		}
	}
}

func (manager *ClientManager) handleRegister(conn *Client) {
	log.Printf("建立新连接: %v", conn.ID)
	Manager.Clients[conn.ID] = conn
	replyMsg := &ReplyMsg{
		Code:    codes.WebsocketSuccess.Int(),
		Content: "已连接至服务器",
	}
	msg, _ := json.Marshal(replyMsg)
	_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
}

func (manager *ClientManager) handleUnregister(conn *Client) {
	log.Printf("连接失败: %v", conn.ID)
	if _, ok := Manager.Clients[conn.ID]; ok {
		replyMsg := &ReplyMsg{
			Code:    codes.WebsocketEnd.Int(),
			Content: "连接已断开",
		}
		msg, _ := json.Marshal(replyMsg)
		_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		close(conn.Send)
		delete(Manager.Clients, conn.ID)
	}
}

func (manager *ClientManager) handleBroadcast(broadcast *Broadcast) {
	message := broadcast.Message
	sendID := broadcast.Client.SendID
	flag := false // 默认对方不在线
	for id, conn := range Manager.Clients {
		if id != sendID {
			continue
		}
		select {
		case conn.Send <- message:
			flag = true
		default:
			close(conn.Send)
			delete(Manager.Clients, conn.ID)
		}
	}
	id := broadcast.Client.ID
	if flag {
		manager.sendReply(broadcast.Client, codes.WebsocketOnlineReply.Int(), "对方在线应答")
		err := dao.InsertMsg(redis.Rdb, id, string(message), 1, int64(3*time.Hour*24*30))
		if err != nil {
			fmt.Println("InsertOneMsg Err", err)
		}
	} else {
		manager.sendReply(broadcast.Client, codes.WebsocketOfflineReply.Int(), "对方不在线应答")
		err := dao.InsertMsg(redis.Rdb, id, string(message), 0, int64(3*time.Hour*24*30))
		if err != nil {
			fmt.Println("InsertOneMsg Err", err)
		}
	}
}

func (manager *ClientManager) sendReply(client *Client, code int, content string) {
	replyMsg := &ReplyMsg{
		Code:    code,
		Content: content,
	}
	msg, _ := json.Marshal(replyMsg)
	_ = client.Socket.WriteMessage(websocket.TextMessage, msg)
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		c.Socket.PongHandler()
		sendMsg := new(SendMsg)
		err := c.Socket.ReadJSON(&sendMsg)
		if err != nil {
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}
		if sendMsg.Type == 1 { // 1->2 发送方->接收方
			r1, _ := redis.Rdb.Get(context.Background(), c.ID).Result()
			r2, _ := redis.Rdb.Get(context.Background(), c.SendID).Result()
			if r1 > "3" && r2 == "" {
				reply := ReplyMsg{
					Code:    codes.WebsocketLimit.Int(),
					Content: "达到发送限制",
				}
				msg, _ := json.Marshal(reply)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
				continue
			} else {
				redis.Rdb.Incr(context.Background(), c.ID)                                      // 1->2 发送方->接收方 发送次数+1
				_, _ = redis.Rdb.Expire(context.Background(), c.ID, time.Hour*24*30*3).Result() // 1->2 发送方->接收方 防止过快分手,设置三个月过期时间

			}
			Manager.Broadcast <- &Broadcast{
				Client:  c,
				Message: []byte(sendMsg.Content),
			}
		} else if sendMsg.Type == 2 { //拉取历史消息

		} else if sendMsg.Type == 3 {

		}
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			replyMsg := ReplyMsg{
				Code:    codes.WebsocketSuccessMessage.Int(),
				Content: fmt.Sprintf("%s", msg),
			}
			msg, _ = json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

func CreateID(id string, toId string) string {
	return id + "->" + toId // 1->2
}

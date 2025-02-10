package websockets

import (
	wsmodel "chatross-api/internal/model/ws_model"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type HubFunc interface{
	Run()
	SendMessage(msg *wsmodel.Message)
	PingPong(conn *websocket.Conn, client *Client)
}

type Hub struct {
	HubFunc
	Clients map[string]*Client
	register chan *Client
	unregister chan *Client
	broadcast chan *wsmodel.Message
	mu 			sync.Mutex
}


func NewHub() *Hub {
	return &Hub{
		broadcast: make(chan *wsmodel.Message),
		Clients: make(map[string]*Client),
		register: make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.mu.Lock()
			hub.Clients[client.ID] = client
			hub.mu.Unlock()

		case client := <-hub.unregister:
			hub.mu.Lock()
			delete(hub.Clients, client.ID)
			hub.mu.Unlock()
		case msg := <-hub.broadcast:
			hub.SendMessage(msg)
		}
	}
}

func (hub *Hub) PingPong(conn *websocket.Conn, client *Client) {
	ticker := time.NewTicker(300 * time.Millisecond)
	defer func ()  {
		ticker.Stop()
		if len(client.Conn) == 0{
			hub.unregister <- client
		}
	}()

	for range ticker.C {
		err := conn.WriteMessage(websocket.PingMessage, nil)
		if err != nil {
			log.Println("Ping failed, client disconnected:", err)
			client.RemoveConnection(conn)
			return
		}
	}
	
}

func (hub *Hub) SendMessage(msg *wsmodel.Message) {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	recipient, exist := hub.Clients[msg.To]
	if !exist {
		return
	}
	
	recipient.Send <- msg
}
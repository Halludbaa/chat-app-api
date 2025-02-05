package websockets

import (
	wsmodel "chatross-api/internal/model/ws_model"
	"sync"
)

type Hub struct {
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

func (hub *Hub) SendMessage(msg *wsmodel.Message) {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	recipient, exist := hub.Clients[msg.To]
	if !exist {
		return
	}
	
	recipient.Send <- msg
}
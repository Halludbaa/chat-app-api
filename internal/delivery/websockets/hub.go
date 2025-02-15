package websockets

import (
	wsmodel "chatross-api/internal/model/wsmodel"
	"chatross-api/internal/usecase"
	"context"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type HubFunc interface{
	Run()
	SendMessage(msg *wsmodel.Message)
	PingPong(conn *websocket.Conn, client *Client)
}

type UseCaseList struct {
	ChatUsecase 	*usecase.ChatUsecase
	MessageUsecase 	*usecase.MessageUsecase
}

type Hub struct {
	HubFunc
	log 		*logrus.Logger
	UC			*UseCaseList
	Clients 	map[string]*Client
	register 	chan *Client
	unregister 	chan *Client
	broadcast 	chan *wsmodel.Message
	mu 			sync.Mutex
}


func NewHub(log *logrus.Logger, uc *UseCaseList) *Hub {
	return &Hub{
		UC: uc,
		broadcast: make(chan *wsmodel.Message),
		Clients: make(map[string]*Client),
		register: make(chan *Client),
		unregister: make(chan *Client),
		log: log,
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
			sMsg, err := hub.StoreMessage(msg)
			if err == nil {
				hub.SendMessage(sMsg)
			}
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
	
	sender, exist := hub.Clients[msg.From]
	if !exist {
		return
	}

	sender.Send <- msg

	recipient, exist := hub.Clients[msg.To]
	if !exist {
		return
	}
	
	recipient.Send <- msg

	
}

func (hub *Hub) StoreMessage(msg *wsmodel.Message) (*wsmodel.Message, error) {
	hub.log.Debug("Stored")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if msg.ChatID == 0 {
		chatID, err := hub.UC.ChatUsecase.GetChatID(ctx, msg)
		if err != nil || chatID == 0 {
			return nil, err
		}

		msg.ChatID = chatID
	}
	
	// Call Store Usecase of Messages
	newMsg, err := hub.UC.MessageUsecase.Store(ctx, msg)
	if  err != nil {
		hub.log.WithError(err).Error("Failure in Database")
		return nil, err
	}
	newMsg.Type = msg.Type

	return newMsg, nil
}
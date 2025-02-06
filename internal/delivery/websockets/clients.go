package websockets

import (
	wsmodel "chatross-api/internal/model/ws_model"
	"log"

	"github.com/gorilla/websocket"
)
type ClientContract interface {
	ReadMessage()
	WriteMessage()
}

type Client struct {
	ID string
	Conn []*websocket.Conn
	Send chan *wsmodel.Message
	Hub *Hub
}

func NewClient(client *Client) ClientContract{
	return client
}

func (c *Client) ReadMessage() {
	defer func() {
		for _, conn := range c.Conn {
			conn.Close()
		}
		c.Hub.unregister <- c
	}()

	
	for { for _, conn := range c.Conn {
		msg := new(wsmodel.Message)
		if err := conn.ReadJSON(msg); err != nil {
			log.Println("Read Error: ", err)
			break
		}

		log.Printf("Read Success: From %s To %s , Message: %s \n", msg.From, msg.To, msg.Content)
		c.Hub.broadcast <- msg

	}}
}

func (c *Client) WriteMessage() {
	defer func() {
		for _, conn := range c.Conn {
			conn.Close()
		}
	}()
	for msg := range c.Send {
		for _, conn := range c.Conn {
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Println("Write Error: ", err)
				break
			}
		}
	}
}
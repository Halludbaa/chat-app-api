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
	Conn *websocket.Conn
	Send chan *wsmodel.Message
	Hub *Hub
}

func (c *Client) ReadMessage() {
	defer func() {
		c.Conn.Close()
		c.Hub.unregister <- c
	}()

	for {
		msg := new(wsmodel.Message)
		if err := c.Conn.ReadJSON(msg); err != nil {
			log.Println("Read Error: ", err)
			break
		}

		log.Println("Read Success: ", msg.Content)

		c.Hub.broadcast <- msg
	}
}

func (c *Client) WriteMessage() {
	defer c.Conn.Close()
	for msg := range c.Send {
		err := c.Conn.WriteJSON(msg)
		if err != nil {
			log.Println("Write Error: ", err)
			break
		}
	}
}
package websockets

import (
	wsmodel "chatross-api/internal/model/ws_model"
	"log"
	"slices"

	"github.com/gorilla/websocket"
)
type ClientContract interface {
	ReadMessage(conn *websocket.Conn)
	WriteMessage()
	RemoveConnection(conn *websocket.Conn)
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

func (c *Client) ReadMessage(conn *websocket.Conn) {
	defer func() {
		c.RemoveConnection(conn)
		if len(c.Conn) == 0 {
			c.Hub.unregister <- c
		}
	}()

	for  {
		msg := new(wsmodel.Message)
		if err := conn.ReadJSON(msg); err != nil {
			log.Println("Read Error: ", err)
			
			break
		}
		log.Printf("Read Success: From %s To %s , Message: %s \n", msg.From, msg.To, msg.Content)
		c.Hub.broadcast <- msg
	}
}
func (c *Client) RemoveConnection(conn *websocket.Conn){
	conn.Close()
	if idx := slices.Index(c.Conn, conn); idx != -1 {
		c.Conn = append(c.Conn[:idx], c.Conn[idx+1:]... )
	}
}

func (c *Client) WriteMessage() {
	defer func() {
		log.Println("Write Off")
	}()

	for msg := range c.Send {
		if  num := len(c.Conn); num == 0 {
			log.Printf("connection: %d", num )
			break
		}
		for _, conn := range c.Conn {
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Println("Write Error: ", err)
				c.RemoveConnection(conn)
				continue
			}
		}
		
	}
}
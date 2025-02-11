package websockets

import (
	wsmodel "chatross-api/internal/model/wsmodel"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}	

func ServeWS(hub *Hub, ctx *gin.Context, userID *string){

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Websocket Failed to Serve: ", err)
		return
	}

	if client, exist := hub.Clients[*userID]; exist {
		client.Conn = append(client.Conn, conn)
		go client.ReadMessage(conn)
		return
	}

	client := &Client{
		ID: *userID,
		Conn: []*websocket.Conn{
			conn,
		},
		Send: make(chan *wsmodel.Message),
		Hub: hub,
	}

	hub.register <- client

	go client.ReadMessage(conn)
	go client.WriteMessage()
} 
package routes

import (
	"fmt"
	"github.com/devproje/plog/log"
	"github.com/devproje/simple-chat/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func ws(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Errorln("Failed to upgrade to WebSocket:", err)
		return
	}
	defer func(conn *websocket.Conn) {
		err = conn.Close()
		if err != nil {

		}
	}(conn)

	socket := &model.User{Sock: conn, Name: ""}
	clients[socket] = true
	log.Infof("connection established to chat server with addr: %s\n", conn.RemoteAddr().String())

	for {
		var recv model.MessageData
		err = conn.ReadJSON(&recv)
		if err != nil {
			delete(clients, socket)
			break
		}

		broadcast <- recv
	}
}

func HandleBroadcasts() {
	for {
		recv := <-broadcast

		for client := range clients {
			switch recv.Type {
			case "set_nickname":
				client.Name = recv.Payload
				err := client.Sock.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s has joined the chat room!", recv.Payload)))
				if err != nil {
					_ = client.Sock.Close()
					delete(clients, client)
				}
			case "new_message":
				fmt.Println(client.Name)
				err := client.Sock.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s: %s", client.Name, recv.Payload)))
				if err != nil {
					_ = client.Sock.Close()
					delete(clients, client)
				}
			}
		}
	}
}

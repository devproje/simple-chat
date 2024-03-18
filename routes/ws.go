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
		log.Errorln(err)
		return
	}

	client := &model.User{Sock: conn}
	clients[conn] = client

	for {
		var msg string
		var recv model.MessageData
		err = conn.ReadJSON(&recv)
		if err != nil {
			delete(clients, conn)
			break
		}

		switch recv.Type {
		case "new_message":
			msg = fmt.Sprintf("%s: %s", client.Name, recv.Payload)
			recv.Payload = msg
		case "set_username":
			client.Name = recv.Payload
			msg = fmt.Sprintf("%s joined the chat.", clients[conn].Name)
			recv.Payload = msg
		}

		broadcast <- recv
	}
}

func HandleConnections() {
	for {
		recv := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(recv.Payload))
			if err != nil {
				_ = client.Close()
				delete(clients, client)
			}
		}
	}
}

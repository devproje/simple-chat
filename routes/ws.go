package routes

import (
	"github.com/devproje/plog/log"
	"github.com/devproje/simple-chat/model"
	"github.com/gin-gonic/gin"
)

func ws(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Errorln(err)
		return
	}

	client := model.NewUser()
	clients[conn] = client

	for {
		var recv model.MessageData
		err = conn.ReadJSON(&recv)
		if err != nil {
			if client.Name != "" {
				broadcast <- model.MessageData{
					Type:    "left_user",
					Payload: client.Name,
				}
			}

			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			break
		}

		switch recv.Type {
		case "new_message":
			recv.Author = client.Name
		case "set_username":
			client.Name = recv.Payload
			recv.Payload = client.Name
		}

		broadcast <- recv
	}
}

func HandleConnections() {
	for {
		recv := <-broadcast
		for client := range clients {
			err := client.WriteJSON(recv)
			if err != nil {
				_ = client.Close()
				mutex.Lock()
				delete(clients, client)
				mutex.Unlock()
			}
		}
	}
}

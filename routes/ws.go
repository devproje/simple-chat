package routes

import (
	"github.com/devproje/plog/log"
	"github.com/devproje/simple-chat/config"
	"github.com/devproje/simple-chat/controller"
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
					Type:    model.LeftUser,
					Payload: client.Name,
				}
			}

			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			break
		}

		switch recv.Type {
		case model.NewMessage:
			recv.Author = client.Name
		case model.SetUsername:
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
			if len(recv.Payload) > config.Load().ContentLength {
				continue
			}

			if config.Load().Logging {
				controller.Logging(&model.Log{
					Type:    recv.Type.ToString(),
					Author:  clients[client].Name,
					Content: recv.Payload,
				})
			}

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

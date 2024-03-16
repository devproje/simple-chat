package routes

import (
	"encoding/json"

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
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Errorln(err)
		}

		var recv model.MessageData
		json.Unmarshal(msg, &recv)

		var data = model.MessageData{
			Type:    "recv_message",
			Payload: recv.Payload,
		}

		raw, err := json.Marshal(&data)
		if err != nil {
			log.Errorln(err)
			return
		}

		switch data.Type {
		case "new_message":
			log.Debugln(raw)
			err := conn.WriteMessage(websocket.TextMessage, []byte(raw))
			if err != nil {
				log.Errorln(err)
				return
			}
		}
	}
}

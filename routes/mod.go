package routes

import (
	"github.com/devproje/simple-chat/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	clients   = make(map[*websocket.Conn]*model.User)
	broadcast = make(chan model.MessageData)
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func Build(app *gin.Engine) {
	app.GET("/ws", ws)
}

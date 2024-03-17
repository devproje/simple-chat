package routes

import (
	"github.com/devproje/simple-chat/model"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	mutex     sync.Mutex
	clients   = make(map[*model.User]bool)
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

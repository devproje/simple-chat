package routes

import (
	"github.com/devproje/simple-chat/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var (
	mutex     sync.Mutex
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
	v1 := app.Group("/v1")
	{
		v1.GET("users", users)
	}
}

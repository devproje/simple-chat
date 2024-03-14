package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/devproje/plog/log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	port     int
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type MessageData struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

func init() {
	flag.IntVar(&port, "port", 3000, "service port")
	flag.Parse()
}

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
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

			var data MessageData
			json.Unmarshal(msg, &data)

			log.Infof("{\"type\": \"%s\", \"payload\":\"%s\"}", data.Type, data.Payload)
		}
	})

	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln("Failed to start server:", err)
	}
}

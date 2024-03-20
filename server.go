package main

import (
	"flag"
	"fmt"
	"github.com/devproje/simple-chat/config"
	"github.com/devproje/simple-chat/database"
	"github.com/devproje/simple-chat/middleware"

	"github.com/devproje/plog/log"
	"github.com/devproje/simple-chat/routes"
	"github.com/gin-gonic/gin"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 3000, "service port")
	flag.Parse()
}

func main() {
	app := gin.Default()
	app.Use(middleware.CORS)
	routes.Build(app)
	go routes.HandleConnections()

	config.Load()
	if config.Load().Logging {
		err := database.Init()
		if err != nil {
			log.Errorln("database is not opened: %s", err.Error())
			return
		}
	}

	err := app.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln("Failed to start server:", err)
	}
}

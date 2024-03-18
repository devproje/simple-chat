package main

import (
	"flag"
	"fmt"

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
	routes.Build(app)
	go routes.HandleConnections()

	err := app.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln("Failed to start server:", err)
	}
}

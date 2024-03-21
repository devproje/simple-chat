package routes

import (
	"github.com/devproje/simple-chat/config"
	"github.com/gin-gonic/gin"
)

func server(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"name": config.Load().ServerName,
	})
}

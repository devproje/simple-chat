package routes

import (
	"github.com/devproje/simple-chat/config"
	"github.com/gin-gonic/gin"
)

func server(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"ok":             1,
		"status":         200,
		"name":           config.Load().ServerName,
		"content_length": config.Load().ContentLength,
	})
}

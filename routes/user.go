package routes

import (
	"github.com/devproje/simple-chat/model"
	"github.com/gin-gonic/gin"
)

func users(ctx *gin.Context) {
	var users []*model.User
	for _, client := range clients {
		users = append(users, client)
	}

	ctx.JSON(200, gin.H{
		"users": users,
	})
}

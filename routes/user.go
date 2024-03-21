package routes

import (
	"github.com/devproje/simple-chat/model"
	"github.com/gin-gonic/gin"
	"sort"
)

type Users []*model.User

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return u[i].Name < u[j].Name
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func users(ctx *gin.Context) {
	var u Users
	for _, client := range clients {
		if client.Name == "" {
			continue
		}

		u = append(u, client)
	}

	sort.Sort(u)
	ctx.JSON(200, gin.H{
		"len":   u.Len(),
		"users": u,
	})
}

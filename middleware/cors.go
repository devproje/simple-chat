package middleware

import "github.com/gin-gonic/gin"

func CORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
}

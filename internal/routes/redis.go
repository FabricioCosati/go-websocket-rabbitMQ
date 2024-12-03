package routes

import (
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/di"
	"github.com/gin-gonic/gin"
)

func InitRedis(r *gin.Engine, app *di.App) {
	handlers := app.RedisInit.RedisHandler

	r.GET("/auth/guest", handlers.ConnectUser)
}

package routes

import (
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/di"
	"github.com/gin-gonic/gin"
)

func InitWebsocket(r *gin.Engine, app *di.App) {
	handlers := app.WebsocketInit.WebsocketHandler

	r.GET("/ws", handlers.ConnectWs)
}

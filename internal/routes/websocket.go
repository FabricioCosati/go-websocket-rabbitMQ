package routes

import (
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/handlers"
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/websocket"
	"github.com/gin-gonic/gin"
)

func InitWebsocket(r *gin.Engine) {

	hub := websocket.NewHub()
	go hub.InitHub()

	r.GET("/ws", func(ctx *gin.Context) { handlers.ConnectWs(ctx, hub) })
}

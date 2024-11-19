package routes

import (
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitWebsocket(r *gin.Engine) {
	r.GET("/ws", handlers.ConnectWs)
}

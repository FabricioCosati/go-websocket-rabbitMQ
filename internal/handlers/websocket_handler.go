package handlers

import (
	"fmt"

	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/usecase"
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/websocket"
	"github.com/gin-gonic/gin"
)

func ConnectWs(ctx *gin.Context, hub *websocket.Hub) {
	err := usecase.ConnectWs(ctx, hub)
	if err != nil {
		fmt.Printf("error on connect WS: %s", err)
		ctx.JSON(500, "internal error")
		return
	}
}

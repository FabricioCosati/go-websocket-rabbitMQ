package usecase

import (
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/websocket"
	"github.com/gin-gonic/gin"
)

func ConnectWs(ctx *gin.Context, hub *websocket.Hub) error {
	client, err := websocket.InitClient(ctx, hub)
	if err != nil {
		return err
	}

	client.Register()

	go client.ReadMessages()
	go client.WriteMessages()

	return nil
}

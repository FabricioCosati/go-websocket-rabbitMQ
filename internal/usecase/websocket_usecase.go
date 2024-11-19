package usecase

import (
	"fmt"

	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/websocket"
	"github.com/gin-gonic/gin"
)

func ConnectWs(ctx *gin.Context) error {
	client, err := websocket.InitClient(ctx)
	if err != nil {
		return err
	}

	amqp, err := websocket.InitAmqp()
	if err != nil {
		fmt.Printf("error on init rabbitmq: %s", err)
		return err
	}

	err = amqp.InitExchange()
	if err != nil {
		fmt.Printf("error on init rabbitmq exchange: %s", err)
		return err
	}
	err = amqp.InitQueue()
	if err != nil {
		fmt.Printf("error on init rabbitmq queue: %s", err)
		return err
	}

	go client.SendMessage(ctx, amqp.Channel)
	go client.ReadMessage(amqp.Channel, amqp.Queue.Name)

	return nil
}

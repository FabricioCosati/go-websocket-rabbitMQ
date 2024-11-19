package usecase

import (
	"fmt"
	"sync"

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

	defer amqp.CloseAmqp()

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

	var wg sync.WaitGroup
	wg.Add(2)
	go client.SendMessage(ctx, amqp.Channel, &wg)
	go client.ReadMessage(amqp.Channel, amqp.Queue.Name, &wg)
	wg.Wait()

	return nil
}

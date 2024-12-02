package usecase

import (
	"fmt"
	"sync"

	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/services/websocket"
	"github.com/gin-gonic/gin"
)

type WebsocketUsecase interface {
	ConnectWs(ctx *gin.Context) error
}

type WebsocketUsecaseImpl struct {
	BrokerService websocket.WebsocketBrokerService
	ClientService websocket.WebsocketClientService
	Hub           *websocket.Hub
}

func (impl *WebsocketUsecaseImpl) ConnectWs(ctx *gin.Context) error {
	hub := impl.Hub
	brokerService := impl.BrokerService
	clientService := impl.ClientService

	client, err := clientService.InitClient(ctx, hub)
	if err != nil {
		return err
	}

	hub.RegisterClient(client)

	amqp, err := brokerService.InitAmqp()
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

	msgs, err := clientService.ConsumeQueueMessages(amqp.Channel, amqp.Queue.Name)
	if err != nil {
		fmt.Printf("error on consume queue messages: %s", err)
		return err
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go client.SendMessage(ctx, amqp.Channel, &wg)
	go client.ReadMessage(msgs, &wg)
	wg.Wait()

	return nil
}

func InitWebsocketUsecase(
	brokerService websocket.WebsocketBrokerService,
	clientService websocket.WebsocketClientService,
	hub *websocket.Hub,
) *WebsocketUsecaseImpl {
	return &WebsocketUsecaseImpl{
		BrokerService: brokerService,
		ClientService: clientService,
		Hub:           hub,
	}
}

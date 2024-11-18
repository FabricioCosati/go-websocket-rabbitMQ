package usecase

import (
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/websocket"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectWs(ctx *gin.Context, hub *websocket.Hub) error {
	client, err := websocket.InitClient(ctx, hub)
	if err != nil {
		return err
	}

	client.Register()

	ch, err := initExchange()
	if err != nil {
		return err
	}
	msgs, err := initConsumer()
	if err != nil {
		return nil
	}

	go client.InitReading(ctx, ch)
	go client.InitSending(msgs)

	return nil
}

type Teste struct {
	Message string
}

func initExchange() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		return &amqp.Channel{}, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return ch, err
	}

	err = ch.ExchangeDeclare(
		"chat",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return ch, err
	}

	return ch, nil
}

func initConsumer() (<-chan amqp.Delivery, error) {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		return make(<-chan amqp.Delivery), err
	}

	ch, err := conn.Channel()
	if err != nil {
		return make(<-chan amqp.Delivery), err
	}

	err = ch.ExchangeDeclare(
		"chat",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return make(<-chan amqp.Delivery), err
	}

	q, err := ch.QueueDeclare(
		"",
		true,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return make(<-chan amqp.Delivery), err
	}

	err = ch.QueueBind(
		q.Name,
		"",
		"chat",
		false,
		nil,
	)
	if err != nil {
		return make(<-chan amqp.Delivery), err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		return make(<-chan amqp.Delivery), err
	}

	return msgs, nil
}

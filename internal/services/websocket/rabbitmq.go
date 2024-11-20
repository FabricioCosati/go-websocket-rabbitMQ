package websocket

import (
	"fmt"

	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type WebsocketBrokerService interface {
	InitAmqp() (*AmqpClient, error)
}

type WebsocketBrokerServiceImpl struct {
	Config config.Config
}

type AmqpClient struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   *amqp.Queue
}

func (impl *WebsocketBrokerServiceImpl) InitAmqp() (*AmqpClient, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@localhost:5672/", impl.Config.BrokerUser, impl.Config.BrokerPass))
	if err != nil {
		return &AmqpClient{}, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return &AmqpClient{}, err
	}

	return &AmqpClient{
		Conn:    conn,
		Channel: ch,
		Queue:   &amqp.Queue{},
	}, nil
}

func (c *AmqpClient) CloseAmqp() {
	c.Conn.Close()
	c.Channel.Close()
}

func (c *AmqpClient) InitExchange() error {
	err := c.Channel.ExchangeDeclare(
		"chat",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *AmqpClient) InitQueue() error {
	q, err := c.Channel.QueueDeclare(
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = c.Channel.QueueBind(
		q.Name,
		"",
		"chat",
		false,
		nil,
	)
	if err != nil {
		return err
	}
	c.appendQueue(&q)
	return nil
}

func (c *AmqpClient) appendQueue(queue *amqp.Queue) {
	c.Queue = queue
}

func InitWebsocketBrokerService(cfg config.Config) *WebsocketBrokerServiceImpl {
	return &WebsocketBrokerServiceImpl{
		Config: cfg,
	}
}

package websocket

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Amqp struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   *amqp.Queue
}

func InitAmqp() (*Amqp, error) {
	user := os.Getenv("RABBITMQ_USER")
	pass := os.Getenv("RABBITMQ_PASSWORD")

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@localhost:5672/", user, pass))
	if err != nil {
		return &Amqp{}, nil
	}

	ch, err := conn.Channel()
	if err != nil {
		return &Amqp{}, err
	}

	return &Amqp{
		Conn:    conn,
		Channel: ch,
		Queue:   &amqp.Queue{},
	}, nil
}

func (a *Amqp) CloseAmqp() {
	a.Conn.Close()
	a.Channel.Close()
}

func (a *Amqp) InitExchange() error {
	err := a.Channel.ExchangeDeclare(
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

func (a *Amqp) InitQueue() error {
	q, err := a.Channel.QueueDeclare(
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

	err = a.Channel.QueueBind(
		q.Name,
		"",
		"chat",
		false,
		nil,
	)
	if err != nil {
		return err
	}
	a.appendQueue(&q)
	return nil
}

func (a *Amqp) appendQueue(queue *amqp.Queue) {
	a.Queue = queue
}

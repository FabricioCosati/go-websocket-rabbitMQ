package websocket

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *Client) SendMessage(ctx *gin.Context, ch *amqp.Channel) {
	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Printf("error on read messages: %s", err)
			break
		}

		ch.PublishWithContext(
			ctx,
			"chat",
			"",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        m,
			},
		)
	}
}

func (c *Client) ReadMessage(ch *amqp.Channel, qn string) {
	msgs, err := ch.Consume(
		qn,
		"",
		true,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("error on init consumer: %s", err)
	}

	for m := range msgs {
		err := c.Conn.WriteMessage(websocket.TextMessage, m.Body)
		if err != nil {
			fmt.Printf("error on write message: %s", err)
			break
		}
	}
}

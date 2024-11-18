package websocket

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *Client) ReadMessages() {
	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Printf("error on read messages: %s", err)
			break
		}

		c.Hub.Broadcast <- m
	}
}

func (c *Client) WriteMessages() {
	for {
		select {
		case m, ok := <-c.Send:
			if !ok {
				fmt.Printf("error on receive broadcast message")
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				break
			}

			err := c.Conn.WriteMessage(websocket.TextMessage, m)
			if err != nil {
				fmt.Printf("error on write message: %s", err)
				break
			}
		}
	}
}

func (c *Client) InitReading(ctx *gin.Context, ch *amqp.Channel) {
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

func (c *Client) InitSending(msgs <-chan amqp.Delivery) error {
	for m := range msgs {
		err := c.Conn.WriteMessage(websocket.TextMessage, m.Body)
		if err != nil {
			fmt.Printf("error on write message: %s", err)
			break
		}
	}

	return nil
}

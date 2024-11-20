package websocket

import (
	"time"

	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
)

type WebsocketClientService interface {
	InitClient(ctx *gin.Context) (*Client, error)
	ConsumeQueueMessages(ctch *amqp.Channel, qn string) (<-chan amqp.Delivery, error)
}

type WebsocketClientServiceImpl struct{}

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:   512,
	WriteBufferSize:  512,
	HandshakeTimeout: time.Second * 10,
	CheckOrigin:      websocket.IsWebSocketUpgrade,
}

func (impl *WebsocketClientServiceImpl) InitClient(ctx *gin.Context) (*Client, error) {
	var client Client
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, ctx.Request.Trailer)
	if err != nil {
		return &client, err
	}

	client.Conn = conn
	client.Send = make(chan []byte, 256)

	return &client, nil
}

func (c *Client) SendMessage(ctx *gin.Context, ch *amqp.Channel, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		c.Conn.Close()
	}()

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

func (c *Client) ReadMessage(msgs <-chan amqp.Delivery, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		c.Conn.Close()
	}()

	for m := range msgs {
		err := c.Conn.WriteMessage(websocket.TextMessage, m.Body)
		if err != nil {
			fmt.Printf("error on write message: %s", err)
			break
		}
	}
}

func (impl *WebsocketClientServiceImpl) ConsumeQueueMessages(ch *amqp.Channel, qn string) (<-chan amqp.Delivery, error) {
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
		return msgs, err
	}

	return msgs, nil
}

func InitWebsocketClientService() WebsocketClientService {
	return &WebsocketClientServiceImpl{}
}

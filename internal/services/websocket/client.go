package websocket

import (
	"encoding/json"
	"time"

	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
)

type WebsocketClientService interface {
	InitClient(ctx *gin.Context, hub *Hub) (*Client, error)
	ConsumeQueueMessages(ctch *amqp.Channel, qn string) (<-chan amqp.Delivery, error)
}

type WebsocketClientServiceImpl struct{}

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
	Hub  *Hub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:   512,
	WriteBufferSize:  512,
	HandshakeTimeout: time.Second * 10,
	CheckOrigin:      websocket.IsWebSocketUpgrade,
}

type Message struct {
	User    User
	Message string
	Time    string
}

type User struct {
	Name  string
	Photo string
}

func (impl *WebsocketClientServiceImpl) InitClient(ctx *gin.Context, hub *Hub) (*Client, error) {
	var client Client
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, ctx.Request.Trailer)
	if err != nil {
		return &client, err
	}

	client.Conn = conn
	client.Send = make(chan []byte, 256)
	client.Hub = hub

	return &client, nil
}

func (c *Client) SendMessage(ctx *gin.Context, ch *amqp.Channel, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		c.Hub.UnRegisterClient(c)
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error on read messages: %s", err)
			}
			break
		}

		var body Message
		err = json.Unmarshal(m, &body)
		if err != nil {
			fmt.Printf("error on decode body: %s", err)
			break
		}

		body.Time = time.Now().Format("15:04")
		body.User.Photo = "guest.png"

		m, err = json.Marshal(body)
		if err != nil {
			fmt.Printf("error on encode body: %s", err)
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
		w, err := c.Conn.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}

		w.Write(m.Body)
		err = w.Close()

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

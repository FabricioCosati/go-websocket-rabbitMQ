package websocket

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

type Message struct {
	Message string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:   512,
	WriteBufferSize:  512,
	HandshakeTimeout: time.Second * 10,
	CheckOrigin:      websocket.IsWebSocketUpgrade,
}

func InitClient(ctx *gin.Context, hub *Hub) (*Client, error) {
	var client Client
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, ctx.Request.Trailer)
	if err != nil {
		return &client, err
	}

	client.Conn = conn
	client.Hub = hub
	client.Send = make(chan []byte, 256)

	return &client, nil
}

func (c *Client) Register() {
	c.Hub.Register <- c
}

func (c *Client) Unregister() {
	c.Hub.Unregister <- c
}

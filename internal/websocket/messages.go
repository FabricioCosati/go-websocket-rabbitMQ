package websocket

import (
	"fmt"

	"github.com/gorilla/websocket"
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

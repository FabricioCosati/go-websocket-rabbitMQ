package main

import (
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	routes.InitWebsocket(r)

	r.Run(":8080")
}

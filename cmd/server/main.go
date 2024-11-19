package main

import (
	"log"

	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/config"
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Fatal("error on init config: %s", err)
	}

	r := gin.New()
	routes.InitWebsocket(r)
	r.Run(":8080")
}

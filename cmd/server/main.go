package main

import (
	"log"

	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/config"
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/di"
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("error on init config: %s", err)
	}

	app, err := di.InitApp(cfg)

	r := gin.New()
	routes.InitWebsocket(r, app)
	r.Run(":8080")
}

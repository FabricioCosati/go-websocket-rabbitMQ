package handlers

import (
	"fmt"

	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/usecase"
	"github.com/gin-gonic/gin"
)

func ConnectWs(ctx *gin.Context) {
	err := usecase.ConnectWs(ctx)
	if err != nil {
		fmt.Printf("error on connect WS: %s", err)
		ctx.JSON(500, "internal server error")
		return
	}
}

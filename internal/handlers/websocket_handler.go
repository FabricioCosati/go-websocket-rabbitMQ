package handlers

import (
	"fmt"

	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/usecase"
	"github.com/gin-gonic/gin"
)

type WebsocketHandler interface {
	ConnectWs(ctx *gin.Context)
}

type WebsocketHandlerImpl struct {
	Usecase usecase.WebsocketUsecase
}

func (impl *WebsocketHandlerImpl) ConnectWs(ctx *gin.Context) {
	err := impl.Usecase.ConnectWs(ctx)
	if err != nil {
		fmt.Printf("error on connect WS: %s", err)
		ctx.JSON(500, "internal server error")
		return
	}
}

func InitWebsocketHandler(websocketUsecase usecase.WebsocketUsecase) *WebsocketHandlerImpl {
	return &WebsocketHandlerImpl{
		Usecase: websocketUsecase,
	}
}

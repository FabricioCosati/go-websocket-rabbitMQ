package handlers

import (
	"fmt"

	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/usecase"
	"github.com/gin-gonic/gin"
)

type RedisHandler interface {
	ConnectUser(ctx *gin.Context)
}

type RedisHandlerImpl struct {
	Usecase usecase.RedisUsecase
}

func (impl *RedisHandlerImpl) ConnectUser(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	value, err := impl.Usecase.GetUser(ctx)
	if err == nil {
		ctx.JSON(200, value)
		return
	}

	err = impl.Usecase.SetUser(ctx)
	if err != nil {
		fmt.Printf("error on set client user: %s", err)
		ctx.JSON(500, "internal server error")
		return
	}

	value, err = impl.Usecase.GetUser(ctx)
	if err != nil {
		fmt.Printf("error on get client user: %s", err)
		ctx.JSON(500, "internal server error")
		return
	}

	ctx.JSON(200, value)
}

func InitRedisHandler(redisUsecase usecase.RedisUsecase) *RedisHandlerImpl {
	return &RedisHandlerImpl{
		Usecase: redisUsecase,
	}
}

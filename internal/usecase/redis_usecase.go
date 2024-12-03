package usecase

import (
	"fmt"

	r "github.com/FabricioCosati/go-websocket-rabbitMQ/internal/services/redis"
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/services/websocket"
	"github.com/gin-gonic/gin"
)

type RedisUsecase interface {
	SetUser(ctx *gin.Context) error
	GetUser(ctx *gin.Context) (interface{}, error)
}

type RedisUsecaseImpl struct {
	RedisClient *r.RedisClient
	Hub         *websocket.Hub
}

func (impl *RedisUsecaseImpl) SetUser(ctx *gin.Context) error {
	redisClient := impl.RedisClient
	hub := impl.Hub

	var user websocket.User
	user.Id = len(hub.Clients)
	user.Name = fmt.Sprintf("Guest %d", user.Id)
	user.Photo = "guest.png"

	err := redisClient.Set(ctx, fmt.Sprintf("user%d", user.Id), user)
	if err != nil {
		return err
	}

	return nil
}

func (impl *RedisUsecaseImpl) GetUser(ctx *gin.Context) (interface{}, error) {
	redisClient := impl.RedisClient
	hub := impl.Hub

	return redisClient.Get(ctx, fmt.Sprintf("user%d", len(hub.Clients)))
}

func InitRedisUsecase(redisClient *r.RedisClient, hub *websocket.Hub) *RedisUsecaseImpl {
	return &RedisUsecaseImpl{
		RedisClient: redisClient,
		Hub:         hub,
	}
}

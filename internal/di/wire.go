//go:build wireinject
// +build wireinject

package di

import (
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/config"
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/services/redis"
	"github.com/google/wire"
)

type App struct {
	WebsocketInit *WebsocketInit
	RedisInit     *RedisInit
}

func NewApp(websocketInit *WebsocketInit, redisInit *RedisInit) *App {
	return &App{
		WebsocketInit: websocketInit,
		RedisInit:     redisInit,
	}
}

func InitApp(cfg config.Config) (*App, error) {
	wire.Build(
		redis.NewRedisClient,
		WebsocketProviders,
		RedisProviders,
		NewApp,
	)
	return &App{}, nil
}

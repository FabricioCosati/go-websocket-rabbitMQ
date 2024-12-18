// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/config"
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/handlers"
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/services/redis"
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/services/websocket"
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/usecase"
)

// Injectors from wire.go:

func InitApp(cfg config.Config) (*App, error) {
	websocketBrokerServiceImpl := websocket.InitWebsocketBrokerService(cfg)
	websocketClientService := websocket.InitWebsocketClientService()
	hub := websocket.HubInit()
	websocketUsecaseImpl := usecase.InitWebsocketUsecase(websocketBrokerServiceImpl, websocketClientService, hub)
	websocketHandlerImpl := handlers.InitWebsocketHandler(websocketUsecaseImpl)
	websocketInit := NewWebsocketInit(websocketHandlerImpl, hub)
	redisClient := redis.NewRedisClient(cfg)
	redisUsecaseImpl := usecase.InitRedisUsecase(redisClient, hub)
	redisHandlerImpl := handlers.InitRedisHandler(redisUsecaseImpl)
	redisInit := NewRedisInit(redisHandlerImpl)
	app := NewApp(websocketInit, redisInit)
	return app, nil
}

// wire.go:

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

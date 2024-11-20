//go:build wireinject
// +build wireinject

package di

import (
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/config"
	"github.com/google/wire"
)

type App struct {
	WebsocketInit *WebsocketInit
}

func NewApp(websocketInit *WebsocketInit) *App {
	return &App{
		WebsocketInit: websocketInit,
	}
}

func InitApp(cfg config.Config) (*App, error) {
	wire.Build(
		WebsocketProviders,
		NewApp,
	)
	return &App{}, nil
}

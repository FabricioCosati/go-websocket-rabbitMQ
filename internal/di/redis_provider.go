package di

import (
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/handlers"
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/usecase"
	"github.com/google/wire"
)

var RedisUsecase = wire.NewSet(
	usecase.InitRedisUsecase,
	wire.Bind(new(usecase.RedisUsecase), new(*usecase.RedisUsecaseImpl)),
)
var RedisHandler = wire.NewSet(
	handlers.InitRedisHandler,
	wire.Bind(new(handlers.RedisHandler), new(*handlers.RedisHandlerImpl)),
)

var RedisProviders = wire.NewSet(
	RedisUsecase,
	RedisHandler,
	NewRedisInit,
)

type RedisInit struct {
	RedisHandler handlers.RedisHandler
}

func NewRedisInit(
	RedisHandler handlers.RedisHandler,
) *RedisInit {
	return &RedisInit{
		RedisHandler: RedisHandler,
	}
}

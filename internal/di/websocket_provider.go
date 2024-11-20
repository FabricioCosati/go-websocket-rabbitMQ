package di

import (
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/handlers"
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/services/websocket"
	"github.com/FabricioCosati/go-websocket-rabbitMQ/internal/usecase"
	"github.com/google/wire"
)

var WebsocketBrokerService = wire.NewSet(
	websocket.InitWebsocketBrokerService,
	wire.Bind(new(websocket.WebsocketBrokerService), new(*websocket.WebsocketBrokerServiceImpl)),
)
var WebsocketUsecase = wire.NewSet(
	usecase.InitWebsocketUsecase,
	wire.Bind(new(usecase.WebsocketUsecase), new(*usecase.WebsocketUsecaseImpl)),
)
var WebsocketHandler = wire.NewSet(
	handlers.InitWebsocketHandler,
	wire.Bind(new(handlers.WebsocketHandler), new(*handlers.WebsocketHandlerImpl)),
)

var WebsocketProviders = wire.NewSet(
	WebsocketBrokerService,
	websocket.InitWebsocketClientService,
	WebsocketUsecase,
	WebsocketHandler,
	NewWebsocketInit,
)

type WebsocketInit struct {
	WebsocketBrokerService websocket.WebsocketBrokerService
	WebsocketClientService websocket.WebsocketClientService
	WebsocketUsecase       usecase.WebsocketUsecase
	WebsocketHandler       handlers.WebsocketHandler
}

func NewWebsocketInit(
	WebsocketBrokerService websocket.WebsocketBrokerService,
	WebsocketClientService websocket.WebsocketClientService,
	WebsocketUsecase usecase.WebsocketUsecase,
	WebsocketHandler handlers.WebsocketHandler,
) *WebsocketInit {
	return &WebsocketInit{
		WebsocketBrokerService: WebsocketBrokerService,
		WebsocketClientService: WebsocketClientService,
		WebsocketUsecase:       WebsocketUsecase,
		WebsocketHandler:       WebsocketHandler,
	}
}

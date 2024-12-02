package websocket

type HubService interface {
	RegisterClient(*Client)
	UnRegisterClient(*Client)
	HubRun()
}

type Hub struct {
	Clients map[*Client]bool
}

func HubInit() *Hub {
	return &Hub{
		Clients: make(map[*Client]bool),
	}
}

func (h *Hub) RegisterClient(client *Client) {
	h.Clients[client] = true
}

func (h *Hub) UnRegisterClient(client *Client) {
	delete(h.Clients, client)
}

func (h *Hub) HubRun() {

}

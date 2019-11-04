package socket

import (
	"awise-messenger/socket/action"
	"awise-messenger/socket/info"
	"log"
)

// DisseminateToTheOthers for send a message to all account except the broadcaster
type DisseminateToTheOthers struct {
	broadcaster int
	message     []byte
}

// DisseminateToTheTargets for send message to the target
type DisseminateToTheTargets struct {
	Targets []int
	Message []byte
}

// Hub of the clients
type Hub struct {
	clients map[*Client]bool

	Broadcast               chan []byte
	disseminateToTheOthers  chan *DisseminateToTheOthers
	DisseminateToTheTargets chan *DisseminateToTheTargets

	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		Broadcast:               make(chan []byte),
		DisseminateToTheTargets: make(chan *DisseminateToTheTargets),

		disseminateToTheOthers: make(chan *DisseminateToTheOthers),

		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			for other := range h.clients {
				other.send <- action.NewConnection(client.account.ID).Send()
			}
			h.clients[client] = true
			info.Infos.Add(client.account.ID)
			log.Printf("New client register %s (%d) alive now : %d", client.account.Firstname, client.account.ID, info.Infos.NbAlive())
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {

				delete(h.clients, client)
				close(client.send)

				info.Infos.Del(client.account.ID)
				for other := range h.clients {
					other.send <- action.NewDisconnection(client.account.ID).Send()
				}
				log.Printf("Client unregister %s (%d) alive now : %d", client.account.Firstname, client.account.ID, info.Infos.NbAlive())
			}
		case disseminateToTheOthers := <-h.disseminateToTheOthers:
			for client := range h.clients {
				if client.account.ID != disseminateToTheOthers.broadcaster {
					select {
					case client.send <- disseminateToTheOthers.message:
					default:
						h.unregister <- client
					}
				}
			}
		case disseminateToTheTargets := <-h.DisseminateToTheTargets:
			for _, target := range disseminateToTheTargets.Targets {
				if info.Infos.Alive(target) {
					for client := range h.clients {
						if client.account.ID == target {
							select {
							case client.send <- disseminateToTheTargets.Message:
							default:
								h.unregister <- client
							}
						}
					}
				}
			}
		case message := <-h.Broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					h.unregister <- client
				}
			}
		}
	}
}

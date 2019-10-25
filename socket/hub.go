package socket

import (
	"awise-messenger/socket/action"
	"log"
)

// DisseminateToTheOthers for send a message to all account except the broadcaster
type DisseminateToTheOthers struct {
	broadcaster int
	message     []byte
}

// DisseminateToTheTarget for send message to the target
type DisseminateToTheTarget struct {
	target  int
	message []byte
}

// Hub of the clients
type Hub struct {
	clients map[*Client]bool

	broadcast              chan []byte
	disseminateToTheOthers chan *DisseminateToTheOthers
	disseminateToTheTarget chan *DisseminateToTheTarget

	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:              make(chan []byte),
		disseminateToTheOthers: make(chan *DisseminateToTheOthers),
		disseminateToTheTarget: make(chan *DisseminateToTheTarget),
		register:               make(chan *Client),
		unregister:             make(chan *Client),
		clients:                make(map[*Client]bool),
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
			Infos.add(client.account.ID)
			log.Printf("New client register %s (%d) alive now : %d", client.account.Firstname, client.account.ID, Infos.nbAlive())
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {

				delete(h.clients, client)
				close(client.send)

				Infos.del(client.account.ID)
				for other := range h.clients {
					other.send <- action.NewDisconnection(client.account.ID).Send()
				}
				log.Printf("Client unregister %s (%d) alive now : %d", client.account.Firstname, client.account.ID, Infos.nbAlive())
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
		case disseminateToTheTarget := <-h.disseminateToTheTarget:
			for client := range h.clients {
				if client.account.ID == disseminateToTheTarget.target {
					select {
					case client.send <- disseminateToTheTarget.message:
					default:
						h.unregister <- client
					}
				}
			}
		case message := <-h.broadcast:
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

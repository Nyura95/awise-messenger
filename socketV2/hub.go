package socketv2

import "log"

// DisseminateToTheOthers for send a message to all user except the broadcaster
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
			for client := range h.clients {
				client.send <- []byte("New user connected")
			}
			h.clients[client] = true
			Infos.add(client.user.UserID)
			log.Printf("New client register %s (%d) alive now : %d", client.user.Fname, client.user.UserID, Infos.nbAlive())
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				Infos.del(client.user.UserID)
				log.Printf("Client unregister %s (%d) alive now : %d", client.user.Fname, client.user.UserID, Infos.nbAlive())
			}
		case disseminateToTheOthers := <-h.disseminateToTheOthers:
			for client := range h.clients {
				if client.user.UserID != disseminateToTheOthers.broadcaster {
					select {
					case client.send <- disseminateToTheOthers.message:
					default:
						h.unregister <- client
					}
				}
			}
		case disseminateToTheTarget := <-h.disseminateToTheTarget:
			for client := range h.clients {
				if client.user.UserID == disseminateToTheTarget.target {
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

func 
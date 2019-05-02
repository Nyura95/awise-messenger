package socket

import (
	"log"
	"messenger/config"
	"messenger/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

// Transactionnal return
type Transactionnal struct {
	Action  string
	Message string
	ID      int
	Token   string
}

func handler(ws *websocket.Conn) {
	var err error

	customer := Customer{Ws: ws}

	for {
		var transactionnal Transactionnal
		if err = decryptMessage(ws, &transactionnal); err != nil {
			customer.sendMessage(Transactionnal{Action: "Error", Message: "Error parsing"})
			break
		}

		switch transactionnal.Action {
		case "onload":
			log.Println("onload")
			if exists := getCustomerByID(transactionnal.ID); exists != nil {
				deleteCustomer(exists)
			}
			customer.Info.ID = transactionnal.ID
			customer.Info.Token = transactionnal.Token
			Customers = append(Customers, &customer)
			break
		case "onclose":
			log.Println("onclose")
			deleteCustomer(&customer)
			return
		case "onread":
			log.Println("onread")

			// Research conversation by token
			conversation := models.Conversation{Token: customer.Info.Token}
			conversation.FindOneByToken()

			var idTarget int
			if conversation.IDReceiver == customer.Info.ID {
				idTarget = conversation.IDCreator
			} else {
				idTarget = conversation.IDReceiver
			}

			message := models.Message{IDConversation: conversation.IDConversation, IDUser: idTarget}
			message.UpdateMessageRead()

			// Update Status
			if conversation.IDStatus != 2 {
				conversation.IDStatus = 2
				conversation.Update()
			}

			break
		case "send":
			log.Println("send")

			// Research conversation
			conversation := models.Conversation{Token: customer.Info.Token}
			if err = conversation.FindOneByToken(); err != nil {
				customer.sendMessage(Transactionnal{Action: "Error", ID: customer.Info.ID, Message: "Conversation lost", Token: customer.Info.Token})
				break
			}

			// Create message
			message := models.Message{IDUser: customer.Info.ID, IDStatus: 1, IDConversation: conversation.IDConversation, Message: transactionnal.Message}
			message.Create()

			// Update status conversation and IDFirst/IDLast
			conversation.IDStatus = 1
			conversation.IDLastMessage = message.IDMessage
			if conversation.IDFirstMessage == 0 {
				conversation.IDFirstMessage = message.IDMessage
			}
			conversation.Update()

			// Send message to target
			if customer.Info.ID != conversation.IDCreator {
				customer.sendMessageToCustomer(conversation.IDCreator, Transactionnal{Action: "newMessage", Message: transactionnal.Message})
				break
			}
			if customer.Info.ID != conversation.IDReceiver {
				customer.sendMessageToCustomer(conversation.IDReceiver, Transactionnal{Action: "newMessage", Message: transactionnal.Message})
				break
			}
			break
		}
	}
}

// Start the socket server
func Start() {
	config := config.GetConfig()
	Customers = make([]*Customer, 0)
	r := mux.NewRouter()

	r.Handle("/", websocket.Handler(handler))

	log.Println("Start Socket server on localhost:" + strconv.Itoa(config.SocketPort))
	if err := http.ListenAndServe("127.0.0.1:"+strconv.Itoa(config.SocketPort), r); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

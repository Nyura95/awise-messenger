package socket

import (
	"awise-messenger/config"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

// Transactionnal return
type Transactionnal struct {
	Action  string
	Success bool
	Comment string
	Data    string
}

type targetConversation struct {
	ID int
}

func handler(ws *websocket.Conn) {
	var err error

	customer := Customer{Ws: ws}
	log.Println("New")

	for {
		var transactionnal Transactionnal
		if err = decryptMessage(ws, &transactionnal); err != nil {
			log.Println("Error: Parsing error")
			customer.sendMessage(Transactionnal{Action: "Error", Comment: "Error parsing", Success: false, Data: "{}"})
			break
		}

		log.Println(transactionnal)
		if transactionnal.Action != "onload" && customer.Info.Token == "" {
			log.Println("Error: You are not initialize")
			customer.sendMessage(Transactionnal{Action: "Error", Comment: "You are not initialize", Success: false, Data: "{}"})
			break
		}

		switch transactionnal.Action {
		case "onload":
			if err := onLoad(transactionnal, &customer); err != nil {
				log.Printf("Error: %s", err.Error())
				customer.sendMessage(Transactionnal{Action: "Error", Comment: err.Error(), Success: false, Data: "{}"})
			}
			target, err := json.Marshal(targetConversation{ID: customer.Info.ConversationID})
			if err != nil {
				log.Printf("Error: %s", err.Error())
				customer.sendMessage(Transactionnal{Action: "Error", Comment: err.Error(), Success: false, Data: "{}"})
				break
			}
			customer.sendMessage(Transactionnal{Action: "newTargetConversation", Success: true, Data: string(target)})
			break
		case "onclose":
			customer.sendMessage(Transactionnal{Action: "close", Success: true, Data: "{}"})
			if err := onClose(&customer); err != nil {
				log.Printf("Error: %s", err.Error())
				customer.sendMessage(Transactionnal{Action: "Error", Comment: err.Error(), Success: false, Data: "{}"})
			}
			return
		case "onread":
			if err := onRead(&customer); err != nil {
				log.Printf("Error: %s", err.Error())
				customer.sendMessage(Transactionnal{Action: "Error", Comment: err.Error(), Success: false, Data: "{}"})
			}
			break
		case "send":
			newMessage, err := onSend(transactionnal, &customer)
			if err != nil {
				log.Printf("Error: %s", err.Error())
				customer.sendMessage(Transactionnal{Action: "Error", Comment: err.Error(), Success: false, Data: "{}"})
				break
			}
			data, err := json.Marshal(newMessage)
			if err != nil {
				customer.sendMessage(Transactionnal{Action: "Error", Comment: "Error convert message", Success: false, Data: "{}"})
				break
			}

			customer.sendMessageToCustomer(customer.Info.UserID, Transactionnal{Action: "newMessage", Success: true, Data: string(data)})
			log.Printf("check if target %d is online", customer.Info.TargetID)
			if target := getCustomerByID(customer.Info.TargetID); target != nil {
				log.Printf("Target find")
				target.sendMessage(Transactionnal{Action: "newMessage", Success: true, Data: string(data)})
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

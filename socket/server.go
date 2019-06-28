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

func handler(ws *websocket.Conn) {
	var err error

	customer := Customer{Ws: ws}
	log.Println("New")

	for {
		var transactionnal Transactionnal
		if err = decryptMessage(ws, &transactionnal); err != nil {
			customer.sendMessage(Transactionnal{Action: "Error", Comment: "Error parsing", Success: false})
			break
		}

		log.Println(transactionnal)

		switch transactionnal.Action {
		case "onload":
			if err := onLoad(transactionnal, &customer); err != nil {
				log.Printf("Error: %s", err.Error())
				customer.sendMessage(Transactionnal{Action: "Error", Comment: err.Error(), Success: false})
			}
			break
		case "onclose":
			if err := onClose(&customer); err != nil {
				log.Printf("Error: %s", err.Error())
				customer.sendMessage(Transactionnal{Action: "Error", Comment: err.Error(), Success: false})
			}
			return
		case "onread":
			if err := onRead(&customer); err != nil {
				log.Printf("Error: %s", err.Error())
				customer.sendMessage(Transactionnal{Action: "Error", Comment: err.Error(), Success: false})
			}
			break
		case "send":
			newMessage, err := onSend(transactionnal, &customer)
			if err != nil {
				log.Printf("Error: %s", err.Error())
				customer.sendMessage(Transactionnal{Action: "Error", Comment: err.Error(), Success: false})
				break
			}
			data, err := json.Marshal(newMessage)
			if err != nil {
				customer.sendMessage(Transactionnal{Action: "Error", Comment: "Error convert message", Success: false})
				break
			}

			customer.sendMessageToCustomer(customer.Info.UserID, Transactionnal{Action: "newMessage", Success: true, Data: string(data)})
			if target := getCustomerByID(customer.Info.TargetID); target != nil {
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

package socket

import (
	"encoding/json"
	"log"

	"golang.org/x/net/websocket"
)

func decryptMessage(ws *websocket.Conn, transaction *Transactionnal) error {
	var message string
	websocket.Message.Receive(ws, &message)
	log.Printf("Message: %s", message)
	if message == "" {
		transaction.Action = "onclose"
		return nil
	}
	err := json.Unmarshal([]byte(message), &transaction)
	if err != nil {
		return err
	}
	return nil
}

func encryptMessage(transaction Transactionnal) (string, error) {
	message, err := json.Marshal(transaction)
	if err != nil {
		return "", err
	}
	return string(message), nil
}

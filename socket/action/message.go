package action

import (
	"awise-messenger/models"
	"encoding/json"
)

// Message action front
type Message struct {
	Action  string
	IDUser  int
	Message *models.Message
}

// NewMessage create a new instance of message
func NewMessage(IDUser int, IDConversation int, msg string) *Message {
	message, err := models.CreateMessage(IDUser, IDConversation, msg, 1)
	if err != nil {
		return &Message{}
	}
	return &Message{Action: "Message", IDUser: IDUser, Message: message}
}

// Send to the front
func (n *Message) Send() []byte {
	json, _ := json.Marshal(n)
	return append(json)
}

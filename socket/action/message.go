package action

import (
	"awise-messenger/models"
	"encoding/json"
)

// Message action front
type Message struct {
	Action  string          `json:"action"`
	IDUser  int             `json:"idUser"`
	Message *models.Message `json:"message"`
}

// NewMessage create a new instance of message
func NewMessage(IDUser int, IDConversation int, msg string) *Message {
	message, err := models.CreateMessage(IDUser, IDConversation, msg, 1)
	if err != nil {
		return &Message{}
	}
	return &Message{Action: "message", IDUser: IDUser, Message: message}
}

// Send to the front
func (n *Message) Send() []byte {
	json, _ := json.Marshal(n)
	return append(json)
}

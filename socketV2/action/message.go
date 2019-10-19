package action

import (
	"awise-messenger/modelsv2"
	"encoding/json"
)

// Message action front
type Message struct {
	Action  string
	IDUser  int
	Message *modelsv2.Message
}

// NewMessage create a new instance of message
func NewMessage(IDUser int, IDConversation int, msg string) *Message {
	message, err := modelsv2.CreateMessage(IDUser, IDConversation, msg, 1)
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

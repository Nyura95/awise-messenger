package action

import (
	"encoding/json"
)

// Message action front
type Message struct {
	Action    string
	User      int
	Message   string
	IDMessage int
}

// NewMessage create a new instance of message
func NewMessage(user int, message string) *Message {
	return &Message{Action: "Message", User: user, Message: message}
}

// Send to the front
func (n *Message) Send() []byte {
	json, _ := json.Marshal(n)
	return append(json)
}

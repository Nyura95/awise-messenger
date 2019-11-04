package action

import (
	"awise-messenger/models"
	"encoding/json"
)

// Update action front
type Update struct {
	Action  string          `json:"action"`
	Message *models.Message `json:"message"`
}

// NewUpdate create a new instance of message
func NewUpdate(message *models.Message) *Message {
	return &Message{Action: "update", Message: message}
}

// Send to the front
func (m *Update) Send() []byte {
	json, _ := json.Marshal(m)
	return append(json)
}

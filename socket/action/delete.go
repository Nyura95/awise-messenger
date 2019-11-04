package action

import (
	"awise-messenger/models"
	"encoding/json"
)

// Delete action front
type Delete struct {
	Action  string          `json:"action"`
	Message *models.Message `json:"message"`
}

// NewDelete create a new instance of message
func NewDelete(message *models.Message) *Message {
	return &Message{Action: "delete", Message: message}
}

// Send to the front
func (d *Delete) Send() []byte {
	json, _ := json.Marshal(d)
	return append(json)
}

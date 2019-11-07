package action

import (
	"encoding/json"
)

// Private action front
type Private struct {
	Action string `json:"action"`
	Token  string `json:"token"`
}

// NewPrivate create a new instance of message
func NewPrivate(token string) *Private {
	return &Private{Action: "private", Token: token}
}

// Send to the front
func (p *Private) Send() []byte {
	json, _ := json.Marshal(p)
	return append(json)
}

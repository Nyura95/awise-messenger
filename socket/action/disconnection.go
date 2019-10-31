package action

import "encoding/json"

// Disconnection action front
type Disconnection struct {
	Action string `json:"action"`
	User   int    `json:"user"`
}

// NewDisconnection create a new instance of new connection
func NewDisconnection(user int) *Disconnection {
	return &Disconnection{Action: "disconnection", User: user}
}

// Send to the front
func (n *Disconnection) Send() []byte {
	json, _ := json.Marshal(n)
	return append(json)
}

package action

import "encoding/json"

// Disconnection action front
type Disconnection struct {
	Action string
	User   int
}

// NewDisconnection create a new instance of new connection
func NewDisconnection(user int) *Disconnection {
	return &Disconnection{Action: "Disconnection", User: user}
}

// Send to the front
func (n *Disconnection) Send() []byte {
	json, _ := json.Marshal(n)
	return append(json)
}

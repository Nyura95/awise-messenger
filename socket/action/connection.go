package action

import "encoding/json"

// Connection action front
type Connection struct {
	Action string `json:"action"`
	User   int    `json:"user"`
}

// NewConnection create a new instance of new connection
func NewConnection(user int) *Connection {
	return &Connection{Action: "connection", User: user}
}

// Send to the front
func (n *Connection) Send() []byte {
	json, _ := json.Marshal(n)
	return append(json)
}

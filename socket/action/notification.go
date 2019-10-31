package action

import "encoding/json"

// Notification action front
type Notification struct {
	Action  string `json:"action"`
	User    int    `json:"user"`
	Message string `json:"message"`
	Title   string `json:"title"`
}

// NewNotification create a new instance of notification
func NewNotification(user int, message string, title string) *Notification {
	return &Notification{Action: "notification", User: user, Message: message, Title: title}
}

// Send to the front
func (n *Notification) Send() []byte {
	json, _ := json.Marshal(n)
	return append(json)
}

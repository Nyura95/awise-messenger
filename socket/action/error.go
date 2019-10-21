package action

import (
	"encoding/json"
	"strings"
)

// Error action front
type Error struct {
	Action  string
	LocKey  string
	Message string
}

// NewError create a new instance of new error
func NewError(message string) *Error {
	return &Error{Action: "Error", LocKey: strings.ReplaceAll(message, " ", "_"), Message: message}
}

// Send to the front
func (e *Error) Send() []byte {
	json, _ := json.Marshal(e)
	return append(json)
}

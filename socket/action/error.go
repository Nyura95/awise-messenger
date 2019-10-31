package action

import (
	"encoding/json"
	"strings"
)

// Error action front
type Error struct {
	Action  string `json:"action"`
	LocKey  string `json:"lockey"`
	Message string `json:"message"`
}

// NewError create a new instance of new error
func NewError(message string) *Error {
	return &Error{Action: "error", LocKey: strings.ReplaceAll(message, " ", "_"), Message: message}
}

// Send to the front
func (e *Error) Send() []byte {
	json, _ := json.Marshal(e)
	return append(json)
}

package worker

import (
	"awise-messenger/server/response"
)

// PrivateModePayload for call PrivateMode
type PrivateModePayload struct {
	IDUser         int
	IDConversation int
}

// PrivateMode return a basic response
func PrivateMode(payload interface{}) interface{} {
	// context := payload.(GetMessagesPayload)
	return response.BasicResponse(new(interface{}), "ok", 1)
}

package worker

import (
	"awise-messenger/models"
	"awise-messenger/server/response"
	"awise-messenger/socket"
	"awise-messenger/socket/action"
	"log"
)

// SendHashPrivateModePayload for call Login
type SendHashPrivateModePayload struct {
	IDConversation int
	Token          string
}

// SendHashPrivateMode return a basic response
func SendHashPrivateMode(payload interface{}) interface{} {
	context := payload.(SendHashPrivateModePayload)

	_, targets, err := models.FindAllRoomsByIDConversation(context.IDConversation)

	if err != nil {
		log.Println("Error, getting romms")
		return response.BasicResponse(new(interface{}), "Error gettings rooms", -2)
	}

	if len(targets) == 0 {
		log.Println("Error, romms not find")
		return response.BasicResponse(new(interface{}), "Rooms not find", -3)
	}

	socket.ShadowLands.DisseminateToTheTargets <- &socket.DisseminateToTheTargets{Message: action.NewPrivate(context.Token).Send(), Targets: targets}

	return response.BasicResponse("ok", "ok", 1)
}

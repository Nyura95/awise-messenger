package worker

import (
	"awise-messenger/helpers"
	"awise-messenger/models"
	"awise-messenger/server/response"
	"awise-messenger/socket"
	"awise-messenger/socket/action"
	"log"
)

// DeleteMessagePayload for call DeleteMessage
type DeleteMessagePayload struct {
	IDUser         int
	IDConversation int
	IDMessage      int
}

// DeleteMessage return a basic response
func DeleteMessage(payload interface{}) interface{} {
	context := payload.(DeleteMessagePayload)

	_, targets, err := models.FindAllRoomsByIDConversation(context.IDConversation)
	if err != nil {
		log.Println("Error, room not found")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "room not found", -2)
	}

	if exist := helpers.ArrayContainsInt(targets, context.IDUser); exist == false {
		log.Println("Error, user is not on this conversation")
		return response.BasicResponse(new(interface{}), "user is not on this conversation", -3)
	}

	message, err := models.FindMessage(context.IDMessage)
	if err != nil {
		log.Println("Error, message not found")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "message not found", -3)
	}

	if message.IDAccount != context.IDUser {
		log.Println("Error, the user is not the creator")
		return response.BasicResponse(new(interface{}), "the user is not the creator", -4)
	}

	err = message.Delete()
	if err != nil {
		log.Printf("Error, delete message")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "delete message", -5)
	}

	socket.ShadowLands.DisseminateToTheTargets <- &socket.DisseminateToTheTargets{Message: action.NewDelete(message).Send(), Targets: targets}

	return response.BasicResponse(message, "ok", 1)
}

package worker

import (
	"awise-messenger/helpers"
	"awise-messenger/models"
	"awise-messenger/server/response"
	"awise-messenger/socket"
	"awise-messenger/socket/action"
	"log"
)

// UpdateMessagePayload for call UpdateMessage
type UpdateMessagePayload struct {
	IDUser         int
	IDConversation int
	IDMessage      int
	Message        string
}

// UpdateMessage return a basic response
func UpdateMessage(payload interface{}) interface{} {
	context := payload.(UpdateMessagePayload)

	conversation, err := models.FindConversation(context.IDConversation)
	if err != nil || conversation.ID == 0 {
		log.Println("Error, conversion not found")
		return response.BasicResponse(new(interface{}), "conversion not found", -1)
	}

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

	message.Message = context.Message

	err = message.Update()
	if err != nil {
		log.Printf("Error, update message")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "update message", -5)
	}

	socket.ShadowLands.DisseminateToTheTargets <- &socket.DisseminateToTheTargets{Message: action.NewUpdate(message).Send(), Targets: targets}

	return response.BasicResponse(message, "ok", 1)
}

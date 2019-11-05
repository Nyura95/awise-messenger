package worker

import (
	"awise-messenger/enum"
	"awise-messenger/models"
	"awise-messenger/server/response"
	"log"
)

// GetMessagesPayload for call GetMessages
type GetMessagesPayload struct {
	IDUser         int
	IDConversation int
	Page           int
}

// GetMessages return a basic response
func GetMessages(payload interface{}) interface{} {
	context := payload.(GetMessagesPayload)

	room, err := models.FindRoomByIDConversationAndIDAccount(context.IDConversation, context.IDUser)
	if err != nil {
		log.Println("Error, fetch room")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error fetch room", -1)
	}

	if room.ID == 0 {
		log.Println("Error, room not found")
		return response.BasicResponse(new(interface{}), "Error room not found", -2)
	}

	messages, err := models.FindAllMessageByIDConversation(room.IDConversation, enum.NbMessages, context.Page)
	if err != nil {
		log.Println("Error find messages conversation")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error find messages conversation", -4)
	}

	return response.BasicResponse(messages, "ok", 1)
}

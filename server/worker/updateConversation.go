package worker

import (
	"awise-messenger/models"
	"awise-messenger/server/response"
	"log"
)

// UpdateConversationPayload for call UpdateConversation
type UpdateConversationPayload struct {
	IDUser         int
	IDConversation int

	Title string
	Image string
}

// UpdateConversation return a basic response
func UpdateConversation(payload interface{}) interface{} {
	context := payload.(UpdateConversationPayload)

	rooms, _, err := models.FindAllRoomsByIDConversation(context.IDConversation)
	if err != nil {
		log.Println("Error fetch rooms")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error fetch rooms", -1)
	}

	if len(rooms) == 0 {
		log.Println("Error, rooms not find")
		return response.BasicResponse(new(interface{}), "Error rooms not find", -1)
	}

	conversation, err := models.FindConversation(context.IDConversation)
	if err != nil {
		log.Println("Error fetch conversation")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error fetch conversation", -2)
	}

	if conversation.ID == 0 {
		log.Println("Error, conversation not find")
		return response.BasicResponse(new(interface{}), "Error conversation not find", -2)
	}

	if context.Title != "" {
		conversation.Title = context.Title
	}

	if context.Image != "" {
		conversation.Image = context.Image
	}

	if context.Image != "" || context.Title != "" {
		conversation.Update()
	}

	return response.BasicResponse(conversation, "ok", 1)
}

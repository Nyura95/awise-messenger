package worker

import (
	"awise-messenger/models"
	"awise-messenger/server/response"
	"log"
	"sync"
)

// GetConversationPayload for call GetConversation
type GetConversationPayload struct {
	IDUser int
}

// GetConversation return a basic response
func GetConversation(payload interface{}) interface{} {
	context := payload.(GetConversationPayload)

	rooms, err := models.FindAllRoomsByIDAccount(context.IDUser)
	if err != nil {
		log.Println("Error fetch rooms")
		return response.BasicResponse(new(interface{}), "Error fetch rooms", -1)
	}

	var wg sync.WaitGroup
	conversations := []*models.Conversation{}
	wg.Add(len(rooms))
	for _, room := range rooms {
		go func(IDConversation int) {
			defer wg.Done()
			conversation, err := models.FindConversation(IDConversation)
			if err != nil {
				log.Printf("Error, get conversation failed (%d)", IDConversation)
				log.Println(err)
				return
			}
			if conversation.ID != 0 {
				conversations = append(conversations, conversation)
			}
		}(room.IDConversation)
	}

	wg.Wait()

	return response.BasicResponse(conversations, "ok", 1)
}

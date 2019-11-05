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
	conversations := []*models.ConversationWithAllInfos{}
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
				// get all rooms for this conversation
				rooms, _, _ := models.FindAllRoomsByIDConversation(conversation.ID)

				// search for conversation accounts
				var wg2 sync.WaitGroup
				accounts := []*models.Account{}
				token := ""
				wg2.Add(len(rooms))
				for _, room := range rooms {
					if room.IDAccount == context.IDUser {
						token = room.Token
					}
					go func(IDAccount int) {
						defer wg2.Done()
						account, _ := models.FindAccount(IDAccount)
						accounts = append(accounts, account)
					}(room.IDAccount)
				}

				// find all messages for this conversation
				messages, _ := models.FindAllMessageByIDConversation(conversation.ID, 1, 1)

				// wait the accounts search
				wg2.Wait()

				// add the ConversationWithAllInfos
				conversations = append(conversations, &models.ConversationWithAllInfos{Conversation: conversation, Messages: messages, Accounts: accounts, Token: token})
			}
		}(room.IDConversation)
	}

	wg.Wait()

	return response.BasicResponse(conversations, "ok", 1)
}

package worker

import (
	"awise-messenger/enum"
	"awise-messenger/models"
	"awise-messenger/server/response"
	"log"
	"sync"
)

// GetConversationPayload for call GetConversation
type GetConversationPayload struct {
	IDUser         int
	IDConversation int
}

// GetConversation return a basic response
func GetConversation(payload interface{}) interface{} {
	context := payload.(GetConversationPayload)

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

	var wg sync.WaitGroup
	wg.Add(len(rooms))
	accounts := []*models.Account{}
	token := ""
	for _, room := range rooms {
		if room.IDAccount == context.IDUser {
			token = room.Token
			wg.Done()
			continue
		}
		go func(IDAccount int) {
			defer wg.Done()
			account, _ := models.FindAccount(IDAccount)
			accounts = append(accounts, account)
		}(room.IDAccount)
	}

	messages, err := models.FindAllMessageByIDConversation(context.IDConversation, enum.NbMessages, 1)
	if err != nil {
		log.Println("Error fetch messages")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error fetch messages", -3)
	}

	wg.Wait()

	return response.BasicResponse(models.ConversationWithAllInfos{Conversation: conversation, Accounts: accounts, Token: token, Messages: messages}, "ok", 1)
}

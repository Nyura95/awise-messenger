package worker

import (
	"awise-messenger/enum"
	"awise-messenger/helpers"
	"awise-messenger/models"
	"awise-messenger/server/response"
	"log"
	"sync"
)

// CreateConversationInMultiRoomPayload for call CreateConversationInMultiRoom
type CreateConversationInMultiRoomPayload struct {
	IDUser     int
	IDAccounts []int
}

// CreateConversationInMultiRoom return a basic response
func CreateConversationInMultiRoom(payload interface{}) interface{} {
	context := payload.(CreateConversationInMultiRoomPayload)

	context.IDAccounts = append(context.IDAccounts, context.IDUser)

	// check if user exist
	if exist := models.CheckAccountExist(context.IDAccounts...); exist == false {
		log.Println("Error, account of the target not find")
		return response.BasicResponse(new(interface{}), "Target not find", -1)
	}

	conversation, err := models.FindConversationByHash(helpers.Uniqhash(context.IDAccounts...))
	if err != nil {
		log.Println("Error fetch conversation")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error fetch conversation", -2)
	}

	if conversation.ID == 0 {
		conversation, err = models.CreateConversation(helpers.Uniqhash(context.IDAccounts...), "", "", 0, 0, 1, 0)
		if err != nil {
			log.Println("Error create conversation")
			log.Println(err)
			return response.BasicResponse(new(interface{}), "Error create conversation", -3)
		}
		if conversation.ID == 0 {
			log.Println("Error create conversation (0)")
			return response.BasicResponse(new(interface{}), "Error create conversation", -3)
		}

		errors := models.CreateRoomForMultipleAccount(conversation.ID, context.IDAccounts...)
		if len(errors) != 0 {
			log.Println("Error create rooms")
			for err := range errors {
				log.Println(err)
			}
			return response.BasicResponse(new(interface{}), "Error create rooms", -4)
		}
	}

	rooms, _, err := models.FindAllRoomsByIDConversation(conversation.ID)
	if err != nil {
		log.Println("Error find rooms")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error find rooms", -5)
	}

	var wg sync.WaitGroup
	accounts := []*models.Account{}
	token := ""
	wg.Add(len(rooms))
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

	messages, err := models.FindAllMessageByIDConversation(conversation.ID, enum.NbMessages, 1)
	if err != nil {
		log.Println("Error find messages conversation")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error find messages conversation", -6)
	}

	wg.Wait()

	return response.BasicResponse(models.ConversationWithAllInfos{Conversation: conversation, Accounts: accounts, Messages: messages, Token: token}, "ok", 1)
}

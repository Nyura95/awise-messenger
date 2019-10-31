package worker

import (
	"awise-messenger/enum"
	"awise-messenger/helpers"
	"awise-messenger/models"
	"awise-messenger/server/response"
	"log"
)

// GetConversationWithATargetPayload for call GetConversationWithATarget
type GetConversationWithATargetPayload struct {
	IDUser   int
	IDTarget int
}

// GetConversationWithATarget return a basic response
func GetConversationWithATarget(payload interface{}) interface{} {
	context := payload.(GetConversationWithATargetPayload)

	jobs := make(chan *models.Account, 2)
	getAccount := func(ID int, job chan *models.Account) {
		account, _ := models.FindAccount(ID)
		job <- account
	}
	go getAccount(context.IDUser, jobs)
	go getAccount(context.IDTarget, jobs)

	account1 := <-jobs
	account2 := <-jobs
	close(jobs)

	if account1.ID == 0 || account2.ID == 0 {
		log.Println("User or target not find")
		return response.BasicResponse(new(interface{}), "User or target not find", -1)
	}

	conversation, err := models.FindConversationBetweenTwoAccount(account1.ID, account2.ID)
	if err != nil {
		log.Println("Error fetch conversation")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error fetch conversation", -2)
	}

	if conversation.ID == 0 {
		conversation, err := models.CreateConversation(helpers.Uniqhash(account1.ID, account2.ID), "", 0, 0, 1)
		if err != nil {
			log.Println("Error create conversation")
			log.Println(err)
			return response.BasicResponse(new(interface{}), "Error create conversation", -3)
		}
		if conversation.ID == 0 {
			log.Println("Error create conversation (0)")
			return response.BasicResponse(new(interface{}), "Error create conversation", -3)
		}

		err = models.CreateRoomForMultipleAccount(conversation.ID, account1.ID, account2.ID)
		if err != nil {
			log.Println("Error create rooms")
			log.Println(err)
			return response.BasicResponse(new(interface{}), "Error create rooms", -4)
		}
	}

	room, err := models.FindRoomByIDConversationAndIDAccount(conversation.ID, context.IDUser)
	if err != nil {
		log.Println("Error find rooms")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error find rooms", -5)
	}

	messages, err := models.FindAllMessageByIDConversation(conversation.ID, enum.NbMessages, 1)
	if err != nil {
		log.Println("Error find messages conversation")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error find messages conversation", -6)
	}

	var accounts [2]*models.Account
	accounts[0] = account1
	accounts[1] = account2

	return response.BasicResponse(models.ConversationWithToken{Conversation: conversation, Accounts: accounts, Messages: messages, Token: room.Token}, "ok", 1)
}

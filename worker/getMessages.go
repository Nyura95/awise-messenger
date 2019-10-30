package worker

import (
	"awise-messenger/enum"
	"awise-messenger/models"
	"awise-messenger/server/response"
	"log"
)

// GetMessagesPayload for call GetMessages
type GetMessagesPayload struct {
	IDUser   int
	IDTarget int
	Page     int
}

// GetMessages return a basic response
func GetMessages(payload interface{}) interface{} {
	context := payload.(GetMessagesPayload)

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
		log.Printf("User or target does not exist")
		return response.BasicResponse(new(interface{}), "User or target does not exist", -2)
	}

	conversation, err := models.FindConversationBetweenTwoAccount(account1.ID, account2.ID)
	if err != nil {
		log.Printf("Error when getting the room between the accounts")
		return response.BasicResponse(new(interface{}), "Error when getting the room between the accounts", -2)
	}

	if conversation.ID == 0 {
		log.Printf("Conversation does not exist")
		return response.BasicResponse(new(interface{}), "Conversation does not exist", -2)
	}

	messages, err := models.FindAllMessageByIDConversation(conversation.ID, enum.NbMessages, context.Page)
	if err != nil {
		log.Printf("Error when getting the messages")
		return response.BasicResponse(new(interface{}), "Error when getting the messages", -2)
	}

	return response.BasicResponse(messages, "ok", 1)
}

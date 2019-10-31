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
		log.Println("Error create conversation (0)")
		return response.BasicResponse(new(interface{}), "Error create conversation", -3)
	}

	messages, err := models.FindAllMessageByIDConversation(conversation.ID, enum.NbMessages, context.Page)
	if err != nil {
		log.Println("Error find messages conversation")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error find messages conversation", -4)
	}

	return response.BasicResponse(messages, "ok", 1)
}

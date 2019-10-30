package worker

import (
	"awise-messenger/models"
	"awise-messenger/server/response"
	"log"
)

// UpdateMessagePayload for call UpdateMessage
type UpdateMessagePayload struct {
	IDUser    int
	IDTarget  int
	IDMessage int
	Message   string
}

// UpdateMessage return a basic response
func UpdateMessage(payload interface{}) interface{} {
	context := payload.(UpdateMessagePayload)

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

	message, err := models.FindMessage(context.IDMessage)
	if err != nil {
		log.Printf("Error when getting the message")
		return response.BasicResponse(new(interface{}), "Error when getting the message", -2)
	}

	message.Message = context.Message
	err = message.Update()
	if err != nil {
		log.Println(err)
		log.Printf("Error when uddate the message")
		return response.BasicResponse(new(interface{}), "Error when uddate the message", -2)
	}

	return message
}

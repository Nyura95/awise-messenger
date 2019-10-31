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

	message, err := models.FindMessage(context.IDMessage)
	if err != nil {
		log.Println("Error find message")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error find message", -4)
	}

	message.Message = context.Message

	err = message.Update()
	if err != nil {
		log.Printf("Error update message")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error update message", -5)
	}

	return message
}

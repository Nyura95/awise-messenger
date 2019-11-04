package worker

import (
	"awise-messenger/models"
	"awise-messenger/server/response"
	"awise-messenger/socket"
	"awise-messenger/socket/action"
	"log"
)

// DeleteMessagePayload for call UpdateMessage
type DeleteMessagePayload struct {
	IDUser    int
	IDTarget  int
	IDMessage int
	Message   string
}

// DeleteMessage return a basic response
func DeleteMessage(payload interface{}) interface{} {
	context := payload.(DeleteMessagePayload)

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

	if message.IDAccount != context.IDUser {
		log.Println("Error delete message because the user is not the creator")
		return response.BasicResponse(new(interface{}), "Error creator", -5)
	}

	err = message.Delete()
	if err != nil {
		log.Printf("Error update message")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error update message", -6)
	}

	socket.ShadowLands.DisseminateToTheTargets <- &socket.DisseminateToTheTargets{Message: action.NewDelete(message).Send(), Targets: []int{account2.ID, account1.ID}}

	return response.BasicResponse(message, "ok", 1)
}

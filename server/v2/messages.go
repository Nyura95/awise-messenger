package v2

import (
	"awise-messenger/enum"
	"awise-messenger/models"
	"awise-messenger/server/response"
	"awise-messenger/worker"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type getMessagesPayload struct {
	IDUser   int
	IDTarget int
	page     int
}

// GetMessages with the status of the connection socket
func GetMessages(w http.ResponseWriter, r *http.Request) {
	IDUser := context.Get(r, "IDUser").(int)
	IDTarget, err := strconv.Atoi(mux.Vars(r)["IDTarget"])
	page, err2 := strconv.Atoi(mux.Vars(r)["page"])

	if err != nil || err2 != nil {
		log.Printf("The IDTarget is not valid")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "The IDTarget is not valid", -1))
		return
	}

	if IDUser == IDTarget {
		log.Printf("The id's are similar !")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "The id's are similar !", -2))
		return
	}

	pool := worker.CreateWorkerPool(getMessages)
	defer pool.Close()
	json.NewEncoder(w).Encode(pool.Process(getMessagesPayload{IDTarget: IDTarget, IDUser: IDUser, page: page}))
}

func getMessages(payload interface{}) interface{} {
	context := payload.(getMessagesPayload)

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

	messages, err := models.FindAllMessageByIDConversation(conversation.ID, enum.NbMessages, context.page)
	if err != nil {
		log.Printf("Error when getting the messages")
		return response.BasicResponse(new(interface{}), "Error when getting the messages", -2)
	}

	return response.BasicResponse(messages, "ok", 1)
}

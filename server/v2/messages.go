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

type updateMessagePayload struct {
	IDUser    int
	IDTarget  int
	IDMessage int
	Message   string
}

type updateMessagePost struct {
	Message string
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

// UpdateMessage for update a message
func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	IDUser := context.Get(r, "IDUser").(int)
	IDTarget, err := strconv.Atoi(mux.Vars(r)["IDTarget"])
	IDMessage, err2 := strconv.Atoi(mux.Vars(r)["IDMessage"])

	if err != nil || err2 != nil {
		log.Println("The IDTarget is not valid")
		log.Println(err)
		log.Println(err2)
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "The IDTarget is not valid", -1))
		return
	}

	var body updateMessagePost
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&body)

	log.Println(body)

	if IDUser == IDTarget {
		log.Printf("The id's are similar !")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "The id's are similar !", -2))
		return
	}

	pool := worker.CreateWorkerPool(updateMessage)
	defer pool.Close()
	json.NewEncoder(w).Encode(pool.Process(updateMessagePayload{IDTarget: IDTarget, IDUser: IDUser, IDMessage: IDMessage, Message: body.Message}))

}

func updateMessage(payload interface{}) interface{} {
	context := payload.(updateMessagePayload)

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

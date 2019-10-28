package v2

import (
	"awise-messenger/helpers"
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

type conversationWithToken struct {
	*models.Conversation
	Token    string
	Messages []*models.Message
	Accounts [2]*models.Account
}

type getConversationWithATargetPayload struct {
	IDUser   int
	IDTarget int
}

// GetConversationWithATarget get or create a conversation with a other account
func GetConversationWithATarget(w http.ResponseWriter, r *http.Request) {

	IDUser := context.Get(r, "IDUser").(int)
	IDTarget, err := strconv.Atoi(mux.Vars(r)["IDTarget"])
	if err != nil {
		log.Printf("The IDTarget is not valid")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "The IDTarget is not valid", -1))
		return
	}

	if IDUser == IDTarget {
		log.Printf("The id's are similar !")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "The id's are similar !", -2))
		return
	}

	pool := worker.CreateWorkerPool(getConversationWithATarget)
	defer pool.Close()
	json.NewEncoder(w).Encode(pool.Process(getConversationWithATargetPayload{IDUser: IDUser, IDTarget: IDTarget}))
}

func getConversationWithATarget(payload interface{}) interface{} {
	context := payload.(getConversationWithATargetPayload)

	jobs := make(chan *models.Account, 2)
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
		conversation, err := models.CreateConversation(helpers.Uniqhash(account1.ID, account2.ID), "", 0, 0, 1)
		if err != nil || conversation.ID == 0 {
			log.Println(err)
			return response.BasicResponse(new(interface{}), "Error when creating the conversation into the datatable", -2)
		}
		err = models.CreateRoomForMultipleAccount(conversation.ID, account1.ID, account2.ID)
		if err != nil {
			log.Println(err)
			return response.BasicResponse(new(interface{}), "Error when creating the rooms into the datatable", -2)
		}
	}

	if conversation.ID == 0 {
		log.Printf("Error when creating the conversation into the datatable")
		return response.BasicResponse(new(interface{}), "Error when creating the conversation into the datatable", -2)
	}

	room, err := models.FindRoomByIDConversationAndIDAccount(conversation.ID, context.IDUser)
	if err != nil {
		log.Printf("Error when getting the room for the token")
		return response.BasicResponse(new(interface{}), "Error when getting the room for the token", -2)
	}

	messages, err := models.FindAllMessageByIDConversation(conversation.ID, 20)
	if err != nil {
		log.Printf("Error when getting the messages")
		return response.BasicResponse(new(interface{}), "Error when getting the messages", -2)
	}
	reverse(messages)

	var accounts [2]*models.Account
	accounts[0] = account1
	accounts[1] = account2

	return response.BasicResponse(conversationWithToken{Conversation: conversation, Accounts: accounts, Messages: messages, Token: room.Token}, "ok", 1)
}

func getAccount(ID int, job chan *models.Account) {
	account, _ := models.FindAccount(ID)
	job <- account
}

func createNewConversation(account1 *models.Account, account2 *models.Account, create chan *models.Conversation) {
	uniqHash := helpers.Uniqhash(account1.ID, account2.ID)
	conversation, err := models.CreateConversation(uniqHash, "", 0, 0, 1)
	if err != nil {
		log.Println(err)
		create <- conversation
		return
	}
	if conversation.ID == 0 {
		log.Println("Error create new conversation")
		create <- conversation
		return
	}
	models.CreateRoomForMultipleAccount(conversation.ID, account1.ID, account2.ID)
	create <- conversation
}

func reverse(a []*models.Message) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

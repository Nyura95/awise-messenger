package v2

import (
	"awise-messenger/helpers"
	"awise-messenger/modelsv2"
	"awise-messenger/server/response"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

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

	jobs := make(chan *modelsv2.Account, 2)
	go getAccount(IDUser, jobs)
	go getAccount(IDTarget, jobs)

	account1 := <-jobs
	account2 := <-jobs
	close(jobs)

	if account1.ID == 0 || account2.ID == 0 {
		log.Printf("User or target does not exist")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "User or target does not exist", -2))
		return
	}

	IDConversation, err := modelsv2.FindRoomBetweenTwoAccount(account1.ID, account2.ID)
	if err != nil {
		log.Printf("Error when getting the room between the accounts")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "Error when getting the room between the accounts", -2))
		return
	}

	var conversation *modelsv2.Conversation
	if IDConversation == 0 {
		jobs := make(chan *modelsv2.Conversation, 1)
		go createNewConversation(account1, account2, jobs)
		conversation = <-jobs
		if conversation.ID == 0 {
			log.Printf("Error when creating the conversation into the datatable")
			json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "Error when creating the conversation into the datatable", -2))
			return
		}
	} else {
		conversation, err = modelsv2.FindConversation(IDConversation)
		if conversation.ID == 0 {
			log.Printf("Error when creating the conversation into the datatable")
			json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "Error when creating the conversation into the datatable", -2))
			return
		}
	}

	json.NewEncoder(w).Encode(response.BasicResponse(conversation, "ok", 1))
}

func getAccount(ID int, job chan *modelsv2.Account) {
	account, _ := modelsv2.FindAccount(ID)
	job <- account
}

func createNewConversation(account1 *modelsv2.Account, account2 *modelsv2.Account, create chan *modelsv2.Conversation) {
	conversation := &modelsv2.Conversation{IDFirstMessage: 0, IDLastMessage: 0, Title: "", IDStatus: 1}
	conversation.UniqHash = helpers.Uniqhash(account1.ID, account2.ID)
	conversation.TokenConversation = helpers.Token(conversation.UniqHash)
	err := conversation.Create()
	if err != nil {
		create <- conversation
		return
	}
	modelsv2.CreateRoomForMultipleAccount(conversation.ID, account1.ID, account2.ID)
	create <- conversation
}

package v2

import (
	"awise-messenger/helpers"
	"awise-messenger/server/response"
	"awise-messenger/server/worker"

	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type startConversationInMultiRoom struct {
	IDTargets []int
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

	pool := helpers.CreateWorkerPool(worker.GetConversationWithATarget)
	defer pool.Close()

	basicResponse := pool.Process(worker.GetConversationWithATargetPayload{IDUser: IDUser, IDTarget: IDTarget}).(response.Response)
	if basicResponse.Success == false {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(basicResponse)
}

// GetConversations get all conversation for users
func GetConversations(w http.ResponseWriter, r *http.Request) {
	IDUser := context.Get(r, "IDUser").(int)

	pool := helpers.CreateWorkerPool(worker.GetConversations)
	defer pool.Close()

	basicResponse := pool.Process(worker.GetConversationsPayload{IDUser: IDUser}).(response.Response)
	if basicResponse.Success == false {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(basicResponse)
}

// GetConversation get a conversation by id
func GetConversation(w http.ResponseWriter, r *http.Request) {
	IDUser := context.Get(r, "IDUser").(int)
	IDConversation, err := strconv.Atoi(mux.Vars(r)["IDConversation"])
	if err != nil {
		log.Printf("The IDConversation is not valid")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "The IDConversation is not valid", -1))
		return
	}

	pool := helpers.CreateWorkerPool(worker.GetConversation)
	defer pool.Close()

	basicResponse := pool.Process(worker.GetConversationPayload{IDUser: IDUser, IDConversation: IDConversation}).(response.Response)
	if basicResponse.Success == false {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(basicResponse)
}

// StartConversationInMultiRoom create a room with any account
func StartConversationInMultiRoom(w http.ResponseWriter, r *http.Request) {
	IDUser := context.Get(r, "IDUser").(int)

	var body startConversationInMultiRoom
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	if len(body.IDTargets) == 0 {
		log.Printf("Need targets")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "Need targets", -1))
		return
	}

	pool := helpers.CreateWorkerPool(worker.CreateConversationInMultiRoom)
	defer pool.Close()

	basicResponse := pool.Process(worker.CreateConversationInMultiRoomPayload{IDUser: IDUser, IDAccounts: body.IDTargets}).(response.Response)
	if basicResponse.Success == false {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(basicResponse)
}

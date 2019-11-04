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

type updateMessagePost struct {
	Message string
}

// GetMessages with the status of the connection socket
func GetMessages(w http.ResponseWriter, r *http.Request) {
	IDUser := context.Get(r, "IDUser").(int)
	IDConversation, err := strconv.Atoi(mux.Vars(r)["IDConversation"])
	page, err2 := strconv.Atoi(mux.Vars(r)["page"])

	if err != nil || err2 != nil {
		log.Printf("The IDTarget is not valid")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "The IDTarget is not valid", -1))
		return
	}

	pool := helpers.CreateWorkerPool(worker.GetMessages)
	defer pool.Close()
	json.NewEncoder(w).Encode(pool.Process(worker.GetMessagesPayload{IDUser: IDUser, IDConversation: IDConversation, Page: page}))
}

// UpdateMessage for update a message
func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	IDUser := context.Get(r, "IDUser").(int)
	IDConversation, err := strconv.Atoi(mux.Vars(r)["IDConversation"])
	IDMessage, err2 := strconv.Atoi(mux.Vars(r)["IDMessage"])

	if err != nil || err2 != nil {
		log.Println("The query is not valid")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "The query is not valid", -1))
		return
	}

	var body updateMessagePost
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&body)

	pool := helpers.CreateWorkerPool(worker.UpdateMessage)
	defer pool.Close()
	json.NewEncoder(w).Encode(pool.Process(worker.UpdateMessagePayload{IDConversation: IDConversation, IDUser: IDUser, IDMessage: IDMessage, Message: body.Message}))
}

// DeleteMessage for update a message
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	IDUser := context.Get(r, "IDUser").(int)
	IDConversation, err := strconv.Atoi(mux.Vars(r)["IDConversation"])
	IDMessage, err2 := strconv.Atoi(mux.Vars(r)["IDMessage"])

	if err != nil || err2 != nil {
		log.Println("The IDTarget is not valid")
		log.Println(err)
		log.Println(err2)
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "The IDTarget is not valid", -1))
		return
	}

	pool := helpers.CreateWorkerPool(worker.DeleteMessage)
	defer pool.Close()
	json.NewEncoder(w).Encode(pool.Process(worker.DeleteMessagePayload{IDUser: IDUser, IDConversation: IDConversation, IDMessage: IDMessage}))
}

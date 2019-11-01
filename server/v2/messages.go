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

	pool := helpers.CreateWorkerPool(worker.GetMessages)
	defer pool.Close()
	json.NewEncoder(w).Encode(pool.Process(worker.GetMessagesPayload{IDTarget: IDTarget, IDUser: IDUser, Page: page}))
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

	if IDUser == IDTarget {
		log.Printf("The id's are similar !")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "The id's are similar !", -2))
		return
	}

	pool := helpers.CreateWorkerPool(worker.UpdateMessage)
	defer pool.Close()
	json.NewEncoder(w).Encode(pool.Process(worker.UpdateMessagePayload{IDTarget: IDTarget, IDUser: IDUser, IDMessage: IDMessage, Message: body.Message}))
}

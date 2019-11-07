package v2

import (
	"awise-messenger/helpers"
	"awise-messenger/server/response"
	"awise-messenger/server/worker"
	"log"

	"encoding/json"
	"net/http"
)

type getHashPrivateMode struct {
	Strength int
}

type sendHashPrivateMode struct {
	IDConversation int
	Token          string
}

// GetHashPrivateMode return a hash
func GetHashPrivateMode(w http.ResponseWriter, r *http.Request) {

	var body getHashPrivateMode
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	if body.Strength < 1 {
		log.Printf("Strength needed")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "Strength needed", -1))
		return
	}

	if body.Strength > 10 {
		log.Printf("Strength too hight")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "Strength too hight", -2))
		return
	}

	pool := helpers.CreateWorkerPool(worker.GetHashPrivateMode)
	defer pool.Close()

	basicResponse := pool.Process(worker.GetHashPrivateModePayload{Strength: body.Strength}).(response.Response)
	if basicResponse.Success == false {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(basicResponse)
}

// SendHashPrivateMode send a hash to other account
func SendHashPrivateMode(w http.ResponseWriter, r *http.Request) {

	var body sendHashPrivateMode
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	if body.IDConversation == 0 {
		log.Println("IDConversation needed")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "IDConversation needed", -1))
		return
	}

	if body.Token == "" {
		log.Println("Token needed")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "Token needed", -1))
		return
	}

	pool := helpers.CreateWorkerPool(worker.SendHashPrivateMode)
	defer pool.Close()

	basicResponse := pool.Process(worker.SendHashPrivateModePayload{IDConversation: body.IDConversation, Token: body.Token}).(response.Response)
	if basicResponse.Success == false {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(basicResponse)
}

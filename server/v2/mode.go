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

// PrivateMode start a private mode on a conversation with the users connected
func PrivateMode(w http.ResponseWriter, r *http.Request) {
	IDUser := context.Get(r, "IDUser").(int)
	IDConversation, err := strconv.Atoi(mux.Vars(r)["IDConversation"])

	if err != nil {
		log.Printf("The IDConversation is not valid")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "The IDConversation is not valid", -1))
		return
	}

	pool := helpers.CreateWorkerPool(worker.PrivateMode)
	defer pool.Close()

	basicResponse := pool.Process(worker.PrivateModePayload{IDConversation: IDConversation, IDUser: IDUser}).(response.Response)
	if basicResponse.Success == false {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(basicResponse)
}

package v2

import (
	"awise-messenger/server/response"
	"awise-messenger/worker"
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

	pool := worker.CreateWorkerPool(worker.GetConversationWithATarget)
	defer pool.Close()
	json.NewEncoder(w).Encode(pool.Process(worker.GetConversationWithATargetPayload{IDUser: IDUser, IDTarget: IDTarget}))
}

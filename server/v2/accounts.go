package v2

import (
	"awise-messenger/helpers"
	"awise-messenger/server/response"
	"awise-messenger/server/worker"
	"encoding/json"
	"net/http"

	"github.com/gorilla/context"
)

// GetAccounts with the status of the connection socket
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	IDUser := context.Get(r, "IDUser").(int)

	pool := helpers.CreateWorkerPool(worker.GetAccounts)
	defer pool.Close()

	basicResponse := pool.Process(worker.GetAccountsPayload{IDUser: IDUser}).(response.Response)
	if basicResponse.Success == false {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(basicResponse)

}

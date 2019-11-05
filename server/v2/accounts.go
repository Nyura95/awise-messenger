package v2

import (
	"awise-messenger/helpers"
	"awise-messenger/server/response"
	"awise-messenger/server/worker"
	"encoding/json"
	"net/http"
)

// GetAccounts with the status of the connection socket
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	pool := helpers.CreateWorkerPool(worker.GetAccounts)
	defer pool.Close()

	basicResponse := pool.Process(new(interface{})).(response.Response)
	if basicResponse.Success == false {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(basicResponse)

}

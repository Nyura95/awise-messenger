package v2

import (
	"awise-messenger/helpers"
	"awise-messenger/server/worker"
	"encoding/json"
	"net/http"
)

// GetAccounts with the status of the connection socket
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	pool := helpers.CreateWorkerPool(worker.GetAccounts)
	defer pool.Close()
	json.NewEncoder(w).Encode(pool.Process(new(interface{})))
}

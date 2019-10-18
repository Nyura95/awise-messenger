package v2

import (
	"awise-messenger/modelsv2"
	"awise-messenger/server/response"
	"awise-messenger/socketv2"
	"awise-messenger/worker"
	"encoding/json"
	"log"
	"net/http"
)

const (
	databaseError = "database_error"
)

type onlineAccount struct {
	modelsv2.Account
	Online bool
}

// GetAccounts with the status of the connection socket
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := modelsv2.FindAllAccount()
	if err != nil {
		log.Println("Error when getting the users in the database")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), databaseError, -1))
		return
	}
	pool := worker.CreateWorkerPool(getAccountOnline)
	defer pool.Close()
	json.NewEncoder(w).Encode(response.BasicResponse(pool.Process(accounts), "ok", 1))
}

// check if the accounts passed is alive now from the socket and return a interface
func getAccountOnline(payload interface{}) interface{} {
	accounts := payload.([]*modelsv2.Account)

	onlineAccounts := []*onlineAccount{}

	for _, account := range accounts {
		online := false
		for _, id := range socketv2.Infos.List {
			if account.ID == id {
				online = true
			}
		}
		onlineAccounts = append(onlineAccounts, &onlineAccount{Account: *account, Online: online})
	}

	return &onlineAccounts
}

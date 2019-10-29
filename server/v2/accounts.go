package v2

import (
	"awise-messenger/models"
	"awise-messenger/server/response"
	"awise-messenger/socket"
	"awise-messenger/worker"
	"encoding/json"
	"log"
	"net/http"
)

type onlineAccount struct {
	models.Account
	Online bool
}

// GetAccounts with the status of the connection socket
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := models.FindAllAccount()
	if err != nil {
		log.Println("Error when getting the users in the database")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "Error database", -1))
		return
	}
	pool := worker.CreateWorkerPool(getAccountsWorker)
	defer pool.Close()
	json.NewEncoder(w).Encode(response.BasicResponse(pool.Process(accounts), "ok", 1))
}

func getAccountsWorker(payload interface{}) interface{} {
	accounts := payload.([]*models.Account)

	onlineAccounts := []*onlineAccount{}

	for _, account := range accounts {
		online := false
		for _, id := range socket.Infos.List {
			if account.ID == id {
				online = true
			}
		}
		onlineAccounts = append(onlineAccounts, &onlineAccount{Account: *account, Online: online})
	}

	return &onlineAccounts
}

package worker

import (
	"awise-messenger/models"
	"awise-messenger/server/response"
	"awise-messenger/socket"
	"log"
)

// GetAccountsWithInfo is a struct for the front
type GetAccountsWithInfo struct {
	models.Account
	Online bool
}

// GetAccounts for transform accouts in
func GetAccounts(payload interface{}) interface{} {
	accounts, err := models.FindAllAccount()
	if err != nil {
		log.Printf("Error get accounts")
		return response.BasicResponse(new(interface{}), "Error get accounts", -2)
	}
	accountsWithInfos := []*GetAccountsWithInfo{}
	for _, account := range accounts {
		online := false
		for _, id := range socket.Infos.List {
			if account.ID == id {
				online = true
			}
		}
		accountsWithInfos = append(accountsWithInfos, &GetAccountsWithInfo{Account: *account, Online: online})
	}

	return response.BasicResponse(accountsWithInfos, "ok", 1)
}

package worker

import (
	"awise-messenger/models"
	"awise-messenger/server/response"
	"awise-messenger/socket"
	"log"
)

// GetAccounts for transform accouts in
func GetAccounts(payload interface{}) interface{} {
	accounts, err := models.FindAllAccount()
	if err != nil {
		log.Printf("Error get accounts")
		return response.BasicResponse(new(interface{}), "Error get accounts", -2)
	}
	accountInfos := []*models.AccountInfos{}
	for _, account := range accounts {
		online := false
		for _, id := range socket.Infos.List {
			if account.ID == id {
				online = true
			}
		}
		accountInfos = append(accountInfos, &models.AccountInfos{Account: account, Online: online})
	}

	return response.BasicResponse(accountInfos, "ok", 1)
}

package worker

import (
	"awise-messenger/models"
	"awise-messenger/server/response"
	"awise-messenger/socket/info"
	"log"
)

// GetAccounts return a basic response
func GetAccounts(payload interface{}) interface{} {
	accounts, err := models.FindAllAccount()
	if err != nil {
		log.Println("Error fetch accounts")
		return response.BasicResponse(new(interface{}), "Error fetch accounts", -1)
	}
	accountInfos := []*models.AccountInfos{}
	for _, account := range accounts {
		online := false
		for _, id := range info.Infos.List {
			if account.ID == id {
				online = true
			}
		}
		accountInfos = append(accountInfos, &models.AccountInfos{Account: account, Online: online})
	}

	return response.BasicResponse(accountInfos, "ok", 1)
}
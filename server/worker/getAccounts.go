package worker

import (
	"awise-messenger/models"
	"awise-messenger/server/response"
	"awise-messenger/socket/info"
	"log"
)

// GetAccountsPayload for call DeleteMessage
type GetAccountsPayload struct {
	IDUser int
}

// GetAccounts return a basic response
func GetAccounts(payload interface{}) interface{} {
	context := payload.(GetAccountsPayload)

	accounts, err := models.FindAllAccount()
	if err != nil {
		log.Println("Error fetch accounts")
		return response.BasicResponse(new(interface{}), "Error fetch accounts", -1)
	}
	accountInfos := []*models.AccountInfos{}
	for _, account := range accounts {
		if account.ID == context.IDUser {
			continue
		}
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

package worker

import (
	"awise-messenger/models"
	"awise-messenger/server/response"
	"log"
)

// CreateAccountsPayload for call CreateAccounts
type CreateAccountsPayload struct {
	Avatars   string
	Firstname string
	Lastname  string
	Username  string
	Password  string
}

// CreateAccounts return a basic response
func CreateAccounts(payload interface{}) interface{} {
	context := payload.(CreateAccountsPayload)

	account, err := models.CreateAccount(context.Avatars, context.Firstname, context.Lastname, context.Username, context.Password, 2)
	if err != nil {
		log.Println("Error, account created")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error create account", -2)
	}

	return response.BasicResponse(account, "ok", 1)
}

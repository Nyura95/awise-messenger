package worker

import (
	"awise-messenger/models"
	"awise-messenger/server/response"
	"crypto/md5"
	"encoding/hex"
	"log"
)

// LoginPayload for call Login
type LoginPayload struct {
	Username string
	Password string
}

// Login return a basic response
func Login(payload interface{}) interface{} {
	context := payload.(LoginPayload)

	hasher := md5.New()
	hasher.Write([]byte(context.Password))

	account, err := models.FindAccountByPassword(hex.EncodeToString(hasher.Sum(nil)))
	if err != nil {
		log.Println("Error, get account")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error get accounts", -2)
	}

	if account.ID == 0 || account.Username != context.Username {
		log.Println("Error, password or username not valid")
		return response.BasicResponse(new(interface{}), "Error password or username is not valid", -3)
	}

	err = models.DeleteAllAccessTokenByIDAccount(account.ID)
	if err != nil {
		log.Println("Error, disabled all token")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error disabled token", -4)
	}

	accessToken, err := models.CreateAccessToken(account.ID)
	if err != nil {
		log.Println("Error, create token")
		log.Println(err)
		return response.BasicResponse(new(interface{}), "Error create token", -5)
	}

	return response.BasicResponse(accessToken, "ok", 1)
}

package v2

import (
	"awise-messenger/helpers"
	"awise-messenger/server/response"
	"awise-messenger/server/worker"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/context"
)

type createAccounts struct {
	Avatars   string
	Firstname string
	Lastname  string
	Username  string
	Password  string
}

// GetAccounts with the status of the connection socket
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	IDUser := context.Get(r, "IDUser").(int)

	pool := helpers.CreateWorkerPool(worker.GetAccounts)
	defer pool.Close()

	basicResponse := pool.Process(worker.GetAccountsPayload{IDUser: IDUser}).(response.Response)
	if basicResponse.Success == false {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(basicResponse)
}

// CreateAccounts create a new account
func CreateAccounts(w http.ResponseWriter, r *http.Request) {

	var body createAccounts
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	if body.Avatars == "" || body.Firstname == "" || body.Lastname == "" || body.Password == "" || body.Username == "" {
		log.Printf("Body create invalid")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "The body for create a new account is not valid", -1))
		return
	}

	pool := helpers.CreateWorkerPool(worker.CreateAccounts)
	defer pool.Close()

	basicResponse := pool.Process(worker.CreateAccountsPayload{Avatars: body.Avatars, Firstname: body.Firstname, Lastname: body.Lastname, Password: body.Password, Username: body.Username}).(response.Response)
	if basicResponse.Success == false {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(basicResponse)
}

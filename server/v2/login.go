package v2

import (
	"awise-messenger/helpers"
	"awise-messenger/server/response"
	"awise-messenger/server/worker"
	"encoding/json"
	"log"
	"net/http"
)

type login struct {
	Username string
	Password string
}

// Login authenticate an user
func Login(w http.ResponseWriter, r *http.Request) {

	var body login
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	if body.Username == "" || body.Password == "" {
		log.Printf("Body login invalid")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "The body for login is not valid", -1))
		return
	}

	pool := helpers.CreateWorkerPool(worker.Login)
	defer pool.Close()

	basicResponse := pool.Process(worker.LoginPayload{Password: body.Password, Username: body.Username}).(response.Response)
	if basicResponse.Success == false {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(basicResponse)
}

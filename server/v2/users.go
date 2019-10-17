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

type socketUser struct {
	modelsv2.User
	Online bool
}

// GetUsers with the status of the connection socket
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := modelsv2.FindAllUsers()
	if err != nil {
		log.Println("Error when getting the users in the database")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), databaseError, -1))
		return
	}
	pool := worker.CreateWorkerPool(assignSocket)
	defer pool.Close()
	json.NewEncoder(w).Encode(response.BasicResponse(pool.Process(users), "ok", 1))
}

func assignSocket(payload interface{}) interface{} {
	users := payload.([]*modelsv2.User)

	socketUsers := []*socketUser{}

	for _, user := range users {
		online := false
		for _, id := range socketv2.Infos.List {
			if user.UserID == id {
				online = true
			}
		}
		socketUsers = append(socketUsers, &socketUser{User: *user, Online: online})
	}

	return &socketUsers
}

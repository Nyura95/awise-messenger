package v1

import (
	"awise-messenger/models"
	"awise-messenger/server/response"
	"awise-messenger/socket"
	"awise-messenger/worker"
	"encoding/json"
	"log"
	"net/http"
)

type userMessenger struct {
	models.User
	Online bool
}

// GetAllUser for the user
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all users !")

	users, err := models.FindAllUsers()

	if err != nil {
		log.Printf("Error get users !")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "Error get users", -1))
		return
	}

	pool := worker.CreateWorkerPool(makeUserMessenger)
	defer pool.Close()
	json.NewEncoder(w).Encode(response.BasicResponse(pool.Process(users), "ok", 1))
}

func makeUserMessenger(payload interface{}) interface{} {
	users := payload.([]*models.User)
	var usersMessenger []userMessenger

	for _, user := range users {
		online := false
		for _, n := range socket.Customers {
			if user.UserID == n.Info.UserID {
				online = true
			}
		}
		usersMessenger = append(usersMessenger, userMessenger{User: *user, Online: online})
	}

	return usersMessenger
}

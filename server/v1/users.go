package v1

import (
	"awise-messenger/models"
	"awise-messenger/server/response"
	"encoding/json"
	"log"
	"net/http"
)

type userMessenger struct {
	models.User
	Online bool
}

type allUserMessenger []userMessenger

// GetAllUser for the user
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all users !")

	users, err := models.FindAllUsers()

	if err != nil {
		log.Printf("Error get users !")
		log.Panicln(err)
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "Error get users", -1))
		return
	}

	allUserMessenger := make(allUserMessenger, 0)

	for _, user := range users {
		allUserMessenger = append(allUserMessenger, userMessenger{User: *user, Online: true})
	}

	json.NewEncoder(w).Encode(response.BasicResponse(allUserMessenger, "ok", 1))
}

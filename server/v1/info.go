package v1

import (
	"awise-messenger/models"
	"awise-messenger/server/response"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/context"
)

type info struct {
	ConversationNotRead int
	Conversations       int
}

// GetInfo for the user
func GetInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("Get info !")
	// Get ID token
	idUser := context.Get(r, "user_id").(int)

	conversations, _ := models.FindAllConversationByIDUser(idUser)

	var conversationNotRead int
	for _, conversation := range conversations {
		if conversation.IDStatus == 1 {
			conversationNotRead++
		}
	}
	log.Printf("%d conversation find, %d not read", len(conversations), conversationNotRead)

	json.NewEncoder(w).Encode(response.BasicResponse(info{ConversationNotRead: conversationNotRead, Conversations: len(conversations)}, "ok", 1))
}

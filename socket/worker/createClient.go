package worker

import (
	"awise-messenger/models"
	"awise-messenger/socket/info"
)

// CreateClientReturn return CreateClientReturn
type CreateClientReturn struct {
	Account      *models.Account
	Conversation *models.Conversation
	Target       []int
	Auth         bool
	Msg          string
}

// CreateClient return CreateClientReturn
func CreateClient(payload interface{}) interface{} {
	token := payload.(string)
	middleware := &CreateClientReturn{Auth: false}

	if token == "" {
		middleware.Msg = "query is empty"
		return middleware
	}

	room, err := models.FindRoomByToken(token)
	if err != nil {
		middleware.Msg = "Token not find"
		return middleware
	}

	if alive := info.Infos.Alive(room.IDAccount); alive == true {
		// middleware.Msg = "user already connected"
		// return middleware
	}

	account, err := models.FindAccount(room.IDAccount)
	if account.ID == 0 || err != nil {
		middleware.Msg = "user not found"
		return middleware
	}
	middleware.Account = account

	conversation, err := models.FindConversation(room.IDConversation)
	if err != nil {
		middleware.Msg = "Conversation not find"
		return middleware
	}

	middleware.Conversation = conversation

	rooms, err := models.FindAllRoomsByIDConversation(conversation.ID)
	if err != nil {
		middleware.Msg = "Rooms not find"
		return middleware
	}

	var target []int
	for _, room := range rooms {
		if room.IDAccount != account.ID {
			target = append(target, room.IDAccount)
		}
	}

	middleware.Target = target
	middleware.Auth = true

	return middleware
}

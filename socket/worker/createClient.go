package worker

import (
	"awise-messenger/models"
	"awise-messenger/socket/info"
)

const (
	queryEmpty           = "query is empty"
	tokenDelete          = "this token is delete"
	authNotFound         = "auth does not found"
	userAlreadyConnected = "user already connected"
	userNotFound         = "user not found"
	targetNotFound       = "target not found"
	targetIsNotANumber   = "target is not a number"
	tagetIsUser          = "target is the user"
	conversationNotFound = "Conversation not found"
	roundNotFound        = "Round not found"
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
		middleware.Msg = queryEmpty
		return middleware
	}

	room, err := models.FindRoomByToken(token)
	if err != nil {
		middleware.Msg = roundNotFound
		return middleware
	}

	if alive := info.Infos.Alive(room.IDAccount); alive == true {
		middleware.Msg = userAlreadyConnected
		return middleware
	}

	account, err := models.FindAccount(room.IDAccount)
	if account.ID == 0 || err != nil {
		middleware.Msg = userNotFound
		return middleware
	}
	middleware.Account = account

	conversation, err := models.FindConversation(room.IDConversation)
	if err != nil {
		middleware.Msg = conversationNotFound
		return middleware
	}

	middleware.Conversation = conversation

	rooms, err := models.FindAllRoomsByIDConversation(conversation.ID)
	if err != nil {
		middleware.Msg = conversationNotFound // TPM
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

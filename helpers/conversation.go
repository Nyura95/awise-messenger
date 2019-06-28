package helpers

import (
	"awise-messenger/models"
	"crypto/sha1"
	"encoding/base64"
	"strconv"
	"time"
)

// ConversationFront conversation only front
type ConversationFront struct {
	IDConversation int
	Title          string
	IDuser         int
	Token          string
	IDLastMessage  int
	IDFirstMessage int
	IDStatus       int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// TransformConversationInFront conversation to conversationFront
func TransformConversationInFront(conversation *models.Conversation, userID int) ConversationFront {
	if conversation.IDCreator == userID {
		return ConversationFront{IDConversation: conversation.IDConversation, Title: conversation.Title, IDuser: conversation.IDCreator, Token: conversation.TokenCreator, IDStatus: conversation.IDStatus, IDFirstMessage: conversation.IDFirstMessage, IDLastMessage: conversation.IDLastMessage, CreatedAt: conversation.CreatedAt, UpdatedAt: conversation.UpdatedAt}
	}
	return ConversationFront{IDConversation: conversation.IDConversation, Title: conversation.Title, IDuser: conversation.IDReceiver, Token: conversation.TokenReceiver, IDStatus: conversation.IDStatus, IDFirstMessage: conversation.IDFirstMessage, IDLastMessage: conversation.IDLastMessage, CreatedAt: conversation.CreatedAt, UpdatedAt: conversation.UpdatedAt}
}

// Uniqhash for get the uniq_hash conversation
func Uniqhash(creator int, receiver int) string {
	var hash string
	if creator > receiver {
		hash = strconv.Itoa(receiver) + strconv.Itoa(creator)
	} else {
		hash = strconv.Itoa(creator) + strconv.Itoa(receiver)
	}
	return hash
}

// Token for get the token conversation
func Token(Uniqhash string) string {
	h := sha1.New()
	h.Write([]byte(Uniqhash))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

// GetUsersConversation (get the conversation users)
func GetUsersConversation(idCreator int, idReceiver int, users chan []*models.User) {
	listUser := make([]*models.User, 0)
	userCreator := models.User{UserID: idCreator}
	if err := userCreator.FindOne(); err != nil {
		users <- listUser
		return
	}
	userReceiver := models.User{UserID: idReceiver}
	if err := userReceiver.FindOne(); err != nil {
		users <- listUser
		return
	}

	listUser = append(listUser, &userCreator)
	listUser = append(listUser, &userReceiver)
	users <- listUser
}

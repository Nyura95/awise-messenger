package helpers

import (
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

// Uniqhash for get the uniq_hash conversation
func Uniqhash(IDAccount1 int, IDAccount2 int) string {
	var hash string
	if IDAccount1 > IDAccount2 {
		hash = strconv.Itoa(IDAccount2) + strconv.Itoa(IDAccount1)
	} else {
		hash = strconv.Itoa(IDAccount1) + strconv.Itoa(IDAccount2)
	}
	return hash
}

// Token for get the token conversation
func Token(Uniqhash string) string {
	h := sha1.New()
	h.Write([]byte(Uniqhash + "randomkey"))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

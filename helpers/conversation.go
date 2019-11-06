package helpers

import (
	"crypto/sha1"
	"encoding/base64"
	"sort"
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
func Uniqhash(IDAccounts ...int) string {
	hash := ""
	sort.Sort(sort.Reverse(sort.IntSlice(IDAccounts)))
	for _, IDAccount := range IDAccounts {
		hash += strconv.Itoa(IDAccount)
	}
	return hash
}

// Token for get the token conversation
func Token(Uniqhash string) string {
	h := sha1.New()
	h.Write([]byte(Uniqhash + "randomkey"))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

package worker

import (
	"awise-messenger/server/response"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var privateKey = []byte("f98e4479796e4493c8b21ccaf21e7d50a940701d6ff8c582371b9ba95fae3d08")
var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
	"abcdefghijklmnopqrstuvwxyzåäö" +
	"0123456789")

// GetHashPrivateModePayload for call GetHashPrivateMode
type GetHashPrivateModePayload struct {
	Strength int
}

// GetHashPrivateMode return a basic response
func GetHashPrivateMode(payload interface{}) interface{} {
	context := payload.(GetHashPrivateModePayload)

	h := hmac.New(sha256.New, privateKey)

	var wg sync.WaitGroup
	token := ""
	wg.Add(context.Strength)

	for i := 0; i < context.Strength; i++ {
		go func() {
			defer wg.Done()
			generate := generateRune(context.Strength * 10)
			token += generate
		}()
	}

	wg.Wait()

	h.Write([]byte(token))

	return response.BasicResponse(hex.EncodeToString(h.Sum(nil)), "ok", 1)
}

func generateRune(nbChar int) string {
	var b strings.Builder
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nbChar; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

package socketv2

import (
	"awise-messenger/config"
	"awise-messenger/modelsv2"
	"awise-messenger/worker"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
)

type middleware struct {
	account      *modelsv2.Account
	conversation *modelsv2.Conversation
	target       []int

	auth bool
	msg  string
}

// Start the socket server
func Start() {
	config, _ := config.GetConfig()

	hub := newHub()
	go hub.run()

	r := mux.NewRouter()

	r.HandleFunc("/{token}", func(w http.ResponseWriter, r *http.Request) {
		pool := worker.CreateWorkerPool(checkAuth)
		defer pool.Close()
		middleware := pool.Process(r).(*middleware)
		if middleware.auth == false {
			closeServeWs(middleware.msg, w, r)
			return
		}
		serveWs(hub, middleware.account, middleware.conversation, middleware.target, w, r)
	})

	log.Println("Start Socket server on localhost:" + strconv.Itoa(config.SocketPort))
	if err := http.ListenAndServe("127.0.0.1:"+strconv.Itoa(config.SocketPort), r); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func checkAuth(payload interface{}) interface{} {
	r := payload.(*http.Request)
	middleware := &middleware{auth: false}

	token := mux.Vars(r)["token"]

	if token == "" {
		middleware.msg = queryEmpty
		return middleware
	}

	room, err := modelsv2.FindRoomByToken(token)
	if err != nil {
		middleware.msg = authNotFound // TMP
		return middleware
	}

	if alive := Infos.alive(room.IDAccount); alive == true {
		middleware.msg = userAlreadyConnected
		return middleware
	}

	account, err := modelsv2.FindAccount(room.IDAccount)
	if account.ID == 0 || err != nil {
		middleware.msg = userNotFound
		return middleware
	}
	middleware.account = account

	conversation, err := modelsv2.FindConversation(room.IDConversation)
	if err != nil {
		middleware.msg = conversationNotFound
		return middleware
	}

	middleware.conversation = conversation

	rooms, err := modelsv2.FindAllRoomsByIDConversation(conversation.ID)
	if err != nil {
		middleware.msg = conversationNotFound // TPM
		return middleware
	}

	var target []int
	for _, room := range rooms {
		if room.IDAccount != account.ID {
			target = append(target, room.IDAccount)
		}
	}

	middleware.target = target
	middleware.auth = true

	return middleware
}

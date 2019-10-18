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
)

type middleware struct {
	account *modelsv2.Account
	target  int
	auth    bool
	msg     string
}

// Start the socket server
func Start() {
	config, _ := config.GetConfig()

	hub := newHub()
	go hub.run()

	r := mux.NewRouter()

	r.HandleFunc("/{token}/{target}", func(w http.ResponseWriter, r *http.Request) {
		pool := worker.CreateWorkerPool(checkAuth)
		defer pool.Close()
		middleware := pool.Process(r).(*middleware)
		if middleware.auth == false {
			closeServeWs(middleware.msg, w, r)
			return
		}
		serveWs(hub, middleware.account, middleware.target, w, r)
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
	target := mux.Vars(r)["target"]

	if token == "" || target == "" {
		middleware.msg = queryEmpty
		return middleware
	}

	accessToken, err := modelsv2.FindAccessTokenByToken(token)
	if accessToken.ID == 0 || err != nil {
		middleware.msg = authNotFound
		return middleware
	}
	if accessToken.FlagDelete != 0 {
		middleware.msg = tokenDelete
		return middleware
	}

	if alive := Infos.alive(accessToken.IDAccount); alive == true {
		middleware.msg = userAlreadyConnected
		return middleware
	}

	account, err := modelsv2.FindAccount(accessToken.IDAccount)
	if account.ID == 0 || err != nil {
		middleware.msg = userNotFound
		return middleware
	}
	middleware.account = account

	idTarget, err := strconv.Atoi(target)
	if err != nil {
		middleware.msg = targetIsNotANumber
		return middleware
	}
	accountTarget, err := modelsv2.FindAccount(idTarget)
	if accountTarget.ID == 0 || err != nil {
		middleware.msg = targetNotFound
		return middleware
	}
	middleware.target = accountTarget.ID

	if account.ID == accountTarget.ID {
		middleware.msg = tagetIsUser
		return middleware
	}

	middleware.auth = true

	return middleware
}

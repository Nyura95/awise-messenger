package socketv2

import (
	"awise-messenger/config"
	"awise-messenger/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	queryEmpty           = []byte("query is empty")
	tokenDelete          = []byte("this token is delete")
	authNotFound         = []byte("auth does not found")
	userAlreadyConnected = []byte("user already connected")
	userNotFound         = []byte("user not found")
	targetNotFound       = []byte("target not found")
	targetIsNotANumber   = []byte("target is not a number")
	tagetIsUser          = []byte("target is the user")
)

// Start the socket server
func Start() {
	config := config.GetConfig()

	hub := newHub()
	go hub.run()

	r := mux.NewRouter()

	r.HandleFunc("/{token}/{target}", func(w http.ResponseWriter, r *http.Request) {
		auth := mux.Vars(r)["token"]
		target := mux.Vars(r)["target"]

		if auth == "" || target == "" {
			closeServeWs(queryEmpty, w, r)
			return
		}

		accessToken := models.Token{Token: auth}
		if err := accessToken.FindOneByToken(); err != nil {
			closeServeWs(authNotFound, w, r)
			return
		}
		if accessToken.FlagDelete != 0 {
			closeServeWs(tokenDelete, w, r)
			return
		}

		if alive := Infos.alive(accessToken.UserID); alive == true {
			closeServeWs(userAlreadyConnected, w, r)
			return
		}

		user := models.User{UserID: accessToken.UserID}
		if err := user.FindOne(); err != nil {
			closeServeWs(userNotFound, w, r)
			return
		}

		idTarget, err := strconv.Atoi(target)
		if err != nil {
			closeServeWs(targetIsNotANumber, w, r)
			return
		}
		userTarget := models.User{UserID: idTarget}
		if err := userTarget.FindOne(); err != nil {
			closeServeWs(targetNotFound, w, r)
			return
		}

		if user.UserID == userTarget.UserID {
			closeServeWs(tagetIsUser, w, r)
			return
		}
		serveWs(hub, user, userTarget.UserID, w, r)
	})

	log.Println("Start Socket server on localhost:" + strconv.Itoa(config.SocketPort))
	if err := http.ListenAndServe("127.0.0.1:"+strconv.Itoa(config.SocketPort), r); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

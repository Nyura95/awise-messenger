package socket

import (
	"awise-messenger/config"
	"awise-messenger/helpers"
	"awise-messenger/socket/worker"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ShadowLands is a socket hub
var ShadowLands = newHub()

// Start the socket server
func Start() {
	config, _ := config.GetConfig()

	go ShadowLands.run()

	r := mux.NewRouter()

	r.HandleFunc("/{token}", func(w http.ResponseWriter, r *http.Request) {
		pool := helpers.CreateWorkerPool(worker.CreateClient)
		defer pool.Close()
		client := pool.Process(mux.Vars(r)["token"]).(*worker.CreateClientReturn)
		if client.Auth == false {
			closeServeWs(client.Msg, w, r)
			return
		}
		serveWs(ShadowLands, client.Account, client.Conversation, client.Targets, w, r)
	})

	log.Println("Start Socket server on localhost:" + strconv.Itoa(config.SocketPort))
	if err := http.ListenAndServe("127.0.0.1:"+strconv.Itoa(config.SocketPort), r); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

package socketv2

import (
	"awise-messenger/config"
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var addr = flag.String("addr", ":3001", "http service address")

// Start the socket server
func Start() {
	config := config.GetConfig()

	hub := newHub()
	go hub.run()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	log.Println("Start Socket server on localhost:" + strconv.Itoa(config.SocketPort))
	if err := http.ListenAndServe("127.0.0.1:"+strconv.Itoa(config.SocketPort), r); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

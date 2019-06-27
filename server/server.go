package server

import (
	"awise-messenger/config"
	"awise-messenger/server/middleware"
	v1 "awise-messenger/server/v1"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Start for start the http server
func Start() {
	config := config.GetConfig()
	r := mux.NewRouter()

	// Cors auth
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	headersOk := handlers.AllowedHeaders([]string{"authorization", "content-type"})

	// middleware
	r.Use(middleware.BasicHeader)
	r.Use(middleware.Logger)
	r.Use(middleware.IsGoodToken)

	// Get or Create a conversation with one target
	r.HandleFunc("/api/v1/conversation/target/{id}", v1.GetConvoByTarget).Methods("GET")
	// Get or Create a conversation with one target
	r.HandleFunc("/api/v1/conversation/{id}", v1.GetConvoByID).Methods("GET")
	// Get all conversation for the user
	r.HandleFunc("/api/v1/conversation", v1.GetAllConvo).Methods("GET")
	// Get info user
	r.HandleFunc("/api/v1/info", v1.GetInfo).Methods("GET")

	// Ajax
	r.HandleFunc("/", nil).Methods("OPTIONS")

	log.Println("Start http server on localhost:" + strconv.Itoa(config.Port))
	http.ListenAndServe("127.0.0.1:"+strconv.Itoa(config.Port), handlers.CORS(originsOk, headersOk, methodsOk)(r))
}

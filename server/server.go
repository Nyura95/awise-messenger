package server

import (
	"awise-messenger/config"
	"awise-messenger/server/middleware"
	v2 "awise-messenger/server/v2"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Start for start the http server
func Start() {
	config, _ := config.GetConfig()
	r := mux.NewRouter()

	// Cors auth
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	headersOk := handlers.AllowedHeaders([]string{"authorization", "content-type"})

	// middleware
	r.Use(middleware.BasicHeader)
	r.Use(middleware.Logger)
	r.Use(middleware.IsGoodToken)

	r.HandleFunc("/api/v2/accounts", v2.GetAccounts).Methods("GET")

	r.HandleFunc("/api/v2/conversations/target/{IDTarget}", v2.GetConversationWithATarget).Methods("GET")

	// Ajax
	r.HandleFunc("/", nil).Methods("OPTIONS")

	log.Println("Start http server on localhost:" + strconv.Itoa(config.Port))
	srv := &http.Server{
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk)(r),
		Addr:         "127.0.0.1:" + strconv.Itoa(config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}

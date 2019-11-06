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
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	headersOk := handlers.AllowedHeaders([]string{"authorization", "content-type"})

	// middleware
	r.Use(middleware.BasicHeader)
	r.Use(middleware.Logger)
	r.Use(middleware.IsGoodToken)

	// Account
	r.HandleFunc("/api/v2/accounts", v2.GetAccounts).Methods("GET")
	// conversation
	r.HandleFunc("/api/v2/conversations", v2.GetConversations).Methods("GET")
	r.HandleFunc("/api/v2/conversations/{IDConversation}", v2.GetConversation).Methods("GET")
	r.HandleFunc("/api/v2/conversations/{IDConversation}", v2.UpdateConversation).Methods("PUT")
	r.HandleFunc("/api/v2/conversations/target/{IDTarget}", v2.GetConversationWithATarget).Methods("GET")
	// Message
	r.HandleFunc("/api/v2/conversations/{IDConversation}/messages/{page}", v2.GetMessages).Methods("GET")
	r.HandleFunc("/api/v2/conversations/{IDConversation}/messages/{IDMessage}", v2.UpdateMessage).Methods("PUT")
	r.HandleFunc("/api/v2/conversations/{IDConversation}/messages/{IDMessage}", v2.DeleteMessage).Methods("DELETE")
	// Multirooms
	r.HandleFunc("/api/v2/conversations/multi", v2.StartConversationInMultiRoom).Methods("POST")

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

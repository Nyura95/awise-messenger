package middleware

import (
	"awise-messenger/models"
	"log"
	"net/http"

	"github.com/gorilla/context"
)

// BasicHeader for return json
func BasicHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Logger for log new entry
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("new entry : " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// IsGoodToken check if the token Auth is correct
func IsGoodToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			log.Println("auth is empty")
			http.Error(w, "Not authorized", 401)
			return
		}

		accessToken := models.Token{Token: auth}
		if err := accessToken.FindOne(); err != nil {
			http.Error(w, "Not authorized", 401)
			return
		}

		if accessToken.FlagDelete != 0 {
			log.Println("this token is delete")
			http.Error(w, "Not authorized", 401)
			return
		}

		// set context
		context.Set(r, "user_id", accessToken.UserID)

		// next middleware
		next.ServeHTTP(w, r)
	})
}

package static

import (
	"awise-messenger/config"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func staticServer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)

		if strings.Contains(r.URL.Path, ".js") || strings.Contains(r.URL.Path, ".map") || strings.Contains(r.URL.Path, ".css") {
			http.ServeFile(w, r, "./client/build"+r.URL.Path)
			return
		}

		http.ServeFile(w, r, "./client/build")
		return
	})
}

// Start tttt
func Start() {
	config, _ := config.GetConfig()
	r := mux.NewRouter()
	// create a static server
	r.Use(staticServer)
	// force use middleware for not found route
	r.NotFoundHandler = r.NewRoute().HandlerFunc(http.NotFound).GetHandler()

	log.Println("Start http server on localhost:" + strconv.Itoa(config.StaticPort))
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + strconv.Itoa(config.StaticPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

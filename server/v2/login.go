package v2

import (
	"encoding/json"
	"net/http"
)

// Login for log an user
func Login(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("ok")
}

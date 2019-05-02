### Example

```go
// Exemple middleware
func GetBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Body
		username := r.FormValue("username")
    log.Println(username)
    // body json
    decoder := json.NewDecoder(r.Body)
    var a body
    var err error
    decoder.Decode(&a)
    // header
    ua := r.Header.Get("User-Agent")
    log.Println(ua)

    // params (for /{id}/ example)
    vars := mux.Vars(request)
    id := vars["id"]

		next.ServeHTTP(w, r)
	})
}
```

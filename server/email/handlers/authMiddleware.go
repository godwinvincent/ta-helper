package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

//AuthenticatedHandler ..
type AuthenticatedHandler func(http.ResponseWriter, *http.Request, *User)

//EnsureAuth ..
func EnsureAuth(handler AuthenticatedHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xUser := r.Header.Get("X-User")
		if xUser == "" {
			w.Write([]byte("not authorized"))
			return
		}
		var user User
		json.Unmarshal([]byte(xUser), &user)
		log.Println(user)
		handler(w, r, &user)
	})
}

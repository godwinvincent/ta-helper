package handlers

import (
	"encoding/json"
	"net/http"
)

//AuthenticatedHandler ..
type AuthenticatedHandler func(http.ResponseWriter, *http.Request, *User)

type User struct {
	Email     string `json:"email" bson:"email"`
	UserName  string `json:"username" bson:"username"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Role      string `json:"role" bson:"ro"`
}

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
		handler(w, r, &user)
	})
}

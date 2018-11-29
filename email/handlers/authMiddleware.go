package handlers

import (
	"net/http"
)

type User struct {
	ID           int64  `json:"id"`
	Email        string `json:"-"` //never JSON encoded/decoded
	PassHash     []byte `json:"-"` //never JSON encoded/decoded
	UserName     string `json:"userName"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	VerfiedEmail bool   `json:"-"` //never JSON encoded/decoded
}

//AuthenticatedHandler ..
type AuthenticatedHandler func(http.ResponseWriter, *http.Request, *User)

//EnsureAuth ..
func EnsureAuth(handler AuthenticatedHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, &User{
			FirstName: "Godwin",
			LastName:  "Vincent",
			Email:     "godwinvincent@gmail.com",
		})
	})
}

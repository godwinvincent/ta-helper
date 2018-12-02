package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/alabama/final-project-alabama/server/gateway/models/users"
	"github.com/alabama/final-project-alabama/server/gateway/sessions"
)

//UsersHandler handles users
func (ctx *Context) UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decoder := json.NewDecoder(r.Body)
			var newUser users.NewUser
			err := decoder.Decode(&newUser)
			if err != nil {
				http.Error(w, "Request Body not in right format", http.StatusBadRequest)
				return
			}
			if err := newUser.Validate(); err != nil {
				http.Error(w, "Invalid user details", http.StatusBadRequest)
				return
			}
			user, err := newUser.ToUser()
			if err != nil {
				http.Error(w, "Unable to create new user with provided details", http.StatusBadRequest)
				return
			}
			user, err = ctx.UserStore.Insert(user)
			if err != nil {
				http.Error(w, "Unable to store new user in database", http.StatusBadRequest)
				return
			}
			sessions.BeginSession(ctx.SigningKey, ctx.SessionStore, SessionState{time.Now(), *user, false}, w)
			w.Header().Set("Content-Type", "application/json")
			jsonStr, _ := json.Marshal(user)
			w.WriteHeader(http.StatusCreated)
			w.Write(jsonStr)
		} else {
			http.Error(w, "Request Body must be in JSON", http.StatusUnsupportedMediaType)
			return
		}

	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
}

//SessionsHandler handles sessions
func (ctx *Context) SessionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if !strings.Contains(r.Header.Get("Content-type"), "application/json") {
			http.Error(w, "Error request body should be in JSON", http.StatusUnsupportedMediaType)
			return
		}
		var userCredentials users.Credentials
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&userCredentials)
		if err != nil {
			http.Error(w, "Request Body not in right format", http.StatusBadRequest)
			return
		}
		user, err := ctx.UserStore.GetByEmail(userCredentials.Email)
		if err != nil {
			duration := time.Second
			time.Sleep(duration)
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
		err = user.Authenticate(userCredentials.Password)
		if err != nil {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
		_, err = sessions.BeginSession(ctx.SigningKey, ctx.SessionStore, &SessionState{time.Now(), *user, false}, w)
		if err != nil {
			http.Error(w, "Error starting session", http.StatusInternalServerError)
			return
		}
		if err != nil {
			http.Error(w, "Error Logging succesful signin", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		jsonStr, _ := json.Marshal(user)
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonStr)
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
}

//SpecificSessionHandler handles specific sessions
func (ctx *Context) SpecificSessionHandler(w http.ResponseWriter, r *http.Request, currSession *SessionState) {
	if r.Method == "DELETE" {
		pathSlice := strings.Split(r.URL.Path, "/")
		lastSeg := pathSlice[len(pathSlice)-1]
		if lastSeg != "mine" {
			http.Error(w, "Can only end your session", http.StatusForbidden)
			return
		}
		_, err := sessions.EndSession(r, ctx.SigningKey, ctx.SessionStore)
		if err != nil {
			http.Error(w, "Unable to end session", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("signed out"))

	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

}

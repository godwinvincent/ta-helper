package handlers

import (
	"net/http"

	"github.com/alabama/final-project-alabama/server/gateway/models/users"
	"github.com/alabama/final-project-alabama/server/gateway/sessions"
)

type Context struct {
	SigningKey        string
	SessionStore      sessions.Store
	UserStore         users.Store
	NotificationStore *Notifier
}

type AuthenticatedHandler func(http.ResponseWriter, *http.Request, *SessionState)

func (ctx *Context) EnsureAuth(handler AuthenticatedHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currSession := SessionState{}
		if _, err := sessions.GetState(r, ctx.SigningKey, ctx.SessionStore, &currSession); err != nil {
			http.Error(w, "please sign-in", http.StatusUnauthorized)
			return
		}
		handler(w, r, &currSession)
	})
}

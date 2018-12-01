package handlers

import (
	"github.com/alabama/final-project-alabama/server/gateway/models/users"
	"github.com/alabama/final-project-alabama/server/gateway/sessions"
)

type Context struct {
	SigningKey   string
	SessionStore sessions.Store
	UserStore    users.Store
	// NotificationStore *Notifier

}

package handlers

import (
	"github.com/Patrick-Old/Office-Hour-Helper/models/users"
	"github.com/Patrick-Old/Office-Hour-Helper/sessions"
)

type Context struct {
	SigningKey        string
	SessionStore      sessions.Store
	UserStore         users.Store
	NotificationStore *Notifier
}

package handlers

import (
	"time"

	"github.com/Patrick-Old/Office-Hour-Helper/models/users"
)

type SessionState struct {
	SessionStart time.Time  `json:"starttime"`
	User         users.User `json:"userData"`
	TwoFa        bool       `json:"twoFa"`
}

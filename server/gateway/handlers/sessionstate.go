package handlers

import (
	"time"

	"github.com/alabama/final-project-alabama/server/gateway/models/users"
)

type SessionState struct {
	SessionStart time.Time  `json:"starttime"`
	User         users.User `json:"userData"`
	TwoFa        bool       `json:"twoFa"`
}

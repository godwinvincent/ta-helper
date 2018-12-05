package handlers

import "github.com/alabama/final-project-alabama/server/scheduling/models"

type Context struct {
	QuestionCollection   models.QuestionCollection
	OfficeHourCollection models.OfficeHourCollection
	UsersCollection      models.UsersCollection
	WebSocketStore       models.WebsocketStore
}

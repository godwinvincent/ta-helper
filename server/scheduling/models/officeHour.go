package models

import (
	"gopkg.in/mgo.v2/bson"
)

// OfficeHourSession is represents an office hour session
// for one or multiple TAs
type OfficeHourSession struct {
	ID           bson.ObjectId `json:"id" bson:"_id"`
	Name         string        `json:"name" bson:"name"`
	NumQuestions int           `json:"numQuestions" bson:"numQuestions"`
	// slice of TA usernames
	TAs []string `json:"ta" bson:"ta"`
}

// OfficeHourCollection represents our connection to the
// Office hours collection in our Databse
type OfficeHourCollection MongoCollection

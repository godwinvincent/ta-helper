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

// NewOfficeHourSession allows us to take in a new incoming
// Office hour session safely, it allows us to control what fields
// we want to accept.
type NewOfficeHourSession struct {
	Name string `json:"name" bson:"name"`
	// NumQuestions int    `json:"numQuestions" bson:"numQuestions"`
	// slice of TA usernames
	TAs []string `json:"ta" bson:"ta"`
}

type UpdateOfficeHourSession struct {
	Name string `json:"name" bson:"name"`
}

// OfficeHourCollection represents our connection to the
// Office hours collection in our Databse
type OfficeHourCollection MongoCollection

type UsersCollection MongoCollection

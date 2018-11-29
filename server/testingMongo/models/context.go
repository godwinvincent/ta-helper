package models

import (
	"final-project-alabama/server/testingMongo/mongo"
)

// Context holds on to connections and objects held in memory
// that are important to the Gateway.
type Context struct {
	MongoConnection *mongo.MongoSession
}

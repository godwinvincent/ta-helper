package handlers

/**
 * Hi, welcome to Ben's mongo interface code.
 * If you want to work with a mongo DB you need to:
 * 1) make a db connection
 * 2) make a collection struc
 * 3) use that collection struc to call a function.
 */

import mgo "gopkg.in/mgo.v2"

// -------------  Strucs -------------

type MongoSession struct {
	session *mgo.Session
}

type MongoCollection struct {
	collection *mgo.Collection
}

// ------------- Strucs -------------

/**
 * NewSession creates a new connection to the Mongo Database
 * and returns it as a *MongoSession.
 */
func NewSession(url string) (*MongoSession, error) {
	// Use the URL to make a connection to that URL
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	return &MongoSession{session}, err
}

//GetCollection returns a Collection session.
//It takes in name of the DB and name of a collection in that
//Mongo DB.
func (s *MongoSession) GetCollection(dbName string, collectionName string) *MongoCollection {
	tempCollection := MongoCollection{s.session.DB(dbName).C(collectionName)}
	return &tempCollection
}

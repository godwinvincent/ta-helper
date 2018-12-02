package models

import mgo "gopkg.in/mgo.v2"

type MongoSession struct {
	session *mgo.Session
}

type MongoCollection struct {
	Collection *mgo.Collection
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
func (s *MongoSession) GetCollection(dbName string, collectionName string) *mgo.Collection {
	// tempCollection := MongoCollection{s.session.DB(dbName).C(collectionName)}
	return s.session.DB(dbName).C(collectionName)
}

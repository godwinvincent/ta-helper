package mongo

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// -------------  Strucs -------------

type MongoSession struct {
	session *mgo.Session
}

type MongoCollection struct {
	collection *mgo.Collection
}

// ------------- Connection -------------

func NewSession(url string) (*MongoSession, error) {
	// Use the URL to make a connection to that URL
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	return &MongoSession{session}, err
}

func (s *MongoSession) Copy() *MongoSession {
	return &MongoSession{s.session.Copy()}
}

func (s *MongoSession) Close() {
	if s.session != nil {
		s.session.Close()
	}
}

// ------------- Collections -------------

/**
 * Takes in name of DB and name of collection
 */
func (s *MongoSession) GetCollection(dbName string, collectionName string) *MongoCollection {
	tempCollection := MongoCollection{s.session.DB(dbName).C(collectionName)}
	return &tempCollection
}

/**
 *
 */
func (col *MongoCollection) InsertInCollection(user *UserModel) error {
	fmt.Println("Inseeting into collection")
	return col.collection.Insert(user)
}

/**
 * Help: https://hackernoon.com/make-yourself-a-go-web-server-with-mongodb-go-on-go-on-go-on-48f394f24e
 *
 */
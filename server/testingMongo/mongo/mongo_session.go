package mongo

/**
 * Hi, welcome to Ben's mongo interface code.
 * If you want to work with a mongo DB you need to:
 * 1) make a db connection
 * 2) make a collection struc
 * 3) use that collection struc to call a function.
 */

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func (col *MongoCollection) GetByUsername(username string) (*UserModel, error) {
	model := UserModel{}
	err := col.collection.Find(bson.M{"username": username}).One(&model)

	return &model, err
}

/**
 * Help: https://hackernoon.com/make-yourself-a-go-web-server-with-mongodb-go-on-go-on-go-on-48f394f24e
 *
 */

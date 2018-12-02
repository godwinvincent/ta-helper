package questions

/**
 * Hi, welcome to Ben's mongo interface code.
 * If you want to work with a mongo DB you need to:
 * 1) make a db connection
 * 2) make a collection struc
 * 3) use that collection struc to call a function.
 */

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ------------- Code from Gateway -------------
type MongoSession struct {
	session *mgo.Session
}

type MongoCollection struct {
	collection *mgo.Collection
}

//User represents a user account in the database
type User struct {
	Email     string `json:"email" bson:"email"`
	PassHash  []byte `json:"-" bson:"passHash"` //never JSON encoded/decoded
	UserName  string `json:"username" bson:"username"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`

	EmailActivated bool   `json:"-" bson:"emailActivated"` //never JSON encoded/decoded
	EmailVerifCode string `json:"-" bson:"emailVerifCode"`
}

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
	return s.session.DB(dbName).C(collectionName)
}

// GetByUserName retrives a user from the given collection and returns it as a User
func (col *MongoCollection) GetByUserName(username string) (*User, error) {
	model := User{}
	err := col.collection.Find(bson.M{"username": username}).One(&model)
	return &model, err
}

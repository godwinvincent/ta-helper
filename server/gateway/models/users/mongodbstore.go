package users

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

// ------------- Collections -------------

// InsertUser inserts a User into the given Collection
func (col *MongoCollection) Insert(user *User) (*User, error) {

	if err := col.collection.Insert(user); err != nil {
		return nil, err
	}

	newUser, err := col.GetByUserName(user.UserName)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// GetByUserName retrives a user from the given collection and returns it as a User
func (col *MongoCollection) GetByUserName(username string) (*User, error) {
	model := User{}
	err := col.collection.Find(bson.M{"username": username}).One(&model)
	return &model, err
}

func (col *MongoCollection) GetByEmail(email string) (*User, error) {
	model := User{}
	err := col.collection.Find(bson.M{"email": email}).One(&model)
	return &model, err
}

func (col *MongoCollection) GetByID(id string) (*User, error) {
	model := User{}
	err := col.collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&model)
	return &model, err
}

// Delete
func (col *MongoCollection) Delete(id string) error {
	err := col.collection.Remove(bson.M{"_id": id})
	return err
}

// func (col *MongoCollection) Update(id string, newUser *User) error {

// 	err := col.collection.UpdateId(obj1.Id, bson.M{"$set": &obj1})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

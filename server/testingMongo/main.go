package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoSession struct {
	session *mgo.Session
}

type MongoCollection struct {
	Collection *mgo.Collection
}

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

func main() {

	// ------------- Environment -------------

	mongoDBName := "bens_db"

	// ------------- Mongo -------------
	fmt.Println("Beginning...")
	MongoConnection, err := NewSession("localhost:27017")

	if err != nil {
		// fmt.Errorf("Failed to connecto to Mongo DB: %v \n", err)
		log.Fatalf("Failed to connecto to Mongo DB: %v \n", err)

	}
	fmt.Println("Successfully connected to Mongo!")

	// get users collection
	usersCollections := MongoConnection.GetCollection(mongoDBName, "users")

	// ------------- Playing w/ mongo -------------
	// insert a user
	type ish struct {
		Name  string   `json:"name" bson:"name"`
		Thing []string `json:"thing" bson:"thing"`
	}
	// names := make([]string, 3)
	// names[0] = "ddd"
	// names[1] = "tttt"
	// names[2] = "ffff"
	// testObj := ish{"Ben", names}

	// user := mongo.NewUserModel(testObj)

	// if err := usersCollections.Collection.Insert(testObj); err != nil {
	// 	fmt.Println(err)
	// }

	err2 := usersCollections.Collection.Update(bson.M{"_id": bson.ObjectIdHex("5c0465e44282d56ddbe98bc2")}, bson.M{"$addToSet": bson.M{"thing": "NEW1"}})
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("done")
	}

}

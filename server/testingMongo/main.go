package main

import (
	"final-project-alabama/server/testingMongo/models"
	"final-project-alabama/server/testingMongo/mongo"
	"fmt"
	"log"
)

func main() {

	// ------------- Environment -------------

	mongoDBName := "bens_db"

	// ------------- Mongo -------------
	fmt.Println("Beginning...")
	MongoConnection, err := mongo.NewSession("localhost:27017")
	if err != nil {
		// fmt.Errorf("Failed to connecto to Mongo DB: %v \n", err)
		log.Fatalf("Failed to connecto to Mongo DB: %v \n", err)

	}
	fmt.Println("Successfully connected to Mongo!")
	// Context
	ctx := models.Context{MongoConnection}
	// get users collection
	usersCollections := ctx.MongoConnection.GetCollection(mongoDBName, "users")

	// ------------- Playing w/ mongo -------------
	// insert a user
	user := mongo.NewUserModel("testttt", "testtstst@uw.edu")
	err = usersCollections.InsertInCollection(user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User inserted")
	}
	// find a user
	foundUser, foundErr := usersCollections.GetByUsername("bwalchen")
	if foundErr != nil {
		log.Fatalf("failed to get user: %v\n", foundErr)
	}
	fmt.Println("Found user, here:")
	fmt.Println(foundUser)

}

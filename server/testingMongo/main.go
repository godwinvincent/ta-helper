package main

import (
	"final-project-alabama/server/testingMongo/models"
	"final-project-alabama/server/testingMongo/mongo"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Beginning...")

	MongoConnection, err := mongo.NewSession("localhost:27017")
	if err != nil {
		// fmt.Errorf("Failed to connecto to Mongo DB: %v \n", err)
		log.Fatalf("Failed to connecto to Mongo DB: %v \n", err)

	}

	fmt.Println("Successfully connected to Mongo!")

	ctx := models.Context{MongoConnection}
}

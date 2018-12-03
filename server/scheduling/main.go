package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alabama/final-project-alabama/server/scheduling/handlers"
	"github.com/alabama/final-project-alabama/server/scheduling/models"
	"github.com/go-redis/redis"
)

type ServiceEvent struct {
	ServiceName   string    `json:"name"`
	PathPattern   string    `json:"pathPattern"`
	Address       string    `json:"address"`
	LastHeartbeat time.Time `json:"lastHeartbeat"`
	Priviledged   bool      `json:"priviledged"`
}

//main is the main entry point for the server
func main() {
	addr := os.Getenv("ADDR")
	redisAddr := os.Getenv("REDISADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	mongoAddr := os.Getenv("MONGOADDR")
	mongoDBName := os.Getenv("MONGODB")

	redisdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			event := &ServiceEvent{"scheduling", "/v1/officehours", "schedule:80", time.Now(), true}
			jsonString, err := json.Marshal(event)
			if err != nil {
				log.Fatal(err)
			}
			_, err = redisdb.RPush("ServiceEvents", jsonString).Result()
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	fmt.Println("Beginning...")
	MongoConnection, err := models.NewSession(mongoAddr)
	if err != nil {
		log.Fatalf("Failed to connecto to Mongo DB: %v \n", err)
	}
	fmt.Println("Successfully connected to Mongo!")

	// Context
	// ctx := models.Context{MongoConnection}
	// get users collection

	questionCollection := models.QuestionCollection{MongoConnection.GetCollection(mongoDBName, "questions")}
	officeHoursCollection := models.OfficeHourCollection{MongoConnection.GetCollection(mongoDBName, "officeHours")}

	ctx := handlers.Context{
		QuestionCollection:   questionCollection,
		OfficeHourCollection: officeHoursCollection,
	}

	mux := http.NewServeMux()
	mux.Handle("/v1/officehours", handlers.EnsureAuth(ctx.OfficeHourHandler))
	mux.Handle("/v1/officehours/", handlers.EnsureAuth(ctx.SpecificOfficeHourHandler))
	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

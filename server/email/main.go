package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alabama/final-project-alabama/server/email/handlers"
	"github.com/go-redis/redis"
)

type Context struct {
	MongoStore *handlers.MongoCollection
}

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
	if len(addr) == 0 {
		addr = ":80"
	}

	redisAddr := os.Getenv("REDISADDR")
	redisdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			event := &ServiceEvent{"email", "/v1/email", "email:80", time.Now(), true}
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

	mongoDBName := "tahelper"

	fmt.Println("Beginning...")
	MongoConnection, err := handlers.NewSession("localhost:27017")
	if err != nil {
		log.Fatalf("Failed to connecto to Mongo DB: %v \n", err)
	}
	fmt.Println("Successfully connected to Mongo!")

	// Context
	// ctx := models.Context{MongoConnection}
	// get users collection
	usersCollections := MongoConnection.GetCollection(mongoDBName, "users")

	ctx := handlers.Context{
		MongoStore: usersCollections,
	}

	mux := http.NewServeMux()
	mux.Handle("/v1/email", handlers.EnsureAuth(handlers.EmailSendHandler))
	mux.Handle("/v1/verifyEmail", handlers.EnsureAuth(handlers.EmailVerifyHandler))
	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

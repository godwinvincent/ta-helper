package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alabama/final-project-alabama/server/gateway/handlers"
	"github.com/alabama/final-project-alabama/server/gateway/sessions"
	"github.com/go-redis/redis"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//main is the main entry point for the server
func main() {
	addr := os.Getenv("ADDR")
	// tlsKeyPath := os.Getenv("TLSKEY")
	// tlsCertPath := os.Getenv("TLSCERT")
	sessionKey := os.Getenv("SESSIONKEY")
	redisAddr := os.Getenv("REDISADDR")
	dsn := os.Getenv("DSN")

	// if tlsKeyPath == "" || tlsCertPath == "" || sessionKey == "" || redisAddr == "" || dsn == "" {
	// 	fmt.Printf("error reading env variables")
	// 	os.Exit(1)
	// }
	if len(addr) == 0 {
		// addr = ":443"
		addr = ":80"
	}
	redisdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	
	sr := &handlers.ServiceRegistry{
		Registry: make(map[string]*handlers.ServiceInfo),
		Redis:    redisdb,
	}

	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			sr.Mx.Lock()
			sr.Update()
			sr.Mx.Unlock()
		}
	}()



	// ------------- Mongo -------------
	mongoDBName := "bens_db"

	fmt.Println("Beginning...")
	MongoConnection, err := mongo.NewSession("localhost:27017")
	if err != nil {
		log.Fatalf("Failed to connecto to Mongo DB: %v \n", err)
	}
	fmt.Println("Successfully connected to Mongo!")

	ctx := handlers.Context{
		SigningKey:   sessionKey,
		SessionStore: sessions.NewRedisStore(redisdb, time.Hour),
		// UserStore:         users.NewMySQLStore(db),
		// NotificationStore: handlers.NewNotifier(),
		MongoConnection
	}

	// Context
	// ctx := models.Context{MongoConnection}
	// get users collection
	usersCollections := ctx.MongoConnection.GetCollection(mongoDBName, "users")




	mux := http.NewServeMux()
	mux.Handle("/v1/", ctx.ServiceDiscovery(sr))
	wrappedMux := handlers.NewCorsHeader(mux)

	log.Printf("server is listening at %s...", addr)
	// log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, wrappedMux))
	log.Fatal(http.ListenAndServe(addr, wrappedMux))
}
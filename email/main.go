package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

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
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/email", EmailHandler)
	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alabama/final-project-alabama/server/gateway/handlers"
	"github.com/alabama/final-project-alabama/server/gateway/models/users"
	"github.com/alabama/final-project-alabama/server/gateway/sessions"
	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//main is the main entry point for the server
func main() {
	// ------------- Important Variables -------------
	tlsKeyPath := reqEnv("TLSKEY")
	tlsCertPath := reqEnv("TLSCERT")
	sessionKey := reqEnv("SESSIONKEY")

	addr := os.Getenv("ADDR")
	redisAddr := reqEnv("REDISADDR")
	mongoAddr := reqEnv("MONGOADDR")
	mongoDBName := reqEnv("MONGODB")

	if len(addr) == 0 {
		// addr = ":443"
		addr = ":80"
	}

	// ------------- Strucs -------------
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

	webSocketStore := handlers.NewNotifier()

	// ------------- Mongo -------------

	fmt.Println("Mongo testing beginning...")
	MongoConnection, err := users.NewSession(mongoAddr)
	if err != nil {
		log.Fatalf("Failed to connecto to Mongo DB: %v \n", err)
	}
	fmt.Println("Successfully connected to Mongo!")

	// Context
	// ctx := models.Context{MongoConnection}
	// get users collection
	usersCollections := MongoConnection.GetCollection(mongoDBName, "users")

	// ------------- Context -------------
	ctx := handlers.Context{
		SigningKey:        sessionKey,
		SessionStore:      sessions.NewRedisStore(redisdb, time.Hour),
		UserStore:         usersCollections,
		NotificationStore: webSocketStore,
	}

	// ------------- RabbitMQ -------------
	// holds on to queue of messages in order to notify other users

	// conn, mqErr := amqp.Dial("amqp://guest:guest@ localhost:5672/") _______
	conn, mqErr := amqp.Dial("amqp://guest:guest@rabbit:5672/")
	if mqErr != nil {
		log.Fatalf("error: could not connect to RabbitMQ from gateway: %v\n", mqErr)
	}
	defer conn.Close()

	ch, mqCloseerr := conn.Channel()
	if mqCloseerr != nil {
		log.Fatalf("error: could not open channel for RabbitMQ from gateway: %v\n", mqCloseerr)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"MsgQueue", // name matches what we used in our nodejs services
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	// ------------- WebSockets -------------
	type WebsocketMsg struct {
		Usernames []string `json:"usernames"`
		Event     string   `json:"event"`
	}

	go func() {
		for d := range msgs {
			// d is a single notifcation
			tempMsg := WebsocketMsg{}
			err := json.Unmarshal(d.Body, &tempMsg)
			if err != nil {
				log.Println("failed to unmarshal rabbit MQ msg in Gateway main.go")
			} else {
				ctx.NotificationStore.Dispatch(tempMsg.Usernames, []byte(tempMsg.Event))
			}
			d.Ack(false)
		}
	}()

	// ------------- Mux -------------
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/users", ctx.UsersHandler)
	mux.HandleFunc("/v1/sessions", ctx.SessionsHandler)
	mux.Handle("/v1/ws", ctx.EnsureAuth(ctx.WebSocketConnectionHandler))
	mux.Handle("/v1/", ctx.ServiceDiscovery(sr))
	wrappedMux := handlers.NewCorsHeader(mux)

	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, wrappedMux))
	//log.Fatal(http.ListenAndServe(addr, wrappedMux))
}

// ------------- Helper Functions -------------
// Gets information from a environment variable and returns it.
func reqEnv(name string) string {
	val := os.Getenv(name)
	if len(val) == 0 {
		log.Fatalf("please set the %s environment variable", name)
	}
	return val
}

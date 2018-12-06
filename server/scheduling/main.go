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
	"github.com/streadway/amqp"
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
	addr := ":80"
	redisAddr := os.Getenv("REDISADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	mongoAddr := getenv("MONGOADDR")
	mongoDBName := getenv("MONGODB")
	// rabbitAddr := getenv("RABBITADDR")

	redisdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			event := &ServiceEvent{"scheduling", "/v1/(officehours)|(question).*", "schedule:80", time.Now(), true}
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

	MongoConnection, err := models.NewSession(mongoAddr)
	if err != nil {
		log.Fatalf("Failed to connecto to Mongo DB: %v \n", err)
	}
	fmt.Println("Successfully connected to Mongo!")

	// ---------------- RabbitMQ ----------------
	// guide: https://www.rabbitmq.com/tutorials/tutorial-one-go.html
	conn, err := amqp.Dial("amqp://guest:guest@rabbit:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"MsgQueue", // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	w := models.WebsocketStore{
		Channel: ch,
		Queue:   q,
	}

	// ---------------- Context ----------------
	questionCollection := models.QuestionCollection{MongoConnection.GetCollection(mongoDBName, "questions")}
	officeHoursCollection := models.OfficeHourCollection{MongoConnection.GetCollection(mongoDBName, "officeHours")}
	usersCollection := models.UsersCollection{MongoConnection.GetCollection(mongoDBName, "users")}

	ctx := handlers.Context{
		QuestionCollection:   questionCollection,
		OfficeHourCollection: officeHoursCollection,
		UsersCollection:      usersCollection,
		WebSocketStore:       w,
	}

	mux := http.NewServeMux()
	mux.Handle("/v1/officehours", handlers.EnsureAuth(ctx.OfficeHourHandler))
	mux.Handle("/v1/officehours/", handlers.EnsureAuth(ctx.SpecificOfficeHourHandler))
	mux.Handle("/v1/question/", handlers.EnsureAuth(ctx.SpecificQuestionHandler))
	mux.Handle("/v1/ws/", handlers.EnsureAuth(ctx.WebSocketConnectionHandler))
	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

// ---------------- Helper Functions ----------------

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// getenv retrieves the enviornment variable.
// If it fails to find it, main.go will log the issue
// and exit.
func getenv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Fatalf("Error: main.go failed to retrieve env variable: %s\n", key)
	}
	return value
}

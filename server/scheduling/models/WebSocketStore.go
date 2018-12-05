package models

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

// WebsocketStore holds on to the channel and queue
// for producing notifications to RabbitMQ
type WebsocketStore struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
}

// SendNotifToRabbit sends a _____ message or struc of your choice
// to RabbitMQ at the channel in the given WebsocketStore.
// The consumer will need to deserialize the message.
func (W *WebsocketStore) SendNotifToRabbit(jsonMsg string) error {
	// serialize notification
	serializedMsg, errSer := serialize(jsonMsg)
	if errSer != nil {
		return fmt.Errorf("failed to serialize rabbitMQ notifcation: %s", errSer)
	}
	// send a message
	err := W.Channel.Publish(
		"",           // exchange
		W.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			// ContentType: "text/plain",
			Body: serializedMsg,
		})
	if err != nil {
		return fmt.Errorf("error sending notification to rabbit: %s", err)
	}

	return nil
}

// serializes a ____
func serialize(msg string) ([]byte, error) {

	result, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return result, nil
	// var b bytes.Buffer
	// encoder := json.NewEncoder(&b)
	// err := encoder.Encode(msg)
	// return b.Bytes(), err
}

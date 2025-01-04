package internal

import (
	"SimpleGoEventProcessingExample/types"
	"encoding/json"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func NewRabbitConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	HandleError(err, "Error creating RabbitMQ connection")

	return conn
}

func PublishPerson(s *types.Server, id uuid.UUID, channel *amqp.Channel, queue amqp.Queue) {
	body, err := json.Marshal(types.PersonCreatedEvent{ID: id})
	HandleError(err, "Failed to marshal JSON")

	err = channel.PublishWithContext(
		s.Ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	HandleError(err, "Error publishing message")
	log.Printf("[x] Sent %s\n", body)
}

func ProcessEvent(d amqp.Delivery) types.PersonCreatedEvent {
	var event = types.PersonCreatedEvent{}
	err := json.Unmarshal(d.Body, &event)

	HandleError(err, "Failed to unmarshal the message")

	return event
}

package main

import (
	"SimpleGoEventProcessingExample/internal"
	"SimpleGoEventProcessingExample/types"
	"context"
	"encoding/json"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	server := &types.Server{
		Rdb:       internal.NewRedisClient("localhost:6379", "", 0),
		Ctx:       ctx,
		Conn:      internal.NewRabbitConnection(),
		QueueName: "person_created",
	}

	defer server.Conn.Close()
	defer server.Rdb.Close()

	internal.PingOrPanic(server)

	ch, err := server.Conn.Channel()
	internal.HandleError(err, "Error creating RabbitMQ channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		server.QueueName,
		false,
		false,
		false,
		false,
		nil)

	internal.HandleError(err, "Failed to declare a queue")

	msgs, err := ch.ConsumeWithContext(
		ctx,
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil)

	internal.HandleError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			event := internal.ProcessEvent(d)
			person := internal.GetPerson(server, event.ID)

			body, err := json.Marshal(person)
			internal.HandleError(err, "Failed to marshal JSON")

			log.Printf("Person: %s\n", body)

			internal.RemovePerson(server, event.ID)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

package main

import (
	"SimpleGoEventProcessingExample/internal"
	"SimpleGoEventProcessingExample/types"
	"context"
	"fmt"
	"math/rand"
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

	for i := 0; i <= 10; i++ {
		person := types.Person{
			ID:   internal.NewUUID(),
			Name: fmt.Sprintf("Person %s", internal.NewUUID()),
			Age:  rand.Intn(100),
		}

		internal.SavePerson(server, &person)
		internal.PublishPerson(server, person.ID, ch, queue)

		time.Sleep(1 * time.Second)
	}
}

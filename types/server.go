package types

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	Rdb       *redis.Client
	Ctx       context.Context
	Conn      *amqp.Connection
	QueueName string
}

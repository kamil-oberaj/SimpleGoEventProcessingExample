package internal

import (
	"SimpleGoEventProcessingExample/types"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"log"
)

func PingOrPanic(s *types.Server) {
	if s.Rdb == nil {
		panic("Redis client is nil")
	}

	pong, err := s.Rdb.Ping(s.Ctx).Result()
	HandleError(err, "Failed to ping Redis")

	fmt.Printf("Redis client is connected: %s\n", pong)
}

func NewRedisClient(url, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       db,
	})
}

func SavePerson(s *types.Server, person *types.Person) {
	if s.Rdb == nil {
		log.Panicf("Redis connection is nil")
	}

	if s.Ctx == nil {
		log.Panicf("Context is nil")
	}

	key := fmt.Sprintf("person:%s", person.ID)
	value, err := json.Marshal(person)
	HandleError(err, "Failed to marshal JSON")

	err = s.Rdb.Set(s.Ctx, key, value, 0).Err()
	HandleError(err, "Failed to save person in Redis")

	log.Printf("[x] Saved in Redis: %s\n", key)
}

func GetPerson(s *types.Server, id uuid.UUID) *types.Person {
	if s.Rdb == nil {
		log.Panicf("Redis connection is nil")
	}

	if s.Ctx == nil {
		log.Panicf("Context is nil")
	}

	key := createKey(id)
	value, err := s.Rdb.Get(s.Ctx, key).Result()
	HandleError(err, "Failed to get person from Redis")

	var person types.Person
	err = json.Unmarshal([]byte(value), &person)
	HandleError(err, "Failed to unmarshal JSON")

	log.Printf("[x] Got from Redis: %s\n", key)
	return &person
}

func RemovePerson(s *types.Server, id uuid.UUID) {
	if s.Rdb == nil {
		log.Panicf("Redis connection is nil")
	}

	if s.Ctx == nil {
		log.Panicf("Context is nil")
	}

	key := createKey(id)
	err := s.Rdb.Del(s.Ctx, key).Err()
	HandleError(err, "Failed to remove person from Redis")

	log.Printf("[x] Removed from Redis: %s\n", key)
}

func createKey(id uuid.UUID) string {
	return fmt.Sprintf("person:%s", id)
}

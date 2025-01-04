package internal

import (
	"github.com/google/uuid"
	"log"
)

func HandleError(err error, msg string) {
	if err != nil {
		log.Panicf("%s : %s", msg, err)
	}
}

func NewUUID() uuid.UUID {
	id, err := uuid.NewV7()
	HandleError(err, "Error generating UUID")
	return id
}

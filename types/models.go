package types

import (
	"github.com/google/uuid"
)

type Person struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

type PersonCreatedEvent struct {
	ID uuid.UUID `json:"id"`
}

/*
func (person *Person) String() string {
	return fmt.Sprintf("ID: %s, Name: %s, Age: %d", person.ID, person.Name, person.Age)
}
*/

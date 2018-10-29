package entity

import (
	"github.com/satori/go.uuid"
)

// ID is an entity ID
type ID uuid.UUID

func newID() (ID, error) {
	id, err := uuid.NewV4()
	return ID(id), err
}

// IsZero indicates if the ID is the zero-value
func (id ID) IsZero() bool {
	return id == ID{}
}

// String represents the ID as a string
func (id ID) String() string {
	return uuid.UUID(id).String()
}

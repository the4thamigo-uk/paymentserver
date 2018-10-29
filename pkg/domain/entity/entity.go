package entity

import (
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

// Entity is base type that manages a domain entity's identiy and version
type Entity struct {
	ID      string `json:"entity_id"`
	Version int    `json:"version"`
}

var zero = Entity{}

// New creates a new entity identifier
func New() (Entity, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return zero, err
	}
	return Entity{
		ID:      id.String(),
		Version: 0,
	}, nil
}

// MustNew creates a new entity identifier and panics on error.
// Use for testing only
func MustNew() Entity {
	id, err := New()
	if err != nil {
		panic(err)
	}
	return id
}

// Validate checks the consistency of an entity
func (e *Entity) Validate() error {
	_, err := uuid.FromString(e.ID)
	if err != nil {
		return errors.Wrapf(err, "Entity id is empty")
	}
	if e.Version < 0 {
		return errors.New("Entity version is not valid")
	}
	return nil
}

// NextVersion increments the version number for the entity
func (e *Entity) NextVersion() {
	e.Version++
}

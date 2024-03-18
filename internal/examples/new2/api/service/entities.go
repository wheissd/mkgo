package service

import (
	"github.com/google/uuid"
	"github.com/wheissd/mkgo/internal/examples/new2/ent/gen"
)

type Opt[T any] struct {
	isSet bool
	val   T
}

func (o *Opt[T]) Set(v T) {
	o.val = v
	o.isSet = true
}

func (o *Opt[T]) IsSet() bool {
	return o.isSet
}

func (o *Opt[T]) Get() T {
	return o.val
}

type Human struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func NewHuman(e *gen.Human) Human {
	return Human{
		ID:   e.ID,
		Name: e.Name,
	}
}

// GetID returns the value of ID.
func (s Human) GetID() uuid.UUID {
	return s.ID
}

// GetName returns the value of ID.
func (s Human) GetName() string {
	return s.Name
}

type CreateHuman struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UpdateHuman struct {
	ID   Opt[uuid.UUID] `json:"id"`
	Name Opt[string]    `json:"name"`
}

type HumanList struct {
	Items []Human
}

type HumanListParams struct {
	Page         Opt[int]
	ItemsPerPage Opt[int]
}

func NewHumanList(es []*gen.Human) []Human {
	if len(es) == 0 {
		return nil
	}
	r := make([]Human, len(es))
	for i, e := range es {
		r[i] = NewHuman(e)
	}
	return r
}

type HumanReadParams struct {
	ID uuid.UUID
}

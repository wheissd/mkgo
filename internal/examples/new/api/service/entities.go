package service

import (
	"github.com/google/uuid"
	"github.com/wheissd/mkgo/internal/examples/new/ent/gen"
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

type DefaultModel struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func NewDefaultModel(e *gen.DefaultModel) DefaultModel {
	return DefaultModel{
		ID:   e.ID,
		Name: e.Name,
	}
}

// GetID returns the value of ID.
func (s DefaultModel) GetID() uuid.UUID {
	return s.ID
}

// GetName returns the value of ID.
func (s DefaultModel) GetName() string {
	return s.Name
}

type CreateDefaultModel struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UpdateDefaultModel struct {
	ID   Opt[uuid.UUID] `json:"id"`
	Name Opt[string]    `json:"name"`
}

type DefaultModelList struct {
	Items []DefaultModel
}

type DefaultModelListParams struct {
	Page         Opt[int]
	ItemsPerPage Opt[int]
}

func NewDefaultModelList(es []*gen.DefaultModel) []DefaultModel {
	if len(es) == 0 {
		return nil
	}
	r := make([]DefaultModel, len(es))
	for i, e := range es {
		r[i] = NewDefaultModel(e)
	}
	return r
}

type DefaultModelReadParams struct {
	ID uuid.UUID
}

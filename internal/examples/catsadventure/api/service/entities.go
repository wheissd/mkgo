package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/internal/examples/catsadventure/ent/gen"
)

type Opt[T any] struct {
	isSet bool
	val   T
}

func NewOptP[T any](v *T) Opt[T] {
	if v != nil {
		return Opt[T]{
			val: *v,
		}
	}
	return Opt[T]{}
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

type Breed struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	DeletedTime Opt[time.Time] `json:"deleted_time"`
	CreateTime  time.Time      `json:"create_time"`
	UpdateTime  time.Time      `json:"update_time"`
	Edges       BreedEdges     `json:"edges"`
}

type BreedEdges struct {
	Cats []Cat
}

func NewBreed(e *gen.Breed) Breed {
	res := Breed{
		ID:          e.ID,
		Name:        e.Name,
		DeletedTime: NewOptP(e.DeletedTime),
		CreateTime:  e.CreateTime,
		UpdateTime:  e.UpdateTime,
	}
	if e.Edges.Cats != nil {
		res.Edges.Cats = make([]Cat, 0, len(e.Edges.Cats))
		for i := range e.Edges.Cats {
			res.Edges.Cats = append(res.Edges.Cats, NewCat(e.Edges.Cats[i]))
		}
	}
	return res
}

type CreateBreed struct {
	Name string `json:"name"`
}

type UpdateBreed struct {
	Name Opt[string] `json:"name"`
}

type BreedList struct {
	Items []Breed
}

type BreedQueryEdges struct {
	Cats *CatQueryEdges
}

type BreedListParams struct {
	Page          Opt[int]
	ItemsPerPage  Opt[int]
	FilterCatsIDs Opt[[]uuid.UUID]
	With          *BreedQueryEdges
}

func NewBreedList(es []*gen.Breed) []Breed {
	if len(es) == 0 {
		return nil
	}
	r := make([]Breed, len(es))
	for i, e := range es {
		r[i] = NewBreed(e)
	}
	return r
}

type BreedReadParams struct {
	ID   uuid.UUID
	With *BreedQueryEdges
}

type Cat struct {
	Name         string         `json:"name"`
	Speed        int64          `json:"speed"`
	Type         string         `json:"type"`
	DeletedTime  Opt[time.Time] `json:"deleted_time"`
	CreateTime   time.Time      `json:"create_time"`
	UpdateTime   time.Time      `json:"update_time"`
	ID           uuid.UUID      `json:"id"`
	BreedID      uuid.UUID      `json:"breed_id"`
	DateFrom     time.Time      `json:"date_from"`
	OtherType    string         `json:"other_type"`
	PhantomField Opt[string]    `json:"phantom_field"`
	Edges        CatEdges       `json:"edges"`
}

type CatEdges struct {
	Kittens []Kitten
	Breed   Opt[Breed]
}

func NewCat(e *gen.Cat) Cat {
	res := Cat{
		Name:        e.Name,
		Speed:       e.Speed,
		Type:        string(e.Type),
		DeletedTime: NewOptP(e.DeletedTime),
		CreateTime:  e.CreateTime,
		UpdateTime:  e.UpdateTime,
		ID:          e.ID,
		BreedID:     e.BreedID,
		DateFrom:    e.DateFrom,
		OtherType:   string(e.OtherType),
	}
	if e.Edges.Kittens != nil {
		res.Edges.Kittens = make([]Kitten, 0, len(e.Edges.Kittens))
		for i := range e.Edges.Kittens {
			res.Edges.Kittens = append(res.Edges.Kittens, NewKitten(e.Edges.Kittens[i]))
		}
	}
	if e.Edges.Breed != nil {
		res.Edges.Breed.Set(NewBreed(e.Edges.Breed))
	}
	return res
}

type CreateCat struct {
	Name      string    `json:"name"`
	Speed     int64     `json:"speed"`
	Type      string    `json:"type"`
	BreedID   uuid.UUID `json:"breed_id"`
	DateFrom  time.Time `json:"date_from"`
	OtherType string    `json:"other_type"`
}

type UpdateCat struct {
	Name      Opt[string]    `json:"name"`
	Speed     Opt[int64]     `json:"speed"`
	Type      Opt[string]    `json:"type"`
	BreedID   Opt[uuid.UUID] `json:"breed_id"`
	DateFrom  Opt[time.Time] `json:"date_from"`
	OtherType Opt[string]    `json:"other_type"`
}

type CatList struct {
	Items []Cat
}

type CatQueryEdges struct {
	Kittens *KittenQueryEdges
	Breed   *BreedQueryEdges
}

type CatListParams struct {
	Page             Opt[int]
	ItemsPerPage     Opt[int]
	FilterKittensIDs Opt[[]uuid.UUID]
	FilterBreedID    Opt[uuid.UUID]
	With             *CatQueryEdges
}

func NewCatList(es []*gen.Cat) []Cat {
	if len(es) == 0 {
		return nil
	}
	r := make([]Cat, len(es))
	for i, e := range es {
		r[i] = NewCat(e)
	}
	return r
}

type CatReadParams struct {
	ID   uuid.UUID
	With *CatQueryEdges
}

type FatherCat struct {
	Name        string         `json:"name"`
	DeletedTime Opt[time.Time] `json:"deleted_time"`
	CreateTime  time.Time      `json:"create_time"`
	UpdateTime  time.Time      `json:"update_time"`
	ID          uuid.UUID      `json:"id"`
	Edges       FatherCatEdges `json:"edges"`
}

type FatherCatEdges struct {
}

func NewFatherCat(e *gen.FatherCat) FatherCat {
	res := FatherCat{
		Name:        e.Name,
		DeletedTime: NewOptP(e.DeletedTime),
		CreateTime:  e.CreateTime,
		UpdateTime:  e.UpdateTime,
		ID:          e.ID,
	}
	return res
}

type CreateFatherCat struct {
	Name string `json:"name"`
}

type UpdateFatherCat struct {
	Name Opt[string] `json:"name"`
}

type FatherCatList struct {
	Items []FatherCat
}

type FatherCatQueryEdges struct {
}

type FatherCatListParams struct {
	Page         Opt[int]
	ItemsPerPage Opt[int]
	With         *FatherCatQueryEdges
}

func NewFatherCatList(es []*gen.FatherCat) []FatherCat {
	if len(es) == 0 {
		return nil
	}
	r := make([]FatherCat, len(es))
	for i, e := range es {
		r[i] = NewFatherCat(e)
	}
	return r
}

type FatherCatReadParams struct {
	ID   uuid.UUID
	With *FatherCatQueryEdges
}

type Kitten struct {
	Name        string         `json:"name"`
	MotherID    uuid.UUID      `json:"mother_id"`
	DeletedTime Opt[time.Time] `json:"deleted_time"`
	CreateTime  time.Time      `json:"create_time"`
	UpdateTime  time.Time      `json:"update_time"`
	ID          uuid.UUID      `json:"id"`
	Edges       KittenEdges    `json:"edges"`
}

type KittenEdges struct {
	Mother Opt[Cat]
}

func NewKitten(e *gen.Kitten) Kitten {
	res := Kitten{
		Name:        e.Name,
		MotherID:    e.MotherID,
		DeletedTime: NewOptP(e.DeletedTime),
		CreateTime:  e.CreateTime,
		UpdateTime:  e.UpdateTime,
		ID:          e.ID,
	}
	if e.Edges.Mother != nil {
		res.Edges.Mother.Set(NewCat(e.Edges.Mother))
	}
	return res
}

type CreateKitten struct {
	Name     string    `json:"name"`
	MotherID uuid.UUID `json:"mother_id"`
}

type UpdateKitten struct {
	Name     Opt[string]    `json:"name"`
	MotherID Opt[uuid.UUID] `json:"mother_id"`
}

type KittenList struct {
	Items []Kitten
}

type KittenQueryEdges struct {
	Mother *CatQueryEdges
}

type KittenListParams struct {
	Page           Opt[int]
	ItemsPerPage   Opt[int]
	FilterMotherID Opt[uuid.UUID]
	With           *KittenQueryEdges
}

func NewKittenList(es []*gen.Kitten) []Kitten {
	if len(es) == 0 {
		return nil
	}
	r := make([]Kitten, len(es))
	for i, e := range es {
		r[i] = NewKitten(e)
	}
	return r
}

type KittenReadParams struct {
	ID   uuid.UUID
	With *KittenQueryEdges
}

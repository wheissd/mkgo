package ogen

import (
	"github.com/wheissd/mkgo/internal/examples/catsadventure/api/service"
)

func NewBreed(e service.Breed) Breed {
	var ret Breed
	ret.ID = e.ID
	ret.Name = e.Name
	ret.DeletedTime.SetTo(e.DeletedTime.Get())
	ret.CreateTime = e.CreateTime
	ret.UpdateTime = e.UpdateTime
	if e.Edges.Cats != nil {
		ret.Edges.Cats = make([]Cat, 0, len(e.Edges.Cats))
		for i := range e.Edges.Cats {
			ret.Edges.Cats = append(ret.Edges.Cats, NewCat(e.Edges.Cats[i]))
		}
	}

	return ret
}
func NewBreedList(es []service.Breed) []Breed {
	if len(es) == 0 {
		return nil
	}
	r := make([]Breed, len(es))
	for i, e := range es {
		r[i] = NewBreed(e)
	}
	return r
}

func NewCat(e service.Cat) Cat {
	var ret Cat
	ret.Name = e.Name
	ret.Speed = e.Speed
	ret.Type = e.Type
	ret.DeletedTime.SetTo(e.DeletedTime.Get())
	ret.CreateTime = e.CreateTime
	ret.UpdateTime = e.UpdateTime
	ret.ID = e.ID
	ret.BreedID = e.BreedID
	ret.DateFrom = e.DateFrom
	ret.OtherType = e.OtherType
	ret.PhantomField.SetTo(e.PhantomField.Get())
	if e.Edges.Kittens != nil {
		ret.Edges.Kittens = make([]Kitten, 0, len(e.Edges.Kittens))
		for i := range e.Edges.Kittens {
			ret.Edges.Kittens = append(ret.Edges.Kittens, NewKitten(e.Edges.Kittens[i]))
		}
	}
	if e.Edges.Breed.IsSet() {
		ret.Edges.Breed.SetTo(NewBreed(e.Edges.Breed.Get()))
	}

	return ret
}
func NewCatList(es []service.Cat) []Cat {
	if len(es) == 0 {
		return nil
	}
	r := make([]Cat, len(es))
	for i, e := range es {
		r[i] = NewCat(e)
	}
	return r
}

func NewFatherCat(e service.FatherCat) FatherCat {
	var ret FatherCat
	ret.Name = e.Name
	ret.DeletedTime.SetTo(e.DeletedTime.Get())
	ret.CreateTime = e.CreateTime
	ret.UpdateTime = e.UpdateTime
	ret.ID = e.ID

	return ret
}
func NewFatherCatList(es []service.FatherCat) []FatherCat {
	if len(es) == 0 {
		return nil
	}
	r := make([]FatherCat, len(es))
	for i, e := range es {
		r[i] = NewFatherCat(e)
	}
	return r
}

func NewKitten(e service.Kitten) Kitten {
	var ret Kitten
	ret.Name = e.Name
	ret.MotherID = e.MotherID
	ret.DeletedTime.SetTo(e.DeletedTime.Get())
	ret.CreateTime = e.CreateTime
	ret.UpdateTime = e.UpdateTime
	ret.ID = e.ID
	if e.Edges.Mother.IsSet() {
		ret.Edges.Mother.SetTo(NewCat(e.Edges.Mother.Get()))
	}

	return ret
}
func NewKittenList(es []service.Kitten) []Kitten {
	if len(es) == 0 {
		return nil
	}
	r := make([]Kitten, len(es))
	for i, e := range es {
		r[i] = NewKitten(e)
	}
	return r
}

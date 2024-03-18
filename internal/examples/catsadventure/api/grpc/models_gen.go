package grpc

import (
	"github.com/wheissd/mkgo/internal/examples/catsadventure/api/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewBreed(e service.Breed) *Breed {
	var ret Breed
	ret.Name = e.Name
	ret.ID = e.ID.String()
	if e.Edges.Cats != nil {
		ret.Edges.Cats = make([]*Cat, 0, len(e.Edges.Cats))
		for i := range e.Edges.Cats {
			ret.Edges.Cats = append(ret.Edges.Cats, NewCat(e.Edges.Cats[i]))
		}
	}

	return &ret
}
func NewBreedList(es []service.Breed) []*Breed {
	if len(es) == 0 {
		return nil
	}
	r := make([]*Breed, len(es))
	for i, e := range es {
		item := NewBreed(e)
		r[i] = item
	}
	return r
}
func catOtherTypeProtoToServiceConverter(v int32) string {
	switch v {
	case 0:
		return "merch"
	case 1:
		return "hotel"
	case 2:
		return "tournament"
	}
	return ""
}
func catTypeProtoToServiceConverter(v int32) string {
	switch v {
	case 0:
		return "merch"
	case 1:
		return "hotel"
	case 2:
		return "tournament"
	}
	return ""
}

func NewCat(e service.Cat) *Cat {
	var ret Cat
	ret.Name = e.Name
	ret.BreedID = e.BreedID.String()
	ret.Speed = e.Speed
	ret.DateFrom = timestamppb.New(e.DateFrom)
	ret.OtherType = CatOtherTypeEnum(CatOtherTypeEnum_value[e.OtherType])
	ret.Type = CatTypeEnum(CatTypeEnum_value[e.Type])
	ret.ID = e.ID.String()
	ret.PhantomField = e.PhantomField.Get()
	if e.Edges.Kittens != nil {
		ret.Edges.Kittens = make([]*Kitten, 0, len(e.Edges.Kittens))
		for i := range e.Edges.Kittens {
			ret.Edges.Kittens = append(ret.Edges.Kittens, NewKitten(e.Edges.Kittens[i]))
		}
	}
	if e.Edges.Breed.IsSet() {
		ret.Edges.Breed = NewBreed(e.Edges.Breed.Get())
	}

	return &ret
}
func NewCatList(es []service.Cat) []*Cat {
	if len(es) == 0 {
		return nil
	}
	r := make([]*Cat, len(es))
	for i, e := range es {
		item := NewCat(e)
		r[i] = item
	}
	return r
}

func NewFatherCat(e service.FatherCat) *FatherCat {
	var ret FatherCat
	ret.Name = e.Name
	ret.ID = e.ID.String()

	return &ret
}
func NewFatherCatList(es []service.FatherCat) []*FatherCat {
	if len(es) == 0 {
		return nil
	}
	r := make([]*FatherCat, len(es))
	for i, e := range es {
		item := NewFatherCat(e)
		r[i] = item
	}
	return r
}

func NewKitten(e service.Kitten) *Kitten {
	var ret Kitten
	ret.Name = e.Name
	ret.MotherID = e.MotherID.String()
	ret.ID = e.ID.String()
	if e.Edges.Mother.IsSet() {
		ret.Edges.Mother = NewCat(e.Edges.Mother.Get())
	}

	return &ret
}
func NewKittenList(es []service.Kitten) []*Kitten {
	if len(es) == 0 {
		return nil
	}
	r := make([]*Kitten, len(es))
	for i, e := range es {
		item := NewKitten(e)
		r[i] = item
	}
	return r
}

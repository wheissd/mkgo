package ogen

import (
	"errors"
	"strings"

	"github.com/wheissd/mkgo/internal/examples/catsadventure/api/service"
)

func parseWithBreed(level int, query string, e *service.BreedQueryEdges) (*service.BreedQueryEdges, error) {
	if level > 5 {
		return nil, errors.New("max nesting reached")
	}
	level++
	currentPos, queryRemainder, found := strings.Cut(query, ",")
	if e == nil {
		e = &service.BreedQueryEdges{}
	}
	var err error
	switch currentPos {
	case "cats":
		e.Cats, err = parseWithCat(level, queryRemainder, e.Cats)
		if err != nil {
			return nil, err
		}
	}

	if !found && query == "" {
		return e, nil
	}
	return nil, errors.New("invalid with param")
}
func parseWithCat(level int, query string, e *service.CatQueryEdges) (*service.CatQueryEdges, error) {
	if level > 5 {
		return nil, errors.New("max nesting reached")
	}
	level++
	currentPos, queryRemainder, found := strings.Cut(query, ",")
	if e == nil {
		e = &service.CatQueryEdges{}
	}
	var err error
	switch currentPos {
	case "kittens":
		e.Kittens, err = parseWithKitten(level, queryRemainder, e.Kittens)
		if err != nil {
			return nil, err
		}
	case "breed":
		e.Breed, err = parseWithBreed(level, queryRemainder, e.Breed)
		if err != nil {
			return nil, err
		}
	}

	if !found && query == "" {
		return e, nil
	}
	return nil, errors.New("invalid with param")
}
func parseWithFatherCat(level int, query string, e *service.FatherCatQueryEdges) (*service.FatherCatQueryEdges, error) {
	if level > 5 {
		return nil, errors.New("max nesting reached")
	}
	level++
	if e == nil {
		e = &service.FatherCatQueryEdges{}
	}
	return e, nil
}
func parseWithKitten(level int, query string, e *service.KittenQueryEdges) (*service.KittenQueryEdges, error) {
	if level > 5 {
		return nil, errors.New("max nesting reached")
	}
	level++
	currentPos, queryRemainder, found := strings.Cut(query, ",")
	if e == nil {
		e = &service.KittenQueryEdges{}
	}
	var err error
	switch currentPos {
	case "mother":
		e.Mother, err = parseWithCat(level, queryRemainder, e.Mother)
		if err != nil {
			return nil, err
		}
	}

	if !found && query == "" {
		return e, nil
	}
	return nil, errors.New("invalid with param")
}

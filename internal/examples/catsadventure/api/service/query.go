package service

import (
	"github.com/wheissd/mkgo/internal/examples/catsadventure/internal/examples/catsadventure/ent/gen"
)

func applyWithBreed(edges *BreedQueryEdges) func(*gen.BreedQuery) {
	return func(q *gen.BreedQuery) {
		if edges.Cats != nil {
			q.WithCats()
		}
	}
}

func applyWithCat(edges *CatQueryEdges) func(*gen.CatQuery) {
	return func(q *gen.CatQuery) {
		if edges.Kittens != nil {
			q.WithKittens()
		}
		if edges.Breed != nil {
			q.WithBreed()
		}
	}
}

func applyWithFatherCat(edges *FatherCatQueryEdges) func(*gen.FatherCatQuery) {
	return func(q *gen.FatherCatQuery) {
	}
}

func applyWithKitten(edges *KittenQueryEdges) func(*gen.KittenQuery) {
	return func(q *gen.KittenQuery) {
		if edges.Mother != nil {
			q.WithMother()
		}
	}
}

package service

import (
	"context"

	"github.com/wheissd/mkgo/internal/examples/catsadventure/internal/examples/catsadventure/ent/gen"
)

type BreedQueryModifier func(ctx context.Context, q *gen.BreedQuery) error

func noOpBreedQueryModifier(ctx context.Context, q *gen.BreedQuery) error { return nil }

type CatQueryModifier func(ctx context.Context, q *gen.CatQuery) error

func noOpCatQueryModifier(ctx context.Context, q *gen.CatQuery) error { return nil }

type FatherCatQueryModifier func(ctx context.Context, q *gen.FatherCatQuery) error

func noOpFatherCatQueryModifier(ctx context.Context, q *gen.FatherCatQuery) error { return nil }

type KittenQueryModifier func(ctx context.Context, q *gen.KittenQuery) error

func noOpKittenQueryModifier(ctx context.Context, q *gen.KittenQuery) error { return nil }

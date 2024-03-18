package service

import (
	"context"

	"github.com/wheissd/mkgo/internal/examples/new2/ent/gen"
)

type HumanQueryModifier func(ctx context.Context, q *gen.HumanQuery) error

func noOpHumanQueryModifier(ctx context.Context, q *gen.HumanQuery) error { return nil }

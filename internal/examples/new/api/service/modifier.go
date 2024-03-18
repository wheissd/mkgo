package service

import (
	"context"

	"github.com/wheissd/mkgo/internal/examples/new/ent/gen"
)

type DefaultModelQueryModifier func(ctx context.Context, q *gen.DefaultModelQuery) error

func noOpDefaultModelQueryModifier(ctx context.Context, q *gen.DefaultModelQuery) error { return nil }

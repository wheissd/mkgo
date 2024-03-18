package db

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wheissd/mkgo/internal/examples/cats/config"
	"github.com/wheissd/mkgo/internal/examples/cats/ent/gen"
	"go.uber.org/fx"

	cache "ariga.io/entcache"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	_ "github.com/wheissd/mkgo/internal/examples/cats/ent/gen/runtime"
)

type DB struct {
	Default *pgxpool.Pool
}

func New(cfg config.DB) *DB {
	pgx, err := pgxpool.New(context.TODO(), cfg.Default.URL)
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		Default: pgx,
	}
}

// NewEntClient инициализируем entClient
func NewEntClient(cfg config.DB) *gen.Client {
	driver, err := entsql.Open(dialect.Postgres, cfg.Default.URL)
	if err != nil {
		log.Fatal(err)
	}

	cacheDriver := cache.NewDriver(
		driver, cache.ContextLevel(),
	)
	return gen.NewClient(
		gen.Driver(
			cacheDriver,
		),
	)
}

var Module = fx.Module(
	"db",
	fx.Provide(
		NewEntClient,
		New,
	))

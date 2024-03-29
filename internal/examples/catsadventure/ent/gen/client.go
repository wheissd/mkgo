// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/breed"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/cat"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/fathercat"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/kitten"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Breed is the client for interacting with the Breed builders.
	Breed *BreedClient
	// Cat is the client for interacting with the Cat builders.
	Cat *CatClient
	// FatherCat is the client for interacting with the FatherCat builders.
	FatherCat *FatherCatClient
	// Kitten is the client for interacting with the Kitten builders.
	Kitten *KittenClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Breed = NewBreedClient(c.config)
	c.Cat = NewCatClient(c.config)
	c.FatherCat = NewFatherCatClient(c.config)
	c.Kitten = NewKittenClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("gen: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("gen: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Breed:     NewBreedClient(cfg),
		Cat:       NewCatClient(cfg),
		FatherCat: NewFatherCatClient(cfg),
		Kitten:    NewKittenClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Breed:     NewBreedClient(cfg),
		Cat:       NewCatClient(cfg),
		FatherCat: NewFatherCatClient(cfg),
		Kitten:    NewKittenClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Breed.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Breed.Use(hooks...)
	c.Cat.Use(hooks...)
	c.FatherCat.Use(hooks...)
	c.Kitten.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Breed.Intercept(interceptors...)
	c.Cat.Intercept(interceptors...)
	c.FatherCat.Intercept(interceptors...)
	c.Kitten.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *BreedMutation:
		return c.Breed.mutate(ctx, m)
	case *CatMutation:
		return c.Cat.mutate(ctx, m)
	case *FatherCatMutation:
		return c.FatherCat.mutate(ctx, m)
	case *KittenMutation:
		return c.Kitten.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("gen: unknown mutation type %T", m)
	}
}

// BreedClient is a client for the Breed schema.
type BreedClient struct {
	config
}

// NewBreedClient returns a client for the Breed from the given config.
func NewBreedClient(c config) *BreedClient {
	return &BreedClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `breed.Hooks(f(g(h())))`.
func (c *BreedClient) Use(hooks ...Hook) {
	c.hooks.Breed = append(c.hooks.Breed, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `breed.Intercept(f(g(h())))`.
func (c *BreedClient) Intercept(interceptors ...Interceptor) {
	c.inters.Breed = append(c.inters.Breed, interceptors...)
}

// Create returns a builder for creating a Breed entity.
func (c *BreedClient) Create() *BreedCreate {
	mutation := newBreedMutation(c.config, OpCreate)
	return &BreedCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Breed entities.
func (c *BreedClient) CreateBulk(builders ...*BreedCreate) *BreedCreateBulk {
	return &BreedCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *BreedClient) MapCreateBulk(slice any, setFunc func(*BreedCreate, int)) *BreedCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &BreedCreateBulk{err: fmt.Errorf("calling to BreedClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*BreedCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &BreedCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Breed.
func (c *BreedClient) Update() *BreedUpdate {
	mutation := newBreedMutation(c.config, OpUpdate)
	return &BreedUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BreedClient) UpdateOne(b *Breed) *BreedUpdateOne {
	mutation := newBreedMutation(c.config, OpUpdateOne, withBreed(b))
	return &BreedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BreedClient) UpdateOneID(id uuid.UUID) *BreedUpdateOne {
	mutation := newBreedMutation(c.config, OpUpdateOne, withBreedID(id))
	return &BreedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Breed.
func (c *BreedClient) Delete() *BreedDelete {
	mutation := newBreedMutation(c.config, OpDelete)
	return &BreedDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BreedClient) DeleteOne(b *Breed) *BreedDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BreedClient) DeleteOneID(id uuid.UUID) *BreedDeleteOne {
	builder := c.Delete().Where(breed.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BreedDeleteOne{builder}
}

// Query returns a query builder for Breed.
func (c *BreedClient) Query() *BreedQuery {
	return &BreedQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBreed},
		inters: c.Interceptors(),
	}
}

// Get returns a Breed entity by its id.
func (c *BreedClient) Get(ctx context.Context, id uuid.UUID) (*Breed, error) {
	return c.Query().Where(breed.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BreedClient) GetX(ctx context.Context, id uuid.UUID) *Breed {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCats queries the cats edge of a Breed.
func (c *BreedClient) QueryCats(b *Breed) *CatQuery {
	query := (&CatClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(breed.Table, breed.FieldID, id),
			sqlgraph.To(cat.Table, cat.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, breed.CatsTable, breed.CatsColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BreedClient) Hooks() []Hook {
	return c.hooks.Breed
}

// Interceptors returns the client interceptors.
func (c *BreedClient) Interceptors() []Interceptor {
	return c.inters.Breed
}

func (c *BreedClient) mutate(ctx context.Context, m *BreedMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BreedCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BreedUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BreedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BreedDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("gen: unknown Breed mutation op: %q", m.Op())
	}
}

// CatClient is a client for the Cat schema.
type CatClient struct {
	config
}

// NewCatClient returns a client for the Cat from the given config.
func NewCatClient(c config) *CatClient {
	return &CatClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cat.Hooks(f(g(h())))`.
func (c *CatClient) Use(hooks ...Hook) {
	c.hooks.Cat = append(c.hooks.Cat, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `cat.Intercept(f(g(h())))`.
func (c *CatClient) Intercept(interceptors ...Interceptor) {
	c.inters.Cat = append(c.inters.Cat, interceptors...)
}

// Create returns a builder for creating a Cat entity.
func (c *CatClient) Create() *CatCreate {
	mutation := newCatMutation(c.config, OpCreate)
	return &CatCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Cat entities.
func (c *CatClient) CreateBulk(builders ...*CatCreate) *CatCreateBulk {
	return &CatCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CatClient) MapCreateBulk(slice any, setFunc func(*CatCreate, int)) *CatCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CatCreateBulk{err: fmt.Errorf("calling to CatClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CatCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CatCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Cat.
func (c *CatClient) Update() *CatUpdate {
	mutation := newCatMutation(c.config, OpUpdate)
	return &CatUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CatClient) UpdateOne(ca *Cat) *CatUpdateOne {
	mutation := newCatMutation(c.config, OpUpdateOne, withCat(ca))
	return &CatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CatClient) UpdateOneID(id uuid.UUID) *CatUpdateOne {
	mutation := newCatMutation(c.config, OpUpdateOne, withCatID(id))
	return &CatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Cat.
func (c *CatClient) Delete() *CatDelete {
	mutation := newCatMutation(c.config, OpDelete)
	return &CatDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CatClient) DeleteOne(ca *Cat) *CatDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CatClient) DeleteOneID(id uuid.UUID) *CatDeleteOne {
	builder := c.Delete().Where(cat.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CatDeleteOne{builder}
}

// Query returns a query builder for Cat.
func (c *CatClient) Query() *CatQuery {
	return &CatQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCat},
		inters: c.Interceptors(),
	}
}

// Get returns a Cat entity by its id.
func (c *CatClient) Get(ctx context.Context, id uuid.UUID) (*Cat, error) {
	return c.Query().Where(cat.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CatClient) GetX(ctx context.Context, id uuid.UUID) *Cat {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryKittens queries the kittens edge of a Cat.
func (c *CatClient) QueryKittens(ca *Cat) *KittenQuery {
	query := (&KittenClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cat.Table, cat.FieldID, id),
			sqlgraph.To(kitten.Table, kitten.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, cat.KittensTable, cat.KittensColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBreed queries the breed edge of a Cat.
func (c *CatClient) QueryBreed(ca *Cat) *BreedQuery {
	query := (&BreedClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cat.Table, cat.FieldID, id),
			sqlgraph.To(breed.Table, breed.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cat.BreedTable, cat.BreedColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CatClient) Hooks() []Hook {
	return c.hooks.Cat
}

// Interceptors returns the client interceptors.
func (c *CatClient) Interceptors() []Interceptor {
	return c.inters.Cat
}

func (c *CatClient) mutate(ctx context.Context, m *CatMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CatCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CatUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CatDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("gen: unknown Cat mutation op: %q", m.Op())
	}
}

// FatherCatClient is a client for the FatherCat schema.
type FatherCatClient struct {
	config
}

// NewFatherCatClient returns a client for the FatherCat from the given config.
func NewFatherCatClient(c config) *FatherCatClient {
	return &FatherCatClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `fathercat.Hooks(f(g(h())))`.
func (c *FatherCatClient) Use(hooks ...Hook) {
	c.hooks.FatherCat = append(c.hooks.FatherCat, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `fathercat.Intercept(f(g(h())))`.
func (c *FatherCatClient) Intercept(interceptors ...Interceptor) {
	c.inters.FatherCat = append(c.inters.FatherCat, interceptors...)
}

// Create returns a builder for creating a FatherCat entity.
func (c *FatherCatClient) Create() *FatherCatCreate {
	mutation := newFatherCatMutation(c.config, OpCreate)
	return &FatherCatCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FatherCat entities.
func (c *FatherCatClient) CreateBulk(builders ...*FatherCatCreate) *FatherCatCreateBulk {
	return &FatherCatCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *FatherCatClient) MapCreateBulk(slice any, setFunc func(*FatherCatCreate, int)) *FatherCatCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &FatherCatCreateBulk{err: fmt.Errorf("calling to FatherCatClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*FatherCatCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &FatherCatCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FatherCat.
func (c *FatherCatClient) Update() *FatherCatUpdate {
	mutation := newFatherCatMutation(c.config, OpUpdate)
	return &FatherCatUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FatherCatClient) UpdateOne(fc *FatherCat) *FatherCatUpdateOne {
	mutation := newFatherCatMutation(c.config, OpUpdateOne, withFatherCat(fc))
	return &FatherCatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FatherCatClient) UpdateOneID(id uuid.UUID) *FatherCatUpdateOne {
	mutation := newFatherCatMutation(c.config, OpUpdateOne, withFatherCatID(id))
	return &FatherCatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FatherCat.
func (c *FatherCatClient) Delete() *FatherCatDelete {
	mutation := newFatherCatMutation(c.config, OpDelete)
	return &FatherCatDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FatherCatClient) DeleteOne(fc *FatherCat) *FatherCatDeleteOne {
	return c.DeleteOneID(fc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FatherCatClient) DeleteOneID(id uuid.UUID) *FatherCatDeleteOne {
	builder := c.Delete().Where(fathercat.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FatherCatDeleteOne{builder}
}

// Query returns a query builder for FatherCat.
func (c *FatherCatClient) Query() *FatherCatQuery {
	return &FatherCatQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFatherCat},
		inters: c.Interceptors(),
	}
}

// Get returns a FatherCat entity by its id.
func (c *FatherCatClient) Get(ctx context.Context, id uuid.UUID) (*FatherCat, error) {
	return c.Query().Where(fathercat.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FatherCatClient) GetX(ctx context.Context, id uuid.UUID) *FatherCat {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *FatherCatClient) Hooks() []Hook {
	return c.hooks.FatherCat
}

// Interceptors returns the client interceptors.
func (c *FatherCatClient) Interceptors() []Interceptor {
	return c.inters.FatherCat
}

func (c *FatherCatClient) mutate(ctx context.Context, m *FatherCatMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FatherCatCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FatherCatUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FatherCatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FatherCatDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("gen: unknown FatherCat mutation op: %q", m.Op())
	}
}

// KittenClient is a client for the Kitten schema.
type KittenClient struct {
	config
}

// NewKittenClient returns a client for the Kitten from the given config.
func NewKittenClient(c config) *KittenClient {
	return &KittenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `kitten.Hooks(f(g(h())))`.
func (c *KittenClient) Use(hooks ...Hook) {
	c.hooks.Kitten = append(c.hooks.Kitten, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `kitten.Intercept(f(g(h())))`.
func (c *KittenClient) Intercept(interceptors ...Interceptor) {
	c.inters.Kitten = append(c.inters.Kitten, interceptors...)
}

// Create returns a builder for creating a Kitten entity.
func (c *KittenClient) Create() *KittenCreate {
	mutation := newKittenMutation(c.config, OpCreate)
	return &KittenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Kitten entities.
func (c *KittenClient) CreateBulk(builders ...*KittenCreate) *KittenCreateBulk {
	return &KittenCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *KittenClient) MapCreateBulk(slice any, setFunc func(*KittenCreate, int)) *KittenCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &KittenCreateBulk{err: fmt.Errorf("calling to KittenClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*KittenCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &KittenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Kitten.
func (c *KittenClient) Update() *KittenUpdate {
	mutation := newKittenMutation(c.config, OpUpdate)
	return &KittenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *KittenClient) UpdateOne(k *Kitten) *KittenUpdateOne {
	mutation := newKittenMutation(c.config, OpUpdateOne, withKitten(k))
	return &KittenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *KittenClient) UpdateOneID(id uuid.UUID) *KittenUpdateOne {
	mutation := newKittenMutation(c.config, OpUpdateOne, withKittenID(id))
	return &KittenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Kitten.
func (c *KittenClient) Delete() *KittenDelete {
	mutation := newKittenMutation(c.config, OpDelete)
	return &KittenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *KittenClient) DeleteOne(k *Kitten) *KittenDeleteOne {
	return c.DeleteOneID(k.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *KittenClient) DeleteOneID(id uuid.UUID) *KittenDeleteOne {
	builder := c.Delete().Where(kitten.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &KittenDeleteOne{builder}
}

// Query returns a query builder for Kitten.
func (c *KittenClient) Query() *KittenQuery {
	return &KittenQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeKitten},
		inters: c.Interceptors(),
	}
}

// Get returns a Kitten entity by its id.
func (c *KittenClient) Get(ctx context.Context, id uuid.UUID) (*Kitten, error) {
	return c.Query().Where(kitten.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *KittenClient) GetX(ctx context.Context, id uuid.UUID) *Kitten {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMother queries the mother edge of a Kitten.
func (c *KittenClient) QueryMother(k *Kitten) *CatQuery {
	query := (&CatClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := k.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(kitten.Table, kitten.FieldID, id),
			sqlgraph.To(cat.Table, cat.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, kitten.MotherTable, kitten.MotherColumn),
		)
		fromV = sqlgraph.Neighbors(k.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *KittenClient) Hooks() []Hook {
	return c.hooks.Kitten
}

// Interceptors returns the client interceptors.
func (c *KittenClient) Interceptors() []Interceptor {
	return c.inters.Kitten
}

func (c *KittenClient) mutate(ctx context.Context, m *KittenMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&KittenCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&KittenUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&KittenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&KittenDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("gen: unknown Kitten mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Breed, Cat, FatherCat, Kitten []ent.Hook
	}
	inters struct {
		Breed, Cat, FatherCat, Kitten []ent.Interceptor
	}
)

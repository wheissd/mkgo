// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/wheissd/mkgo/internal/examples/new/ent/gen/defaultmodel"
	"github.com/wheissd/mkgo/internal/examples/new/ent/gen/predicate"
)

// DefaultModelQuery is the builder for querying DefaultModel entities.
type DefaultModelQuery struct {
	config
	ctx        *QueryContext
	order      []defaultmodel.OrderOption
	inters     []Interceptor
	predicates []predicate.DefaultModel
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DefaultModelQuery builder.
func (dmq *DefaultModelQuery) Where(ps ...predicate.DefaultModel) *DefaultModelQuery {
	dmq.predicates = append(dmq.predicates, ps...)
	return dmq
}

// Limit the number of records to be returned by this query.
func (dmq *DefaultModelQuery) Limit(limit int) *DefaultModelQuery {
	dmq.ctx.Limit = &limit
	return dmq
}

// Offset to start from.
func (dmq *DefaultModelQuery) Offset(offset int) *DefaultModelQuery {
	dmq.ctx.Offset = &offset
	return dmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dmq *DefaultModelQuery) Unique(unique bool) *DefaultModelQuery {
	dmq.ctx.Unique = &unique
	return dmq
}

// Order specifies how the records should be ordered.
func (dmq *DefaultModelQuery) Order(o ...defaultmodel.OrderOption) *DefaultModelQuery {
	dmq.order = append(dmq.order, o...)
	return dmq
}

// First returns the first DefaultModel entity from the query.
// Returns a *NotFoundError when no DefaultModel was found.
func (dmq *DefaultModelQuery) First(ctx context.Context) (*DefaultModel, error) {
	nodes, err := dmq.Limit(1).All(setContextOp(ctx, dmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{defaultmodel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dmq *DefaultModelQuery) FirstX(ctx context.Context) *DefaultModel {
	node, err := dmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DefaultModel ID from the query.
// Returns a *NotFoundError when no DefaultModel ID was found.
func (dmq *DefaultModelQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dmq.Limit(1).IDs(setContextOp(ctx, dmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{defaultmodel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dmq *DefaultModelQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DefaultModel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DefaultModel entity is found.
// Returns a *NotFoundError when no DefaultModel entities are found.
func (dmq *DefaultModelQuery) Only(ctx context.Context) (*DefaultModel, error) {
	nodes, err := dmq.Limit(2).All(setContextOp(ctx, dmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{defaultmodel.Label}
	default:
		return nil, &NotSingularError{defaultmodel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dmq *DefaultModelQuery) OnlyX(ctx context.Context) *DefaultModel {
	node, err := dmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DefaultModel ID in the query.
// Returns a *NotSingularError when more than one DefaultModel ID is found.
// Returns a *NotFoundError when no entities are found.
func (dmq *DefaultModelQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dmq.Limit(2).IDs(setContextOp(ctx, dmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{defaultmodel.Label}
	default:
		err = &NotSingularError{defaultmodel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dmq *DefaultModelQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DefaultModels.
func (dmq *DefaultModelQuery) All(ctx context.Context) ([]*DefaultModel, error) {
	ctx = setContextOp(ctx, dmq.ctx, "All")
	if err := dmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DefaultModel, *DefaultModelQuery]()
	return withInterceptors[[]*DefaultModel](ctx, dmq, qr, dmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dmq *DefaultModelQuery) AllX(ctx context.Context) []*DefaultModel {
	nodes, err := dmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DefaultModel IDs.
func (dmq *DefaultModelQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if dmq.ctx.Unique == nil && dmq.path != nil {
		dmq.Unique(true)
	}
	ctx = setContextOp(ctx, dmq.ctx, "IDs")
	if err = dmq.Select(defaultmodel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dmq *DefaultModelQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dmq *DefaultModelQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dmq.ctx, "Count")
	if err := dmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dmq, querierCount[*DefaultModelQuery](), dmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dmq *DefaultModelQuery) CountX(ctx context.Context) int {
	count, err := dmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dmq *DefaultModelQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dmq.ctx, "Exist")
	switch _, err := dmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("gen: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dmq *DefaultModelQuery) ExistX(ctx context.Context) bool {
	exist, err := dmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DefaultModelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dmq *DefaultModelQuery) Clone() *DefaultModelQuery {
	if dmq == nil {
		return nil
	}
	return &DefaultModelQuery{
		config:     dmq.config,
		ctx:        dmq.ctx.Clone(),
		order:      append([]defaultmodel.OrderOption{}, dmq.order...),
		inters:     append([]Interceptor{}, dmq.inters...),
		predicates: append([]predicate.DefaultModel{}, dmq.predicates...),
		// clone intermediate query.
		sql:  dmq.sql.Clone(),
		path: dmq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DefaultModel.Query().
//		GroupBy(defaultmodel.FieldName).
//		Aggregate(gen.Count()).
//		Scan(ctx, &v)
func (dmq *DefaultModelQuery) GroupBy(field string, fields ...string) *DefaultModelGroupBy {
	dmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DefaultModelGroupBy{build: dmq}
	grbuild.flds = &dmq.ctx.Fields
	grbuild.label = defaultmodel.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.DefaultModel.Query().
//		Select(defaultmodel.FieldName).
//		Scan(ctx, &v)
func (dmq *DefaultModelQuery) Select(fields ...string) *DefaultModelSelect {
	dmq.ctx.Fields = append(dmq.ctx.Fields, fields...)
	sbuild := &DefaultModelSelect{DefaultModelQuery: dmq}
	sbuild.label = defaultmodel.Label
	sbuild.flds, sbuild.scan = &dmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DefaultModelSelect configured with the given aggregations.
func (dmq *DefaultModelQuery) Aggregate(fns ...AggregateFunc) *DefaultModelSelect {
	return dmq.Select().Aggregate(fns...)
}

func (dmq *DefaultModelQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dmq.inters {
		if inter == nil {
			return fmt.Errorf("gen: uninitialized interceptor (forgotten import gen/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dmq); err != nil {
				return err
			}
		}
	}
	for _, f := range dmq.ctx.Fields {
		if !defaultmodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
		}
	}
	if dmq.path != nil {
		prev, err := dmq.path(ctx)
		if err != nil {
			return err
		}
		dmq.sql = prev
	}
	return nil
}

func (dmq *DefaultModelQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DefaultModel, error) {
	var (
		nodes = []*DefaultModel{}
		_spec = dmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DefaultModel).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DefaultModel{config: dmq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dmq *DefaultModelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dmq.querySpec()
	_spec.Node.Columns = dmq.ctx.Fields
	if len(dmq.ctx.Fields) > 0 {
		_spec.Unique = dmq.ctx.Unique != nil && *dmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dmq.driver, _spec)
}

func (dmq *DefaultModelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(defaultmodel.Table, defaultmodel.Columns, sqlgraph.NewFieldSpec(defaultmodel.FieldID, field.TypeUUID))
	_spec.From = dmq.sql
	if unique := dmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dmq.path != nil {
		_spec.Unique = true
	}
	if fields := dmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, defaultmodel.FieldID)
		for i := range fields {
			if fields[i] != defaultmodel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dmq *DefaultModelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dmq.driver.Dialect())
	t1 := builder.Table(defaultmodel.Table)
	columns := dmq.ctx.Fields
	if len(columns) == 0 {
		columns = defaultmodel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dmq.sql != nil {
		selector = dmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dmq.ctx.Unique != nil && *dmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dmq.predicates {
		p(selector)
	}
	for _, p := range dmq.order {
		p(selector)
	}
	if offset := dmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DefaultModelGroupBy is the group-by builder for DefaultModel entities.
type DefaultModelGroupBy struct {
	selector
	build *DefaultModelQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dmgb *DefaultModelGroupBy) Aggregate(fns ...AggregateFunc) *DefaultModelGroupBy {
	dmgb.fns = append(dmgb.fns, fns...)
	return dmgb
}

// Scan applies the selector query and scans the result into the given value.
func (dmgb *DefaultModelGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dmgb.build.ctx, "GroupBy")
	if err := dmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DefaultModelQuery, *DefaultModelGroupBy](ctx, dmgb.build, dmgb, dmgb.build.inters, v)
}

func (dmgb *DefaultModelGroupBy) sqlScan(ctx context.Context, root *DefaultModelQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dmgb.fns))
	for _, fn := range dmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dmgb.flds)+len(dmgb.fns))
		for _, f := range *dmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DefaultModelSelect is the builder for selecting fields of DefaultModel entities.
type DefaultModelSelect struct {
	*DefaultModelQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dms *DefaultModelSelect) Aggregate(fns ...AggregateFunc) *DefaultModelSelect {
	dms.fns = append(dms.fns, fns...)
	return dms
}

// Scan applies the selector query and scans the result into the given value.
func (dms *DefaultModelSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dms.ctx, "Select")
	if err := dms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DefaultModelQuery, *DefaultModelSelect](ctx, dms.DefaultModelQuery, dms, dms.inters, v)
}

func (dms *DefaultModelSelect) sqlScan(ctx context.Context, root *DefaultModelQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dms.fns))
	for _, fn := range dms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
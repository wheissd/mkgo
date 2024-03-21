// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/breed"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/cat"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/predicate"
)

// BreedUpdate is the builder for updating Breed entities.
type BreedUpdate struct {
	config
	hooks    []Hook
	mutation *BreedMutation
}

// Where appends a list predicates to the BreedUpdate builder.
func (bu *BreedUpdate) Where(ps ...predicate.Breed) *BreedUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetDeletedTime sets the "deleted_time" field.
func (bu *BreedUpdate) SetDeletedTime(t time.Time) *BreedUpdate {
	bu.mutation.SetDeletedTime(t)
	return bu
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (bu *BreedUpdate) SetNillableDeletedTime(t *time.Time) *BreedUpdate {
	if t != nil {
		bu.SetDeletedTime(*t)
	}
	return bu
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (bu *BreedUpdate) ClearDeletedTime() *BreedUpdate {
	bu.mutation.ClearDeletedTime()
	return bu
}

// SetUpdateTime sets the "update_time" field.
func (bu *BreedUpdate) SetUpdateTime(t time.Time) *BreedUpdate {
	bu.mutation.SetUpdateTime(t)
	return bu
}

// SetName sets the "name" field.
func (bu *BreedUpdate) SetName(s string) *BreedUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bu *BreedUpdate) SetNillableName(s *string) *BreedUpdate {
	if s != nil {
		bu.SetName(*s)
	}
	return bu
}

// AddCatIDs adds the "cats" edge to the Cat entity by IDs.
func (bu *BreedUpdate) AddCatIDs(ids ...uuid.UUID) *BreedUpdate {
	bu.mutation.AddCatIDs(ids...)
	return bu
}

// AddCats adds the "cats" edges to the Cat entity.
func (bu *BreedUpdate) AddCats(c ...*Cat) *BreedUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bu.AddCatIDs(ids...)
}

// Mutation returns the BreedMutation object of the builder.
func (bu *BreedUpdate) Mutation() *BreedMutation {
	return bu.mutation
}

// ClearCats clears all "cats" edges to the Cat entity.
func (bu *BreedUpdate) ClearCats() *BreedUpdate {
	bu.mutation.ClearCats()
	return bu
}

// RemoveCatIDs removes the "cats" edge to Cat entities by IDs.
func (bu *BreedUpdate) RemoveCatIDs(ids ...uuid.UUID) *BreedUpdate {
	bu.mutation.RemoveCatIDs(ids...)
	return bu
}

// RemoveCats removes "cats" edges to Cat entities.
func (bu *BreedUpdate) RemoveCats(c ...*Cat) *BreedUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bu.RemoveCatIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BreedUpdate) Save(ctx context.Context) (int, error) {
	bu.defaults()
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BreedUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BreedUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BreedUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BreedUpdate) defaults() {
	if _, ok := bu.mutation.UpdateTime(); !ok {
		v := breed.UpdateDefaultUpdateTime()
		bu.mutation.SetUpdateTime(v)
	}
}

func (bu *BreedUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(breed.Table, breed.Columns, sqlgraph.NewFieldSpec(breed.FieldID, field.TypeUUID))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.DeletedTime(); ok {
		_spec.SetField(breed.FieldDeletedTime, field.TypeTime, value)
	}
	if bu.mutation.DeletedTimeCleared() {
		_spec.ClearField(breed.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := bu.mutation.UpdateTime(); ok {
		_spec.SetField(breed.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.SetField(breed.FieldName, field.TypeString, value)
	}
	if bu.mutation.CatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   breed.CatsTable,
			Columns: []string{breed.CatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cat.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedCatsIDs(); len(nodes) > 0 && !bu.mutation.CatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   breed.CatsTable,
			Columns: []string{breed.CatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cat.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.CatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   breed.CatsTable,
			Columns: []string{breed.CatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cat.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{breed.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BreedUpdateOne is the builder for updating a single Breed entity.
type BreedUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BreedMutation
}

// SetDeletedTime sets the "deleted_time" field.
func (buo *BreedUpdateOne) SetDeletedTime(t time.Time) *BreedUpdateOne {
	buo.mutation.SetDeletedTime(t)
	return buo
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (buo *BreedUpdateOne) SetNillableDeletedTime(t *time.Time) *BreedUpdateOne {
	if t != nil {
		buo.SetDeletedTime(*t)
	}
	return buo
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (buo *BreedUpdateOne) ClearDeletedTime() *BreedUpdateOne {
	buo.mutation.ClearDeletedTime()
	return buo
}

// SetUpdateTime sets the "update_time" field.
func (buo *BreedUpdateOne) SetUpdateTime(t time.Time) *BreedUpdateOne {
	buo.mutation.SetUpdateTime(t)
	return buo
}

// SetName sets the "name" field.
func (buo *BreedUpdateOne) SetName(s string) *BreedUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (buo *BreedUpdateOne) SetNillableName(s *string) *BreedUpdateOne {
	if s != nil {
		buo.SetName(*s)
	}
	return buo
}

// AddCatIDs adds the "cats" edge to the Cat entity by IDs.
func (buo *BreedUpdateOne) AddCatIDs(ids ...uuid.UUID) *BreedUpdateOne {
	buo.mutation.AddCatIDs(ids...)
	return buo
}

// AddCats adds the "cats" edges to the Cat entity.
func (buo *BreedUpdateOne) AddCats(c ...*Cat) *BreedUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return buo.AddCatIDs(ids...)
}

// Mutation returns the BreedMutation object of the builder.
func (buo *BreedUpdateOne) Mutation() *BreedMutation {
	return buo.mutation
}

// ClearCats clears all "cats" edges to the Cat entity.
func (buo *BreedUpdateOne) ClearCats() *BreedUpdateOne {
	buo.mutation.ClearCats()
	return buo
}

// RemoveCatIDs removes the "cats" edge to Cat entities by IDs.
func (buo *BreedUpdateOne) RemoveCatIDs(ids ...uuid.UUID) *BreedUpdateOne {
	buo.mutation.RemoveCatIDs(ids...)
	return buo
}

// RemoveCats removes "cats" edges to Cat entities.
func (buo *BreedUpdateOne) RemoveCats(c ...*Cat) *BreedUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return buo.RemoveCatIDs(ids...)
}

// Where appends a list predicates to the BreedUpdate builder.
func (buo *BreedUpdateOne) Where(ps ...predicate.Breed) *BreedUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BreedUpdateOne) Select(field string, fields ...string) *BreedUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Breed entity.
func (buo *BreedUpdateOne) Save(ctx context.Context) (*Breed, error) {
	buo.defaults()
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BreedUpdateOne) SaveX(ctx context.Context) *Breed {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BreedUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BreedUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BreedUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdateTime(); !ok {
		v := breed.UpdateDefaultUpdateTime()
		buo.mutation.SetUpdateTime(v)
	}
}

func (buo *BreedUpdateOne) sqlSave(ctx context.Context) (_node *Breed, err error) {
	_spec := sqlgraph.NewUpdateSpec(breed.Table, breed.Columns, sqlgraph.NewFieldSpec(breed.FieldID, field.TypeUUID))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`gen: missing "Breed.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, breed.FieldID)
		for _, f := range fields {
			if !breed.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
			}
			if f != breed.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.DeletedTime(); ok {
		_spec.SetField(breed.FieldDeletedTime, field.TypeTime, value)
	}
	if buo.mutation.DeletedTimeCleared() {
		_spec.ClearField(breed.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := buo.mutation.UpdateTime(); ok {
		_spec.SetField(breed.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.SetField(breed.FieldName, field.TypeString, value)
	}
	if buo.mutation.CatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   breed.CatsTable,
			Columns: []string{breed.CatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cat.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedCatsIDs(); len(nodes) > 0 && !buo.mutation.CatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   breed.CatsTable,
			Columns: []string{breed.CatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cat.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.CatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   breed.CatsTable,
			Columns: []string{breed.CatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cat.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Breed{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{breed.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
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
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/cat"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/kitten"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/predicate"
)

// KittenUpdate is the builder for updating Kitten entities.
type KittenUpdate struct {
	config
	hooks    []Hook
	mutation *KittenMutation
}

// Where appends a list predicates to the KittenUpdate builder.
func (ku *KittenUpdate) Where(ps ...predicate.Kitten) *KittenUpdate {
	ku.mutation.Where(ps...)
	return ku
}

// SetDeletedTime sets the "deleted_time" field.
func (ku *KittenUpdate) SetDeletedTime(t time.Time) *KittenUpdate {
	ku.mutation.SetDeletedTime(t)
	return ku
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (ku *KittenUpdate) SetNillableDeletedTime(t *time.Time) *KittenUpdate {
	if t != nil {
		ku.SetDeletedTime(*t)
	}
	return ku
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (ku *KittenUpdate) ClearDeletedTime() *KittenUpdate {
	ku.mutation.ClearDeletedTime()
	return ku
}

// SetUpdateTime sets the "update_time" field.
func (ku *KittenUpdate) SetUpdateTime(t time.Time) *KittenUpdate {
	ku.mutation.SetUpdateTime(t)
	return ku
}

// SetName sets the "name" field.
func (ku *KittenUpdate) SetName(s string) *KittenUpdate {
	ku.mutation.SetName(s)
	return ku
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ku *KittenUpdate) SetNillableName(s *string) *KittenUpdate {
	if s != nil {
		ku.SetName(*s)
	}
	return ku
}

// SetMotherID sets the "mother_id" field.
func (ku *KittenUpdate) SetMotherID(u uuid.UUID) *KittenUpdate {
	ku.mutation.SetMotherID(u)
	return ku
}

// SetNillableMotherID sets the "mother_id" field if the given value is not nil.
func (ku *KittenUpdate) SetNillableMotherID(u *uuid.UUID) *KittenUpdate {
	if u != nil {
		ku.SetMotherID(*u)
	}
	return ku
}

// SetMother sets the "mother" edge to the Cat entity.
func (ku *KittenUpdate) SetMother(c *Cat) *KittenUpdate {
	return ku.SetMotherID(c.ID)
}

// Mutation returns the KittenMutation object of the builder.
func (ku *KittenUpdate) Mutation() *KittenMutation {
	return ku.mutation
}

// ClearMother clears the "mother" edge to the Cat entity.
func (ku *KittenUpdate) ClearMother() *KittenUpdate {
	ku.mutation.ClearMother()
	return ku
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ku *KittenUpdate) Save(ctx context.Context) (int, error) {
	ku.defaults()
	return withHooks(ctx, ku.sqlSave, ku.mutation, ku.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ku *KittenUpdate) SaveX(ctx context.Context) int {
	affected, err := ku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ku *KittenUpdate) Exec(ctx context.Context) error {
	_, err := ku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ku *KittenUpdate) ExecX(ctx context.Context) {
	if err := ku.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ku *KittenUpdate) defaults() {
	if _, ok := ku.mutation.UpdateTime(); !ok {
		v := kitten.UpdateDefaultUpdateTime()
		ku.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ku *KittenUpdate) check() error {
	if _, ok := ku.mutation.MotherID(); ku.mutation.MotherCleared() && !ok {
		return errors.New(`gen: clearing a required unique edge "Kitten.mother"`)
	}
	return nil
}

func (ku *KittenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ku.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(kitten.Table, kitten.Columns, sqlgraph.NewFieldSpec(kitten.FieldID, field.TypeUUID))
	if ps := ku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ku.mutation.DeletedTime(); ok {
		_spec.SetField(kitten.FieldDeletedTime, field.TypeTime, value)
	}
	if ku.mutation.DeletedTimeCleared() {
		_spec.ClearField(kitten.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := ku.mutation.UpdateTime(); ok {
		_spec.SetField(kitten.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ku.mutation.Name(); ok {
		_spec.SetField(kitten.FieldName, field.TypeString, value)
	}
	if ku.mutation.MotherCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kitten.MotherTable,
			Columns: []string{kitten.MotherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cat.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ku.mutation.MotherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kitten.MotherTable,
			Columns: []string{kitten.MotherColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kitten.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ku.mutation.done = true
	return n, nil
}

// KittenUpdateOne is the builder for updating a single Kitten entity.
type KittenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *KittenMutation
}

// SetDeletedTime sets the "deleted_time" field.
func (kuo *KittenUpdateOne) SetDeletedTime(t time.Time) *KittenUpdateOne {
	kuo.mutation.SetDeletedTime(t)
	return kuo
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (kuo *KittenUpdateOne) SetNillableDeletedTime(t *time.Time) *KittenUpdateOne {
	if t != nil {
		kuo.SetDeletedTime(*t)
	}
	return kuo
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (kuo *KittenUpdateOne) ClearDeletedTime() *KittenUpdateOne {
	kuo.mutation.ClearDeletedTime()
	return kuo
}

// SetUpdateTime sets the "update_time" field.
func (kuo *KittenUpdateOne) SetUpdateTime(t time.Time) *KittenUpdateOne {
	kuo.mutation.SetUpdateTime(t)
	return kuo
}

// SetName sets the "name" field.
func (kuo *KittenUpdateOne) SetName(s string) *KittenUpdateOne {
	kuo.mutation.SetName(s)
	return kuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (kuo *KittenUpdateOne) SetNillableName(s *string) *KittenUpdateOne {
	if s != nil {
		kuo.SetName(*s)
	}
	return kuo
}

// SetMotherID sets the "mother_id" field.
func (kuo *KittenUpdateOne) SetMotherID(u uuid.UUID) *KittenUpdateOne {
	kuo.mutation.SetMotherID(u)
	return kuo
}

// SetNillableMotherID sets the "mother_id" field if the given value is not nil.
func (kuo *KittenUpdateOne) SetNillableMotherID(u *uuid.UUID) *KittenUpdateOne {
	if u != nil {
		kuo.SetMotherID(*u)
	}
	return kuo
}

// SetMother sets the "mother" edge to the Cat entity.
func (kuo *KittenUpdateOne) SetMother(c *Cat) *KittenUpdateOne {
	return kuo.SetMotherID(c.ID)
}

// Mutation returns the KittenMutation object of the builder.
func (kuo *KittenUpdateOne) Mutation() *KittenMutation {
	return kuo.mutation
}

// ClearMother clears the "mother" edge to the Cat entity.
func (kuo *KittenUpdateOne) ClearMother() *KittenUpdateOne {
	kuo.mutation.ClearMother()
	return kuo
}

// Where appends a list predicates to the KittenUpdate builder.
func (kuo *KittenUpdateOne) Where(ps ...predicate.Kitten) *KittenUpdateOne {
	kuo.mutation.Where(ps...)
	return kuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kuo *KittenUpdateOne) Select(field string, fields ...string) *KittenUpdateOne {
	kuo.fields = append([]string{field}, fields...)
	return kuo
}

// Save executes the query and returns the updated Kitten entity.
func (kuo *KittenUpdateOne) Save(ctx context.Context) (*Kitten, error) {
	kuo.defaults()
	return withHooks(ctx, kuo.sqlSave, kuo.mutation, kuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (kuo *KittenUpdateOne) SaveX(ctx context.Context) *Kitten {
	node, err := kuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kuo *KittenUpdateOne) Exec(ctx context.Context) error {
	_, err := kuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kuo *KittenUpdateOne) ExecX(ctx context.Context) {
	if err := kuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kuo *KittenUpdateOne) defaults() {
	if _, ok := kuo.mutation.UpdateTime(); !ok {
		v := kitten.UpdateDefaultUpdateTime()
		kuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kuo *KittenUpdateOne) check() error {
	if _, ok := kuo.mutation.MotherID(); kuo.mutation.MotherCleared() && !ok {
		return errors.New(`gen: clearing a required unique edge "Kitten.mother"`)
	}
	return nil
}

func (kuo *KittenUpdateOne) sqlSave(ctx context.Context) (_node *Kitten, err error) {
	if err := kuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(kitten.Table, kitten.Columns, sqlgraph.NewFieldSpec(kitten.FieldID, field.TypeUUID))
	id, ok := kuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`gen: missing "Kitten.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := kuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, kitten.FieldID)
		for _, f := range fields {
			if !kitten.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
			}
			if f != kitten.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kuo.mutation.DeletedTime(); ok {
		_spec.SetField(kitten.FieldDeletedTime, field.TypeTime, value)
	}
	if kuo.mutation.DeletedTimeCleared() {
		_spec.ClearField(kitten.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := kuo.mutation.UpdateTime(); ok {
		_spec.SetField(kitten.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := kuo.mutation.Name(); ok {
		_spec.SetField(kitten.FieldName, field.TypeString, value)
	}
	if kuo.mutation.MotherCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kitten.MotherTable,
			Columns: []string{kitten.MotherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cat.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kuo.mutation.MotherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kitten.MotherTable,
			Columns: []string{kitten.MotherColumn},
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
	_node = &Kitten{config: kuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kitten.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	kuo.mutation.done = true
	return _node, nil
}

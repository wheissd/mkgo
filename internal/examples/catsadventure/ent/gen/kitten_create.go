// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/cat"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen/kitten"
)

// KittenCreate is the builder for creating a Kitten entity.
type KittenCreate struct {
	config
	mutation *KittenMutation
	hooks    []Hook
}

// SetDeletedTime sets the "deleted_time" field.
func (kc *KittenCreate) SetDeletedTime(t time.Time) *KittenCreate {
	kc.mutation.SetDeletedTime(t)
	return kc
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (kc *KittenCreate) SetNillableDeletedTime(t *time.Time) *KittenCreate {
	if t != nil {
		kc.SetDeletedTime(*t)
	}
	return kc
}

// SetCreateTime sets the "create_time" field.
func (kc *KittenCreate) SetCreateTime(t time.Time) *KittenCreate {
	kc.mutation.SetCreateTime(t)
	return kc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (kc *KittenCreate) SetNillableCreateTime(t *time.Time) *KittenCreate {
	if t != nil {
		kc.SetCreateTime(*t)
	}
	return kc
}

// SetUpdateTime sets the "update_time" field.
func (kc *KittenCreate) SetUpdateTime(t time.Time) *KittenCreate {
	kc.mutation.SetUpdateTime(t)
	return kc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (kc *KittenCreate) SetNillableUpdateTime(t *time.Time) *KittenCreate {
	if t != nil {
		kc.SetUpdateTime(*t)
	}
	return kc
}

// SetName sets the "name" field.
func (kc *KittenCreate) SetName(s string) *KittenCreate {
	kc.mutation.SetName(s)
	return kc
}

// SetMotherID sets the "mother_id" field.
func (kc *KittenCreate) SetMotherID(u uuid.UUID) *KittenCreate {
	kc.mutation.SetMotherID(u)
	return kc
}

// SetID sets the "id" field.
func (kc *KittenCreate) SetID(u uuid.UUID) *KittenCreate {
	kc.mutation.SetID(u)
	return kc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (kc *KittenCreate) SetNillableID(u *uuid.UUID) *KittenCreate {
	if u != nil {
		kc.SetID(*u)
	}
	return kc
}

// SetMother sets the "mother" edge to the Cat entity.
func (kc *KittenCreate) SetMother(c *Cat) *KittenCreate {
	return kc.SetMotherID(c.ID)
}

// Mutation returns the KittenMutation object of the builder.
func (kc *KittenCreate) Mutation() *KittenMutation {
	return kc.mutation
}

// Save creates the Kitten in the database.
func (kc *KittenCreate) Save(ctx context.Context) (*Kitten, error) {
	kc.defaults()
	return withHooks(ctx, kc.sqlSave, kc.mutation, kc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (kc *KittenCreate) SaveX(ctx context.Context) *Kitten {
	v, err := kc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kc *KittenCreate) Exec(ctx context.Context) error {
	_, err := kc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kc *KittenCreate) ExecX(ctx context.Context) {
	if err := kc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kc *KittenCreate) defaults() {
	if _, ok := kc.mutation.CreateTime(); !ok {
		v := kitten.DefaultCreateTime()
		kc.mutation.SetCreateTime(v)
	}
	if _, ok := kc.mutation.UpdateTime(); !ok {
		v := kitten.DefaultUpdateTime()
		kc.mutation.SetUpdateTime(v)
	}
	if _, ok := kc.mutation.ID(); !ok {
		v := kitten.DefaultID()
		kc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kc *KittenCreate) check() error {
	if _, ok := kc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`gen: missing required field "Kitten.create_time"`)}
	}
	if _, ok := kc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`gen: missing required field "Kitten.update_time"`)}
	}
	if _, ok := kc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`gen: missing required field "Kitten.name"`)}
	}
	if _, ok := kc.mutation.MotherID(); !ok {
		return &ValidationError{Name: "mother_id", err: errors.New(`gen: missing required field "Kitten.mother_id"`)}
	}
	if _, ok := kc.mutation.MotherID(); !ok {
		return &ValidationError{Name: "mother", err: errors.New(`gen: missing required edge "Kitten.mother"`)}
	}
	return nil
}

func (kc *KittenCreate) sqlSave(ctx context.Context) (*Kitten, error) {
	if err := kc.check(); err != nil {
		return nil, err
	}
	_node, _spec := kc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	kc.mutation.id = &_node.ID
	kc.mutation.done = true
	return _node, nil
}

func (kc *KittenCreate) createSpec() (*Kitten, *sqlgraph.CreateSpec) {
	var (
		_node = &Kitten{config: kc.config}
		_spec = sqlgraph.NewCreateSpec(kitten.Table, sqlgraph.NewFieldSpec(kitten.FieldID, field.TypeUUID))
	)
	if id, ok := kc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := kc.mutation.DeletedTime(); ok {
		_spec.SetField(kitten.FieldDeletedTime, field.TypeTime, value)
		_node.DeletedTime = &value
	}
	if value, ok := kc.mutation.CreateTime(); ok {
		_spec.SetField(kitten.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := kc.mutation.UpdateTime(); ok {
		_spec.SetField(kitten.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := kc.mutation.Name(); ok {
		_spec.SetField(kitten.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := kc.mutation.MotherIDs(); len(nodes) > 0 {
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
		_node.MotherID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// KittenCreateBulk is the builder for creating many Kitten entities in bulk.
type KittenCreateBulk struct {
	config
	err      error
	builders []*KittenCreate
}

// Save creates the Kitten entities in the database.
func (kcb *KittenCreateBulk) Save(ctx context.Context) ([]*Kitten, error) {
	if kcb.err != nil {
		return nil, kcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(kcb.builders))
	nodes := make([]*Kitten, len(kcb.builders))
	mutators := make([]Mutator, len(kcb.builders))
	for i := range kcb.builders {
		func(i int, root context.Context) {
			builder := kcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KittenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, kcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, kcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kcb *KittenCreateBulk) SaveX(ctx context.Context) []*Kitten {
	v, err := kcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kcb *KittenCreateBulk) Exec(ctx context.Context) error {
	_, err := kcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kcb *KittenCreateBulk) ExecX(ctx context.Context) {
	if err := kcb.Exec(ctx); err != nil {
		panic(err)
	}
}

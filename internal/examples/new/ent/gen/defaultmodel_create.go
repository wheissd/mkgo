// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/wheissd/mkgo/internal/examples/new/ent/gen/defaultmodel"
)

// DefaultModelCreate is the builder for creating a DefaultModel entity.
type DefaultModelCreate struct {
	config
	mutation *DefaultModelMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (dmc *DefaultModelCreate) SetName(s string) *DefaultModelCreate {
	dmc.mutation.SetName(s)
	return dmc
}

// SetID sets the "id" field.
func (dmc *DefaultModelCreate) SetID(u uuid.UUID) *DefaultModelCreate {
	dmc.mutation.SetID(u)
	return dmc
}

// Mutation returns the DefaultModelMutation object of the builder.
func (dmc *DefaultModelCreate) Mutation() *DefaultModelMutation {
	return dmc.mutation
}

// Save creates the DefaultModel in the database.
func (dmc *DefaultModelCreate) Save(ctx context.Context) (*DefaultModel, error) {
	return withHooks(ctx, dmc.sqlSave, dmc.mutation, dmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dmc *DefaultModelCreate) SaveX(ctx context.Context) *DefaultModel {
	v, err := dmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dmc *DefaultModelCreate) Exec(ctx context.Context) error {
	_, err := dmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmc *DefaultModelCreate) ExecX(ctx context.Context) {
	if err := dmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dmc *DefaultModelCreate) check() error {
	if _, ok := dmc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`gen: missing required field "DefaultModel.name"`)}
	}
	return nil
}

func (dmc *DefaultModelCreate) sqlSave(ctx context.Context) (*DefaultModel, error) {
	if err := dmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dmc.driver, _spec); err != nil {
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
	dmc.mutation.id = &_node.ID
	dmc.mutation.done = true
	return _node, nil
}

func (dmc *DefaultModelCreate) createSpec() (*DefaultModel, *sqlgraph.CreateSpec) {
	var (
		_node = &DefaultModel{config: dmc.config}
		_spec = sqlgraph.NewCreateSpec(defaultmodel.Table, sqlgraph.NewFieldSpec(defaultmodel.FieldID, field.TypeUUID))
	)
	if id, ok := dmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dmc.mutation.Name(); ok {
		_spec.SetField(defaultmodel.FieldName, field.TypeString, value)
		_node.Name = value
	}
	return _node, _spec
}

// DefaultModelCreateBulk is the builder for creating many DefaultModel entities in bulk.
type DefaultModelCreateBulk struct {
	config
	err      error
	builders []*DefaultModelCreate
}

// Save creates the DefaultModel entities in the database.
func (dmcb *DefaultModelCreateBulk) Save(ctx context.Context) ([]*DefaultModel, error) {
	if dmcb.err != nil {
		return nil, dmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dmcb.builders))
	nodes := make([]*DefaultModel, len(dmcb.builders))
	mutators := make([]Mutator, len(dmcb.builders))
	for i := range dmcb.builders {
		func(i int, root context.Context) {
			builder := dmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DefaultModelMutation)
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
					_, err = mutators[i+1].Mutate(root, dmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dmcb *DefaultModelCreateBulk) SaveX(ctx context.Context) []*DefaultModel {
	v, err := dmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dmcb *DefaultModelCreateBulk) Exec(ctx context.Context) error {
	_, err := dmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmcb *DefaultModelCreateBulk) ExecX(ctx context.Context) {
	if err := dmcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khang00/health/ent/bfpdatapoint"
	"github.com/khang00/health/ent/user"
)

// BFPDataPointCreate is the builder for creating a BFPDataPoint entity.
type BFPDataPointCreate struct {
	config
	mutation *BFPDataPointMutation
	hooks    []Hook
}

// SetFatPercentage sets the "fat_percentage" field.
func (bdpc *BFPDataPointCreate) SetFatPercentage(f float64) *BFPDataPointCreate {
	bdpc.mutation.SetFatPercentage(f)
	return bdpc
}

// SetTotalWeight sets the "total_weight" field.
func (bdpc *BFPDataPointCreate) SetTotalWeight(i int) *BFPDataPointCreate {
	bdpc.mutation.SetTotalWeight(i)
	return bdpc
}

// SetCreatedAt sets the "created_at" field.
func (bdpc *BFPDataPointCreate) SetCreatedAt(t time.Time) *BFPDataPointCreate {
	bdpc.mutation.SetCreatedAt(t)
	return bdpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bdpc *BFPDataPointCreate) SetNillableCreatedAt(t *time.Time) *BFPDataPointCreate {
	if t != nil {
		bdpc.SetCreatedAt(*t)
	}
	return bdpc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (bdpc *BFPDataPointCreate) SetUserID(id int) *BFPDataPointCreate {
	bdpc.mutation.SetUserID(id)
	return bdpc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (bdpc *BFPDataPointCreate) SetNillableUserID(id *int) *BFPDataPointCreate {
	if id != nil {
		bdpc = bdpc.SetUserID(*id)
	}
	return bdpc
}

// SetUser sets the "user" edge to the User entity.
func (bdpc *BFPDataPointCreate) SetUser(u *User) *BFPDataPointCreate {
	return bdpc.SetUserID(u.ID)
}

// Mutation returns the BFPDataPointMutation object of the builder.
func (bdpc *BFPDataPointCreate) Mutation() *BFPDataPointMutation {
	return bdpc.mutation
}

// Save creates the BFPDataPoint in the database.
func (bdpc *BFPDataPointCreate) Save(ctx context.Context) (*BFPDataPoint, error) {
	bdpc.defaults()
	return withHooks[*BFPDataPoint, BFPDataPointMutation](ctx, bdpc.sqlSave, bdpc.mutation, bdpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bdpc *BFPDataPointCreate) SaveX(ctx context.Context) *BFPDataPoint {
	v, err := bdpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bdpc *BFPDataPointCreate) Exec(ctx context.Context) error {
	_, err := bdpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdpc *BFPDataPointCreate) ExecX(ctx context.Context) {
	if err := bdpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bdpc *BFPDataPointCreate) defaults() {
	if _, ok := bdpc.mutation.CreatedAt(); !ok {
		v := bfpdatapoint.DefaultCreatedAt
		bdpc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bdpc *BFPDataPointCreate) check() error {
	if _, ok := bdpc.mutation.FatPercentage(); !ok {
		return &ValidationError{Name: "fat_percentage", err: errors.New(`ent: missing required field "BFPDataPoint.fat_percentage"`)}
	}
	if v, ok := bdpc.mutation.FatPercentage(); ok {
		if err := bfpdatapoint.FatPercentageValidator(v); err != nil {
			return &ValidationError{Name: "fat_percentage", err: fmt.Errorf(`ent: validator failed for field "BFPDataPoint.fat_percentage": %w`, err)}
		}
	}
	if _, ok := bdpc.mutation.TotalWeight(); !ok {
		return &ValidationError{Name: "total_weight", err: errors.New(`ent: missing required field "BFPDataPoint.total_weight"`)}
	}
	if v, ok := bdpc.mutation.TotalWeight(); ok {
		if err := bfpdatapoint.TotalWeightValidator(v); err != nil {
			return &ValidationError{Name: "total_weight", err: fmt.Errorf(`ent: validator failed for field "BFPDataPoint.total_weight": %w`, err)}
		}
	}
	if _, ok := bdpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "BFPDataPoint.created_at"`)}
	}
	return nil
}

func (bdpc *BFPDataPointCreate) sqlSave(ctx context.Context) (*BFPDataPoint, error) {
	if err := bdpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bdpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bdpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bdpc.mutation.id = &_node.ID
	bdpc.mutation.done = true
	return _node, nil
}

func (bdpc *BFPDataPointCreate) createSpec() (*BFPDataPoint, *sqlgraph.CreateSpec) {
	var (
		_node = &BFPDataPoint{config: bdpc.config}
		_spec = sqlgraph.NewCreateSpec(bfpdatapoint.Table, sqlgraph.NewFieldSpec(bfpdatapoint.FieldID, field.TypeInt))
	)
	if value, ok := bdpc.mutation.FatPercentage(); ok {
		_spec.SetField(bfpdatapoint.FieldFatPercentage, field.TypeFloat64, value)
		_node.FatPercentage = value
	}
	if value, ok := bdpc.mutation.TotalWeight(); ok {
		_spec.SetField(bfpdatapoint.FieldTotalWeight, field.TypeInt, value)
		_node.TotalWeight = value
	}
	if value, ok := bdpc.mutation.CreatedAt(); ok {
		_spec.SetField(bfpdatapoint.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := bdpc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bfpdatapoint.UserTable,
			Columns: []string{bfpdatapoint.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_bf_ps = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BFPDataPointCreateBulk is the builder for creating many BFPDataPoint entities in bulk.
type BFPDataPointCreateBulk struct {
	config
	builders []*BFPDataPointCreate
}

// Save creates the BFPDataPoint entities in the database.
func (bdpcb *BFPDataPointCreateBulk) Save(ctx context.Context) ([]*BFPDataPoint, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bdpcb.builders))
	nodes := make([]*BFPDataPoint, len(bdpcb.builders))
	mutators := make([]Mutator, len(bdpcb.builders))
	for i := range bdpcb.builders {
		func(i int, root context.Context) {
			builder := bdpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BFPDataPointMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bdpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bdpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, bdpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bdpcb *BFPDataPointCreateBulk) SaveX(ctx context.Context) []*BFPDataPoint {
	v, err := bdpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bdpcb *BFPDataPointCreateBulk) Exec(ctx context.Context) error {
	_, err := bdpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdpcb *BFPDataPointCreateBulk) ExecX(ctx context.Context) {
	if err := bdpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
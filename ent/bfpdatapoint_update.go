// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khang00/health/ent/bfpdatapoint"
	"github.com/khang00/health/ent/predicate"
	"github.com/khang00/health/ent/user"
)

// BFPDataPointUpdate is the builder for updating BFPDataPoint entities.
type BFPDataPointUpdate struct {
	config
	hooks    []Hook
	mutation *BFPDataPointMutation
}

// Where appends a list predicates to the BFPDataPointUpdate builder.
func (bdpu *BFPDataPointUpdate) Where(ps ...predicate.BFPDataPoint) *BFPDataPointUpdate {
	bdpu.mutation.Where(ps...)
	return bdpu
}

// SetFatPercentage sets the "fat_percentage" field.
func (bdpu *BFPDataPointUpdate) SetFatPercentage(f float64) *BFPDataPointUpdate {
	bdpu.mutation.ResetFatPercentage()
	bdpu.mutation.SetFatPercentage(f)
	return bdpu
}

// AddFatPercentage adds f to the "fat_percentage" field.
func (bdpu *BFPDataPointUpdate) AddFatPercentage(f float64) *BFPDataPointUpdate {
	bdpu.mutation.AddFatPercentage(f)
	return bdpu
}

// SetTotalWeight sets the "total_weight" field.
func (bdpu *BFPDataPointUpdate) SetTotalWeight(i int) *BFPDataPointUpdate {
	bdpu.mutation.ResetTotalWeight()
	bdpu.mutation.SetTotalWeight(i)
	return bdpu
}

// AddTotalWeight adds i to the "total_weight" field.
func (bdpu *BFPDataPointUpdate) AddTotalWeight(i int) *BFPDataPointUpdate {
	bdpu.mutation.AddTotalWeight(i)
	return bdpu
}

// SetCreatedAt sets the "created_at" field.
func (bdpu *BFPDataPointUpdate) SetCreatedAt(t time.Time) *BFPDataPointUpdate {
	bdpu.mutation.SetCreatedAt(t)
	return bdpu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bdpu *BFPDataPointUpdate) SetNillableCreatedAt(t *time.Time) *BFPDataPointUpdate {
	if t != nil {
		bdpu.SetCreatedAt(*t)
	}
	return bdpu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (bdpu *BFPDataPointUpdate) SetUserID(id int) *BFPDataPointUpdate {
	bdpu.mutation.SetUserID(id)
	return bdpu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (bdpu *BFPDataPointUpdate) SetNillableUserID(id *int) *BFPDataPointUpdate {
	if id != nil {
		bdpu = bdpu.SetUserID(*id)
	}
	return bdpu
}

// SetUser sets the "user" edge to the User entity.
func (bdpu *BFPDataPointUpdate) SetUser(u *User) *BFPDataPointUpdate {
	return bdpu.SetUserID(u.ID)
}

// Mutation returns the BFPDataPointMutation object of the builder.
func (bdpu *BFPDataPointUpdate) Mutation() *BFPDataPointMutation {
	return bdpu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (bdpu *BFPDataPointUpdate) ClearUser() *BFPDataPointUpdate {
	bdpu.mutation.ClearUser()
	return bdpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bdpu *BFPDataPointUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, BFPDataPointMutation](ctx, bdpu.sqlSave, bdpu.mutation, bdpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bdpu *BFPDataPointUpdate) SaveX(ctx context.Context) int {
	affected, err := bdpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bdpu *BFPDataPointUpdate) Exec(ctx context.Context) error {
	_, err := bdpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdpu *BFPDataPointUpdate) ExecX(ctx context.Context) {
	if err := bdpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bdpu *BFPDataPointUpdate) check() error {
	if v, ok := bdpu.mutation.FatPercentage(); ok {
		if err := bfpdatapoint.FatPercentageValidator(v); err != nil {
			return &ValidationError{Name: "fat_percentage", err: fmt.Errorf(`ent: validator failed for field "BFPDataPoint.fat_percentage": %w`, err)}
		}
	}
	if v, ok := bdpu.mutation.TotalWeight(); ok {
		if err := bfpdatapoint.TotalWeightValidator(v); err != nil {
			return &ValidationError{Name: "total_weight", err: fmt.Errorf(`ent: validator failed for field "BFPDataPoint.total_weight": %w`, err)}
		}
	}
	return nil
}

func (bdpu *BFPDataPointUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bdpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(bfpdatapoint.Table, bfpdatapoint.Columns, sqlgraph.NewFieldSpec(bfpdatapoint.FieldID, field.TypeInt))
	if ps := bdpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bdpu.mutation.FatPercentage(); ok {
		_spec.SetField(bfpdatapoint.FieldFatPercentage, field.TypeFloat64, value)
	}
	if value, ok := bdpu.mutation.AddedFatPercentage(); ok {
		_spec.AddField(bfpdatapoint.FieldFatPercentage, field.TypeFloat64, value)
	}
	if value, ok := bdpu.mutation.TotalWeight(); ok {
		_spec.SetField(bfpdatapoint.FieldTotalWeight, field.TypeInt, value)
	}
	if value, ok := bdpu.mutation.AddedTotalWeight(); ok {
		_spec.AddField(bfpdatapoint.FieldTotalWeight, field.TypeInt, value)
	}
	if value, ok := bdpu.mutation.CreatedAt(); ok {
		_spec.SetField(bfpdatapoint.FieldCreatedAt, field.TypeTime, value)
	}
	if bdpu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bdpu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bdpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bfpdatapoint.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bdpu.mutation.done = true
	return n, nil
}

// BFPDataPointUpdateOne is the builder for updating a single BFPDataPoint entity.
type BFPDataPointUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BFPDataPointMutation
}

// SetFatPercentage sets the "fat_percentage" field.
func (bdpuo *BFPDataPointUpdateOne) SetFatPercentage(f float64) *BFPDataPointUpdateOne {
	bdpuo.mutation.ResetFatPercentage()
	bdpuo.mutation.SetFatPercentage(f)
	return bdpuo
}

// AddFatPercentage adds f to the "fat_percentage" field.
func (bdpuo *BFPDataPointUpdateOne) AddFatPercentage(f float64) *BFPDataPointUpdateOne {
	bdpuo.mutation.AddFatPercentage(f)
	return bdpuo
}

// SetTotalWeight sets the "total_weight" field.
func (bdpuo *BFPDataPointUpdateOne) SetTotalWeight(i int) *BFPDataPointUpdateOne {
	bdpuo.mutation.ResetTotalWeight()
	bdpuo.mutation.SetTotalWeight(i)
	return bdpuo
}

// AddTotalWeight adds i to the "total_weight" field.
func (bdpuo *BFPDataPointUpdateOne) AddTotalWeight(i int) *BFPDataPointUpdateOne {
	bdpuo.mutation.AddTotalWeight(i)
	return bdpuo
}

// SetCreatedAt sets the "created_at" field.
func (bdpuo *BFPDataPointUpdateOne) SetCreatedAt(t time.Time) *BFPDataPointUpdateOne {
	bdpuo.mutation.SetCreatedAt(t)
	return bdpuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bdpuo *BFPDataPointUpdateOne) SetNillableCreatedAt(t *time.Time) *BFPDataPointUpdateOne {
	if t != nil {
		bdpuo.SetCreatedAt(*t)
	}
	return bdpuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (bdpuo *BFPDataPointUpdateOne) SetUserID(id int) *BFPDataPointUpdateOne {
	bdpuo.mutation.SetUserID(id)
	return bdpuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (bdpuo *BFPDataPointUpdateOne) SetNillableUserID(id *int) *BFPDataPointUpdateOne {
	if id != nil {
		bdpuo = bdpuo.SetUserID(*id)
	}
	return bdpuo
}

// SetUser sets the "user" edge to the User entity.
func (bdpuo *BFPDataPointUpdateOne) SetUser(u *User) *BFPDataPointUpdateOne {
	return bdpuo.SetUserID(u.ID)
}

// Mutation returns the BFPDataPointMutation object of the builder.
func (bdpuo *BFPDataPointUpdateOne) Mutation() *BFPDataPointMutation {
	return bdpuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (bdpuo *BFPDataPointUpdateOne) ClearUser() *BFPDataPointUpdateOne {
	bdpuo.mutation.ClearUser()
	return bdpuo
}

// Where appends a list predicates to the BFPDataPointUpdate builder.
func (bdpuo *BFPDataPointUpdateOne) Where(ps ...predicate.BFPDataPoint) *BFPDataPointUpdateOne {
	bdpuo.mutation.Where(ps...)
	return bdpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bdpuo *BFPDataPointUpdateOne) Select(field string, fields ...string) *BFPDataPointUpdateOne {
	bdpuo.fields = append([]string{field}, fields...)
	return bdpuo
}

// Save executes the query and returns the updated BFPDataPoint entity.
func (bdpuo *BFPDataPointUpdateOne) Save(ctx context.Context) (*BFPDataPoint, error) {
	return withHooks[*BFPDataPoint, BFPDataPointMutation](ctx, bdpuo.sqlSave, bdpuo.mutation, bdpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bdpuo *BFPDataPointUpdateOne) SaveX(ctx context.Context) *BFPDataPoint {
	node, err := bdpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bdpuo *BFPDataPointUpdateOne) Exec(ctx context.Context) error {
	_, err := bdpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdpuo *BFPDataPointUpdateOne) ExecX(ctx context.Context) {
	if err := bdpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bdpuo *BFPDataPointUpdateOne) check() error {
	if v, ok := bdpuo.mutation.FatPercentage(); ok {
		if err := bfpdatapoint.FatPercentageValidator(v); err != nil {
			return &ValidationError{Name: "fat_percentage", err: fmt.Errorf(`ent: validator failed for field "BFPDataPoint.fat_percentage": %w`, err)}
		}
	}
	if v, ok := bdpuo.mutation.TotalWeight(); ok {
		if err := bfpdatapoint.TotalWeightValidator(v); err != nil {
			return &ValidationError{Name: "total_weight", err: fmt.Errorf(`ent: validator failed for field "BFPDataPoint.total_weight": %w`, err)}
		}
	}
	return nil
}

func (bdpuo *BFPDataPointUpdateOne) sqlSave(ctx context.Context) (_node *BFPDataPoint, err error) {
	if err := bdpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(bfpdatapoint.Table, bfpdatapoint.Columns, sqlgraph.NewFieldSpec(bfpdatapoint.FieldID, field.TypeInt))
	id, ok := bdpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BFPDataPoint.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bdpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bfpdatapoint.FieldID)
		for _, f := range fields {
			if !bfpdatapoint.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bfpdatapoint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bdpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bdpuo.mutation.FatPercentage(); ok {
		_spec.SetField(bfpdatapoint.FieldFatPercentage, field.TypeFloat64, value)
	}
	if value, ok := bdpuo.mutation.AddedFatPercentage(); ok {
		_spec.AddField(bfpdatapoint.FieldFatPercentage, field.TypeFloat64, value)
	}
	if value, ok := bdpuo.mutation.TotalWeight(); ok {
		_spec.SetField(bfpdatapoint.FieldTotalWeight, field.TypeInt, value)
	}
	if value, ok := bdpuo.mutation.AddedTotalWeight(); ok {
		_spec.AddField(bfpdatapoint.FieldTotalWeight, field.TypeInt, value)
	}
	if value, ok := bdpuo.mutation.CreatedAt(); ok {
		_spec.SetField(bfpdatapoint.FieldCreatedAt, field.TypeTime, value)
	}
	if bdpuo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bdpuo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BFPDataPoint{config: bdpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bdpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bfpdatapoint.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bdpuo.mutation.done = true
	return _node, nil
}

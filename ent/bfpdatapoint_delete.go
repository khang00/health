// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khang00/health/ent/bfpdatapoint"
	"github.com/khang00/health/ent/predicate"
)

// BFPDataPointDelete is the builder for deleting a BFPDataPoint entity.
type BFPDataPointDelete struct {
	config
	hooks    []Hook
	mutation *BFPDataPointMutation
}

// Where appends a list predicates to the BFPDataPointDelete builder.
func (bdpd *BFPDataPointDelete) Where(ps ...predicate.BFPDataPoint) *BFPDataPointDelete {
	bdpd.mutation.Where(ps...)
	return bdpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bdpd *BFPDataPointDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, BFPDataPointMutation](ctx, bdpd.sqlExec, bdpd.mutation, bdpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bdpd *BFPDataPointDelete) ExecX(ctx context.Context) int {
	n, err := bdpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bdpd *BFPDataPointDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(bfpdatapoint.Table, sqlgraph.NewFieldSpec(bfpdatapoint.FieldID, field.TypeInt))
	if ps := bdpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bdpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bdpd.mutation.done = true
	return affected, err
}

// BFPDataPointDeleteOne is the builder for deleting a single BFPDataPoint entity.
type BFPDataPointDeleteOne struct {
	bdpd *BFPDataPointDelete
}

// Where appends a list predicates to the BFPDataPointDelete builder.
func (bdpdo *BFPDataPointDeleteOne) Where(ps ...predicate.BFPDataPoint) *BFPDataPointDeleteOne {
	bdpdo.bdpd.mutation.Where(ps...)
	return bdpdo
}

// Exec executes the deletion query.
func (bdpdo *BFPDataPointDeleteOne) Exec(ctx context.Context) error {
	n, err := bdpdo.bdpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bfpdatapoint.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdpdo *BFPDataPointDeleteOne) ExecX(ctx context.Context) {
	if err := bdpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
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
	"github.com/khang00/health/ent/achievement"
	"github.com/khang00/health/ent/predicate"
	"github.com/khang00/health/ent/user"
)

// AchievementUpdate is the builder for updating Achievement entities.
type AchievementUpdate struct {
	config
	hooks    []Hook
	mutation *AchievementMutation
}

// Where appends a list predicates to the AchievementUpdate builder.
func (au *AchievementUpdate) Where(ps ...predicate.Achievement) *AchievementUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetStatus sets the "status" field.
func (au *AchievementUpdate) SetStatus(s string) *AchievementUpdate {
	au.mutation.SetStatus(s)
	return au
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (au *AchievementUpdate) SetNillableStatus(s *string) *AchievementUpdate {
	if s != nil {
		au.SetStatus(*s)
	}
	return au
}

// SetThumbnailImgURL sets the "thumbnail_img_url" field.
func (au *AchievementUpdate) SetThumbnailImgURL(s string) *AchievementUpdate {
	au.mutation.SetThumbnailImgURL(s)
	return au
}

// SetNillableThumbnailImgURL sets the "thumbnail_img_url" field if the given value is not nil.
func (au *AchievementUpdate) SetNillableThumbnailImgURL(s *string) *AchievementUpdate {
	if s != nil {
		au.SetThumbnailImgURL(*s)
	}
	return au
}

// SetBfpGoal sets the "bfp_goal" field.
func (au *AchievementUpdate) SetBfpGoal(f float64) *AchievementUpdate {
	au.mutation.ResetBfpGoal()
	au.mutation.SetBfpGoal(f)
	return au
}

// AddBfpGoal adds f to the "bfp_goal" field.
func (au *AchievementUpdate) AddBfpGoal(f float64) *AchievementUpdate {
	au.mutation.AddBfpGoal(f)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AchievementUpdate) SetCreatedAt(t time.Time) *AchievementUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AchievementUpdate) SetNillableCreatedAt(t *time.Time) *AchievementUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUserID sets the "user" edge to the User entity by ID.
func (au *AchievementUpdate) SetUserID(id int) *AchievementUpdate {
	au.mutation.SetUserID(id)
	return au
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (au *AchievementUpdate) SetNillableUserID(id *int) *AchievementUpdate {
	if id != nil {
		au = au.SetUserID(*id)
	}
	return au
}

// SetUser sets the "user" edge to the User entity.
func (au *AchievementUpdate) SetUser(u *User) *AchievementUpdate {
	return au.SetUserID(u.ID)
}

// Mutation returns the AchievementMutation object of the builder.
func (au *AchievementUpdate) Mutation() *AchievementMutation {
	return au.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (au *AchievementUpdate) ClearUser() *AchievementUpdate {
	au.mutation.ClearUser()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AchievementUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AchievementMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AchievementUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AchievementUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AchievementUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AchievementUpdate) check() error {
	if v, ok := au.mutation.BfpGoal(); ok {
		if err := achievement.BfpGoalValidator(v); err != nil {
			return &ValidationError{Name: "bfp_goal", err: fmt.Errorf(`ent: validator failed for field "Achievement.bfp_goal": %w`, err)}
		}
	}
	return nil
}

func (au *AchievementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(achievement.Table, achievement.Columns, sqlgraph.NewFieldSpec(achievement.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.SetField(achievement.FieldStatus, field.TypeString, value)
	}
	if value, ok := au.mutation.ThumbnailImgURL(); ok {
		_spec.SetField(achievement.FieldThumbnailImgURL, field.TypeString, value)
	}
	if value, ok := au.mutation.BfpGoal(); ok {
		_spec.SetField(achievement.FieldBfpGoal, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.AddedBfpGoal(); ok {
		_spec.AddField(achievement.FieldBfpGoal, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(achievement.FieldCreatedAt, field.TypeTime, value)
	}
	if au.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   achievement.UserTable,
			Columns: []string{achievement.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   achievement.UserTable,
			Columns: []string{achievement.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{achievement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AchievementUpdateOne is the builder for updating a single Achievement entity.
type AchievementUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AchievementMutation
}

// SetStatus sets the "status" field.
func (auo *AchievementUpdateOne) SetStatus(s string) *AchievementUpdateOne {
	auo.mutation.SetStatus(s)
	return auo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (auo *AchievementUpdateOne) SetNillableStatus(s *string) *AchievementUpdateOne {
	if s != nil {
		auo.SetStatus(*s)
	}
	return auo
}

// SetThumbnailImgURL sets the "thumbnail_img_url" field.
func (auo *AchievementUpdateOne) SetThumbnailImgURL(s string) *AchievementUpdateOne {
	auo.mutation.SetThumbnailImgURL(s)
	return auo
}

// SetNillableThumbnailImgURL sets the "thumbnail_img_url" field if the given value is not nil.
func (auo *AchievementUpdateOne) SetNillableThumbnailImgURL(s *string) *AchievementUpdateOne {
	if s != nil {
		auo.SetThumbnailImgURL(*s)
	}
	return auo
}

// SetBfpGoal sets the "bfp_goal" field.
func (auo *AchievementUpdateOne) SetBfpGoal(f float64) *AchievementUpdateOne {
	auo.mutation.ResetBfpGoal()
	auo.mutation.SetBfpGoal(f)
	return auo
}

// AddBfpGoal adds f to the "bfp_goal" field.
func (auo *AchievementUpdateOne) AddBfpGoal(f float64) *AchievementUpdateOne {
	auo.mutation.AddBfpGoal(f)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AchievementUpdateOne) SetCreatedAt(t time.Time) *AchievementUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AchievementUpdateOne) SetNillableCreatedAt(t *time.Time) *AchievementUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (auo *AchievementUpdateOne) SetUserID(id int) *AchievementUpdateOne {
	auo.mutation.SetUserID(id)
	return auo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (auo *AchievementUpdateOne) SetNillableUserID(id *int) *AchievementUpdateOne {
	if id != nil {
		auo = auo.SetUserID(*id)
	}
	return auo
}

// SetUser sets the "user" edge to the User entity.
func (auo *AchievementUpdateOne) SetUser(u *User) *AchievementUpdateOne {
	return auo.SetUserID(u.ID)
}

// Mutation returns the AchievementMutation object of the builder.
func (auo *AchievementUpdateOne) Mutation() *AchievementMutation {
	return auo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (auo *AchievementUpdateOne) ClearUser() *AchievementUpdateOne {
	auo.mutation.ClearUser()
	return auo
}

// Where appends a list predicates to the AchievementUpdate builder.
func (auo *AchievementUpdateOne) Where(ps ...predicate.Achievement) *AchievementUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AchievementUpdateOne) Select(field string, fields ...string) *AchievementUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Achievement entity.
func (auo *AchievementUpdateOne) Save(ctx context.Context) (*Achievement, error) {
	return withHooks[*Achievement, AchievementMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AchievementUpdateOne) SaveX(ctx context.Context) *Achievement {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AchievementUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AchievementUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AchievementUpdateOne) check() error {
	if v, ok := auo.mutation.BfpGoal(); ok {
		if err := achievement.BfpGoalValidator(v); err != nil {
			return &ValidationError{Name: "bfp_goal", err: fmt.Errorf(`ent: validator failed for field "Achievement.bfp_goal": %w`, err)}
		}
	}
	return nil
}

func (auo *AchievementUpdateOne) sqlSave(ctx context.Context) (_node *Achievement, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(achievement.Table, achievement.Columns, sqlgraph.NewFieldSpec(achievement.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Achievement.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, achievement.FieldID)
		for _, f := range fields {
			if !achievement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != achievement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Status(); ok {
		_spec.SetField(achievement.FieldStatus, field.TypeString, value)
	}
	if value, ok := auo.mutation.ThumbnailImgURL(); ok {
		_spec.SetField(achievement.FieldThumbnailImgURL, field.TypeString, value)
	}
	if value, ok := auo.mutation.BfpGoal(); ok {
		_spec.SetField(achievement.FieldBfpGoal, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.AddedBfpGoal(); ok {
		_spec.AddField(achievement.FieldBfpGoal, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(achievement.FieldCreatedAt, field.TypeTime, value)
	}
	if auo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   achievement.UserTable,
			Columns: []string{achievement.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   achievement.UserTable,
			Columns: []string{achievement.UserColumn},
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
	_node = &Achievement{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{achievement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}

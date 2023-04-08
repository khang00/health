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
	"github.com/khang00/health/ent/article"
	"github.com/khang00/health/ent/predicate"
	"github.com/khang00/health/ent/tag"
)

// ArticleUpdate is the builder for updating Article entities.
type ArticleUpdate struct {
	config
	hooks    []Hook
	mutation *ArticleMutation
}

// Where appends a list predicates to the ArticleUpdate builder.
func (au *ArticleUpdate) Where(ps ...predicate.Article) *ArticleUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetThumbnailImgURL sets the "thumbnail_img_url" field.
func (au *ArticleUpdate) SetThumbnailImgURL(s string) *ArticleUpdate {
	au.mutation.SetThumbnailImgURL(s)
	return au
}

// SetNillableThumbnailImgURL sets the "thumbnail_img_url" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableThumbnailImgURL(s *string) *ArticleUpdate {
	if s != nil {
		au.SetThumbnailImgURL(*s)
	}
	return au
}

// SetTitle sets the "title" field.
func (au *ArticleUpdate) SetTitle(s string) *ArticleUpdate {
	au.mutation.SetTitle(s)
	return au
}

// SetContent sets the "content" field.
func (au *ArticleUpdate) SetContent(s string) *ArticleUpdate {
	au.mutation.SetContent(s)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *ArticleUpdate) SetCreatedAt(t time.Time) *ArticleUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableCreatedAt(t *time.Time) *ArticleUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (au *ArticleUpdate) AddTagIDs(ids ...int) *ArticleUpdate {
	au.mutation.AddTagIDs(ids...)
	return au
}

// AddTags adds the "tags" edges to the Tag entity.
func (au *ArticleUpdate) AddTags(t ...*Tag) *ArticleUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddTagIDs(ids...)
}

// Mutation returns the ArticleMutation object of the builder.
func (au *ArticleUpdate) Mutation() *ArticleMutation {
	return au.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (au *ArticleUpdate) ClearTags() *ArticleUpdate {
	au.mutation.ClearTags()
	return au
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (au *ArticleUpdate) RemoveTagIDs(ids ...int) *ArticleUpdate {
	au.mutation.RemoveTagIDs(ids...)
	return au
}

// RemoveTags removes "tags" edges to Tag entities.
func (au *ArticleUpdate) RemoveTags(t ...*Tag) *ArticleUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArticleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ArticleMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArticleUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArticleUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArticleUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *ArticleUpdate) check() error {
	if v, ok := au.mutation.Title(); ok {
		if err := article.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Article.title": %w`, err)}
		}
	}
	return nil
}

func (au *ArticleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(article.Table, article.Columns, sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.ThumbnailImgURL(); ok {
		_spec.SetField(article.FieldThumbnailImgURL, field.TypeString, value)
	}
	if value, ok := au.mutation.Title(); ok {
		_spec.SetField(article.FieldTitle, field.TypeString, value)
	}
	if value, ok := au.mutation.Content(); ok {
		_spec.SetField(article.FieldContent, field.TypeString, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(article.FieldCreatedAt, field.TypeTime, value)
	}
	if au.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   article.TagsTable,
			Columns: []string{article.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedTagsIDs(); len(nodes) > 0 && !au.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   article.TagsTable,
			Columns: []string{article.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   article.TagsTable,
			Columns: []string{article.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ArticleUpdateOne is the builder for updating a single Article entity.
type ArticleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArticleMutation
}

// SetThumbnailImgURL sets the "thumbnail_img_url" field.
func (auo *ArticleUpdateOne) SetThumbnailImgURL(s string) *ArticleUpdateOne {
	auo.mutation.SetThumbnailImgURL(s)
	return auo
}

// SetNillableThumbnailImgURL sets the "thumbnail_img_url" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableThumbnailImgURL(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetThumbnailImgURL(*s)
	}
	return auo
}

// SetTitle sets the "title" field.
func (auo *ArticleUpdateOne) SetTitle(s string) *ArticleUpdateOne {
	auo.mutation.SetTitle(s)
	return auo
}

// SetContent sets the "content" field.
func (auo *ArticleUpdateOne) SetContent(s string) *ArticleUpdateOne {
	auo.mutation.SetContent(s)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *ArticleUpdateOne) SetCreatedAt(t time.Time) *ArticleUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableCreatedAt(t *time.Time) *ArticleUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (auo *ArticleUpdateOne) AddTagIDs(ids ...int) *ArticleUpdateOne {
	auo.mutation.AddTagIDs(ids...)
	return auo
}

// AddTags adds the "tags" edges to the Tag entity.
func (auo *ArticleUpdateOne) AddTags(t ...*Tag) *ArticleUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddTagIDs(ids...)
}

// Mutation returns the ArticleMutation object of the builder.
func (auo *ArticleUpdateOne) Mutation() *ArticleMutation {
	return auo.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (auo *ArticleUpdateOne) ClearTags() *ArticleUpdateOne {
	auo.mutation.ClearTags()
	return auo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (auo *ArticleUpdateOne) RemoveTagIDs(ids ...int) *ArticleUpdateOne {
	auo.mutation.RemoveTagIDs(ids...)
	return auo
}

// RemoveTags removes "tags" edges to Tag entities.
func (auo *ArticleUpdateOne) RemoveTags(t ...*Tag) *ArticleUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveTagIDs(ids...)
}

// Where appends a list predicates to the ArticleUpdate builder.
func (auo *ArticleUpdateOne) Where(ps ...predicate.Article) *ArticleUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArticleUpdateOne) Select(field string, fields ...string) *ArticleUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Article entity.
func (auo *ArticleUpdateOne) Save(ctx context.Context) (*Article, error) {
	return withHooks[*Article, ArticleMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArticleUpdateOne) SaveX(ctx context.Context) *Article {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArticleUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArticleUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *ArticleUpdateOne) check() error {
	if v, ok := auo.mutation.Title(); ok {
		if err := article.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Article.title": %w`, err)}
		}
	}
	return nil
}

func (auo *ArticleUpdateOne) sqlSave(ctx context.Context) (_node *Article, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(article.Table, article.Columns, sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Article.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, article.FieldID)
		for _, f := range fields {
			if !article.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != article.FieldID {
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
	if value, ok := auo.mutation.ThumbnailImgURL(); ok {
		_spec.SetField(article.FieldThumbnailImgURL, field.TypeString, value)
	}
	if value, ok := auo.mutation.Title(); ok {
		_spec.SetField(article.FieldTitle, field.TypeString, value)
	}
	if value, ok := auo.mutation.Content(); ok {
		_spec.SetField(article.FieldContent, field.TypeString, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(article.FieldCreatedAt, field.TypeTime, value)
	}
	if auo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   article.TagsTable,
			Columns: []string{article.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !auo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   article.TagsTable,
			Columns: []string{article.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   article.TagsTable,
			Columns: []string{article.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Article{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}

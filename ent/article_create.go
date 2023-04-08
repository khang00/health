// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khang00/health/ent/article"
	"github.com/khang00/health/ent/tag"
)

// ArticleCreate is the builder for creating a Article entity.
type ArticleCreate struct {
	config
	mutation *ArticleMutation
	hooks    []Hook
}

// SetThumbnailImgURL sets the "thumbnail_img_url" field.
func (ac *ArticleCreate) SetThumbnailImgURL(s string) *ArticleCreate {
	ac.mutation.SetThumbnailImgURL(s)
	return ac
}

// SetNillableThumbnailImgURL sets the "thumbnail_img_url" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableThumbnailImgURL(s *string) *ArticleCreate {
	if s != nil {
		ac.SetThumbnailImgURL(*s)
	}
	return ac
}

// SetTitle sets the "title" field.
func (ac *ArticleCreate) SetTitle(s string) *ArticleCreate {
	ac.mutation.SetTitle(s)
	return ac
}

// SetContent sets the "content" field.
func (ac *ArticleCreate) SetContent(s string) *ArticleCreate {
	ac.mutation.SetContent(s)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *ArticleCreate) SetCreatedAt(t time.Time) *ArticleCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableCreatedAt(t *time.Time) *ArticleCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (ac *ArticleCreate) AddTagIDs(ids ...int) *ArticleCreate {
	ac.mutation.AddTagIDs(ids...)
	return ac
}

// AddTags adds the "tags" edges to the Tag entity.
func (ac *ArticleCreate) AddTags(t ...*Tag) *ArticleCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddTagIDs(ids...)
}

// Mutation returns the ArticleMutation object of the builder.
func (ac *ArticleCreate) Mutation() *ArticleMutation {
	return ac.mutation
}

// Save creates the Article in the database.
func (ac *ArticleCreate) Save(ctx context.Context) (*Article, error) {
	ac.defaults()
	return withHooks[*Article, ArticleMutation](ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArticleCreate) SaveX(ctx context.Context) *Article {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArticleCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArticleCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ArticleCreate) defaults() {
	if _, ok := ac.mutation.ThumbnailImgURL(); !ok {
		v := article.DefaultThumbnailImgURL
		ac.mutation.SetThumbnailImgURL(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := article.DefaultCreatedAt
		ac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArticleCreate) check() error {
	if _, ok := ac.mutation.ThumbnailImgURL(); !ok {
		return &ValidationError{Name: "thumbnail_img_url", err: errors.New(`ent: missing required field "Article.thumbnail_img_url"`)}
	}
	if _, ok := ac.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Article.title"`)}
	}
	if v, ok := ac.mutation.Title(); ok {
		if err := article.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Article.title": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Article.content"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Article.created_at"`)}
	}
	return nil
}

func (ac *ArticleCreate) sqlSave(ctx context.Context) (*Article, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ArticleCreate) createSpec() (*Article, *sqlgraph.CreateSpec) {
	var (
		_node = &Article{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(article.Table, sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt))
	)
	if value, ok := ac.mutation.ThumbnailImgURL(); ok {
		_spec.SetField(article.FieldThumbnailImgURL, field.TypeString, value)
		_node.ThumbnailImgURL = value
	}
	if value, ok := ac.mutation.Title(); ok {
		_spec.SetField(article.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := ac.mutation.Content(); ok {
		_spec.SetField(article.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(article.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ac.mutation.TagsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ArticleCreateBulk is the builder for creating many Article entities in bulk.
type ArticleCreateBulk struct {
	config
	builders []*ArticleCreate
}

// Save creates the Article entities in the database.
func (acb *ArticleCreateBulk) Save(ctx context.Context) ([]*Article, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Article, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArticleMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArticleCreateBulk) SaveX(ctx context.Context) []*Article {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArticleCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArticleCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
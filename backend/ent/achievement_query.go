// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khang00/health/backend/ent/achievement"
	"github.com/khang00/health/backend/ent/predicate"
	"github.com/khang00/health/backend/ent/user"
)

// AchievementQuery is the builder for querying Achievement entities.
type AchievementQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Achievement
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AchievementQuery builder.
func (aq *AchievementQuery) Where(ps ...predicate.Achievement) *AchievementQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AchievementQuery) Limit(limit int) *AchievementQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AchievementQuery) Offset(offset int) *AchievementQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AchievementQuery) Unique(unique bool) *AchievementQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AchievementQuery) Order(o ...OrderFunc) *AchievementQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryUser chains the current query on the "user" edge.
func (aq *AchievementQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(achievement.Table, achievement.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, achievement.UserTable, achievement.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Achievement entity from the query.
// Returns a *NotFoundError when no Achievement was found.
func (aq *AchievementQuery) First(ctx context.Context) (*Achievement, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{achievement.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AchievementQuery) FirstX(ctx context.Context) *Achievement {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Achievement ID from the query.
// Returns a *NotFoundError when no Achievement ID was found.
func (aq *AchievementQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{achievement.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AchievementQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Achievement entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Achievement entity is found.
// Returns a *NotFoundError when no Achievement entities are found.
func (aq *AchievementQuery) Only(ctx context.Context) (*Achievement, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{achievement.Label}
	default:
		return nil, &NotSingularError{achievement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AchievementQuery) OnlyX(ctx context.Context) *Achievement {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Achievement ID in the query.
// Returns a *NotSingularError when more than one Achievement ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AchievementQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{achievement.Label}
	default:
		err = &NotSingularError{achievement.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AchievementQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Achievements.
func (aq *AchievementQuery) All(ctx context.Context) ([]*Achievement, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Achievement, *AchievementQuery]()
	return withInterceptors[[]*Achievement](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AchievementQuery) AllX(ctx context.Context) []*Achievement {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Achievement IDs.
func (aq *AchievementQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(achievement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AchievementQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AchievementQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AchievementQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AchievementQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AchievementQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AchievementQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AchievementQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AchievementQuery) Clone() *AchievementQuery {
	if aq == nil {
		return nil
	}
	return &AchievementQuery{
		config:     aq.config,
		ctx:        aq.ctx.Clone(),
		order:      append([]OrderFunc{}, aq.order...),
		inters:     append([]Interceptor{}, aq.inters...),
		predicates: append([]predicate.Achievement{}, aq.predicates...),
		withUser:   aq.withUser.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AchievementQuery) WithUser(opts ...func(*UserQuery)) *AchievementQuery {
	query := (&UserClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withUser = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status string `json:"status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Achievement.Query().
//		GroupBy(achievement.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AchievementQuery) GroupBy(field string, fields ...string) *AchievementGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AchievementGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = achievement.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status string `json:"status,omitempty"`
//	}
//
//	client.Achievement.Query().
//		Select(achievement.FieldStatus).
//		Scan(ctx, &v)
func (aq *AchievementQuery) Select(fields ...string) *AchievementSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AchievementSelect{AchievementQuery: aq}
	sbuild.label = achievement.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AchievementSelect configured with the given aggregations.
func (aq *AchievementQuery) Aggregate(fns ...AggregateFunc) *AchievementSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AchievementQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !achievement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AchievementQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Achievement, error) {
	var (
		nodes       = []*Achievement{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withUser != nil,
		}
	)
	if aq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, achievement.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Achievement).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Achievement{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withUser; query != nil {
		if err := aq.loadUser(ctx, query, nodes, nil,
			func(n *Achievement, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AchievementQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Achievement, init func(*Achievement), assign func(*Achievement, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Achievement)
	for i := range nodes {
		if nodes[i].user_achievements == nil {
			continue
		}
		fk := *nodes[i].user_achievements
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_achievements" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aq *AchievementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AchievementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(achievement.Table, achievement.Columns, sqlgraph.NewFieldSpec(achievement.FieldID, field.TypeInt))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, achievement.FieldID)
		for i := range fields {
			if fields[i] != achievement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AchievementQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(achievement.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = achievement.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AchievementGroupBy is the group-by builder for Achievement entities.
type AchievementGroupBy struct {
	selector
	build *AchievementQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AchievementGroupBy) Aggregate(fns ...AggregateFunc) *AchievementGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AchievementGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AchievementQuery, *AchievementGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AchievementGroupBy) sqlScan(ctx context.Context, root *AchievementQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AchievementSelect is the builder for selecting fields of Achievement entities.
type AchievementSelect struct {
	*AchievementQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AchievementSelect) Aggregate(fns ...AggregateFunc) *AchievementSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AchievementSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AchievementQuery, *AchievementSelect](ctx, as.AchievementQuery, as, as.inters, v)
}

func (as *AchievementSelect) sqlScan(ctx context.Context, root *AchievementQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
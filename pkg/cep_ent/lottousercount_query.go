// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/lotto"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/lottousercount"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// LottoUserCountQuery is the builder for querying LottoUserCount entities.
type LottoUserCountQuery struct {
	config
	ctx        *QueryContext
	order      []lottousercount.OrderOption
	inters     []Interceptor
	predicates []predicate.LottoUserCount
	withUser   *UserQuery
	withLotto  *LottoQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LottoUserCountQuery builder.
func (lucq *LottoUserCountQuery) Where(ps ...predicate.LottoUserCount) *LottoUserCountQuery {
	lucq.predicates = append(lucq.predicates, ps...)
	return lucq
}

// Limit the number of records to be returned by this query.
func (lucq *LottoUserCountQuery) Limit(limit int) *LottoUserCountQuery {
	lucq.ctx.Limit = &limit
	return lucq
}

// Offset to start from.
func (lucq *LottoUserCountQuery) Offset(offset int) *LottoUserCountQuery {
	lucq.ctx.Offset = &offset
	return lucq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lucq *LottoUserCountQuery) Unique(unique bool) *LottoUserCountQuery {
	lucq.ctx.Unique = &unique
	return lucq
}

// Order specifies how the records should be ordered.
func (lucq *LottoUserCountQuery) Order(o ...lottousercount.OrderOption) *LottoUserCountQuery {
	lucq.order = append(lucq.order, o...)
	return lucq
}

// QueryUser chains the current query on the "user" edge.
func (lucq *LottoUserCountQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: lucq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lucq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lottousercount.Table, lottousercount.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, lottousercount.UserTable, lottousercount.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(lucq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLotto chains the current query on the "lotto" edge.
func (lucq *LottoUserCountQuery) QueryLotto() *LottoQuery {
	query := (&LottoClient{config: lucq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lucq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lottousercount.Table, lottousercount.FieldID, selector),
			sqlgraph.To(lotto.Table, lotto.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, lottousercount.LottoTable, lottousercount.LottoColumn),
		)
		fromU = sqlgraph.SetNeighbors(lucq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LottoUserCount entity from the query.
// Returns a *NotFoundError when no LottoUserCount was found.
func (lucq *LottoUserCountQuery) First(ctx context.Context) (*LottoUserCount, error) {
	nodes, err := lucq.Limit(1).All(setContextOp(ctx, lucq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{lottousercount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lucq *LottoUserCountQuery) FirstX(ctx context.Context) *LottoUserCount {
	node, err := lucq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LottoUserCount ID from the query.
// Returns a *NotFoundError when no LottoUserCount ID was found.
func (lucq *LottoUserCountQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lucq.Limit(1).IDs(setContextOp(ctx, lucq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{lottousercount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lucq *LottoUserCountQuery) FirstIDX(ctx context.Context) int64 {
	id, err := lucq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LottoUserCount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LottoUserCount entity is found.
// Returns a *NotFoundError when no LottoUserCount entities are found.
func (lucq *LottoUserCountQuery) Only(ctx context.Context) (*LottoUserCount, error) {
	nodes, err := lucq.Limit(2).All(setContextOp(ctx, lucq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{lottousercount.Label}
	default:
		return nil, &NotSingularError{lottousercount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lucq *LottoUserCountQuery) OnlyX(ctx context.Context) *LottoUserCount {
	node, err := lucq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LottoUserCount ID in the query.
// Returns a *NotSingularError when more than one LottoUserCount ID is found.
// Returns a *NotFoundError when no entities are found.
func (lucq *LottoUserCountQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lucq.Limit(2).IDs(setContextOp(ctx, lucq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{lottousercount.Label}
	default:
		err = &NotSingularError{lottousercount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lucq *LottoUserCountQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := lucq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LottoUserCounts.
func (lucq *LottoUserCountQuery) All(ctx context.Context) ([]*LottoUserCount, error) {
	ctx = setContextOp(ctx, lucq.ctx, "All")
	if err := lucq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LottoUserCount, *LottoUserCountQuery]()
	return withInterceptors[[]*LottoUserCount](ctx, lucq, qr, lucq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lucq *LottoUserCountQuery) AllX(ctx context.Context) []*LottoUserCount {
	nodes, err := lucq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LottoUserCount IDs.
func (lucq *LottoUserCountQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if lucq.ctx.Unique == nil && lucq.path != nil {
		lucq.Unique(true)
	}
	ctx = setContextOp(ctx, lucq.ctx, "IDs")
	if err = lucq.Select(lottousercount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lucq *LottoUserCountQuery) IDsX(ctx context.Context) []int64 {
	ids, err := lucq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lucq *LottoUserCountQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lucq.ctx, "Count")
	if err := lucq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lucq, querierCount[*LottoUserCountQuery](), lucq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lucq *LottoUserCountQuery) CountX(ctx context.Context) int {
	count, err := lucq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lucq *LottoUserCountQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lucq.ctx, "Exist")
	switch _, err := lucq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lucq *LottoUserCountQuery) ExistX(ctx context.Context) bool {
	exist, err := lucq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LottoUserCountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lucq *LottoUserCountQuery) Clone() *LottoUserCountQuery {
	if lucq == nil {
		return nil
	}
	return &LottoUserCountQuery{
		config:     lucq.config,
		ctx:        lucq.ctx.Clone(),
		order:      append([]lottousercount.OrderOption{}, lucq.order...),
		inters:     append([]Interceptor{}, lucq.inters...),
		predicates: append([]predicate.LottoUserCount{}, lucq.predicates...),
		withUser:   lucq.withUser.Clone(),
		withLotto:  lucq.withLotto.Clone(),
		// clone intermediate query.
		sql:  lucq.sql.Clone(),
		path: lucq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (lucq *LottoUserCountQuery) WithUser(opts ...func(*UserQuery)) *LottoUserCountQuery {
	query := (&UserClient{config: lucq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lucq.withUser = query
	return lucq
}

// WithLotto tells the query-builder to eager-load the nodes that are connected to
// the "lotto" edge. The optional arguments are used to configure the query builder of the edge.
func (lucq *LottoUserCountQuery) WithLotto(opts ...func(*LottoQuery)) *LottoUserCountQuery {
	query := (&LottoClient{config: lucq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lucq.withLotto = query
	return lucq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LottoUserCount.Query().
//		GroupBy(lottousercount.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (lucq *LottoUserCountQuery) GroupBy(field string, fields ...string) *LottoUserCountGroupBy {
	lucq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LottoUserCountGroupBy{build: lucq}
	grbuild.flds = &lucq.ctx.Fields
	grbuild.label = lottousercount.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by"`
//	}
//
//	client.LottoUserCount.Query().
//		Select(lottousercount.FieldCreatedBy).
//		Scan(ctx, &v)
func (lucq *LottoUserCountQuery) Select(fields ...string) *LottoUserCountSelect {
	lucq.ctx.Fields = append(lucq.ctx.Fields, fields...)
	sbuild := &LottoUserCountSelect{LottoUserCountQuery: lucq}
	sbuild.label = lottousercount.Label
	sbuild.flds, sbuild.scan = &lucq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LottoUserCountSelect configured with the given aggregations.
func (lucq *LottoUserCountQuery) Aggregate(fns ...AggregateFunc) *LottoUserCountSelect {
	return lucq.Select().Aggregate(fns...)
}

func (lucq *LottoUserCountQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lucq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lucq); err != nil {
				return err
			}
		}
	}
	for _, f := range lucq.ctx.Fields {
		if !lottousercount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if lucq.path != nil {
		prev, err := lucq.path(ctx)
		if err != nil {
			return err
		}
		lucq.sql = prev
	}
	return nil
}

func (lucq *LottoUserCountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LottoUserCount, error) {
	var (
		nodes       = []*LottoUserCount{}
		_spec       = lucq.querySpec()
		loadedTypes = [2]bool{
			lucq.withUser != nil,
			lucq.withLotto != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LottoUserCount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LottoUserCount{config: lucq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(lucq.modifiers) > 0 {
		_spec.Modifiers = lucq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lucq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lucq.withUser; query != nil {
		if err := lucq.loadUser(ctx, query, nodes, nil,
			func(n *LottoUserCount, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := lucq.withLotto; query != nil {
		if err := lucq.loadLotto(ctx, query, nodes, nil,
			func(n *LottoUserCount, e *Lotto) { n.Edges.Lotto = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lucq *LottoUserCountQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*LottoUserCount, init func(*LottoUserCount), assign func(*LottoUserCount, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*LottoUserCount)
	for i := range nodes {
		fk := nodes[i].UserID
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
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (lucq *LottoUserCountQuery) loadLotto(ctx context.Context, query *LottoQuery, nodes []*LottoUserCount, init func(*LottoUserCount), assign func(*LottoUserCount, *Lotto)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*LottoUserCount)
	for i := range nodes {
		fk := nodes[i].LottoID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(lotto.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "lotto_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lucq *LottoUserCountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lucq.querySpec()
	if len(lucq.modifiers) > 0 {
		_spec.Modifiers = lucq.modifiers
	}
	_spec.Node.Columns = lucq.ctx.Fields
	if len(lucq.ctx.Fields) > 0 {
		_spec.Unique = lucq.ctx.Unique != nil && *lucq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lucq.driver, _spec)
}

func (lucq *LottoUserCountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(lottousercount.Table, lottousercount.Columns, sqlgraph.NewFieldSpec(lottousercount.FieldID, field.TypeInt64))
	_spec.From = lucq.sql
	if unique := lucq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lucq.path != nil {
		_spec.Unique = true
	}
	if fields := lucq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lottousercount.FieldID)
		for i := range fields {
			if fields[i] != lottousercount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if lucq.withUser != nil {
			_spec.Node.AddColumnOnce(lottousercount.FieldUserID)
		}
		if lucq.withLotto != nil {
			_spec.Node.AddColumnOnce(lottousercount.FieldLottoID)
		}
	}
	if ps := lucq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lucq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lucq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lucq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lucq *LottoUserCountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lucq.driver.Dialect())
	t1 := builder.Table(lottousercount.Table)
	columns := lucq.ctx.Fields
	if len(columns) == 0 {
		columns = lottousercount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lucq.sql != nil {
		selector = lucq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lucq.ctx.Unique != nil && *lucq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range lucq.modifiers {
		m(selector)
	}
	for _, p := range lucq.predicates {
		p(selector)
	}
	for _, p := range lucq.order {
		p(selector)
	}
	if offset := lucq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lucq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lucq *LottoUserCountQuery) Modify(modifiers ...func(s *sql.Selector)) *LottoUserCountSelect {
	lucq.modifiers = append(lucq.modifiers, modifiers...)
	return lucq.Select()
}

// LottoUserCountGroupBy is the group-by builder for LottoUserCount entities.
type LottoUserCountGroupBy struct {
	selector
	build *LottoUserCountQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lucgb *LottoUserCountGroupBy) Aggregate(fns ...AggregateFunc) *LottoUserCountGroupBy {
	lucgb.fns = append(lucgb.fns, fns...)
	return lucgb
}

// Scan applies the selector query and scans the result into the given value.
func (lucgb *LottoUserCountGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lucgb.build.ctx, "GroupBy")
	if err := lucgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LottoUserCountQuery, *LottoUserCountGroupBy](ctx, lucgb.build, lucgb, lucgb.build.inters, v)
}

func (lucgb *LottoUserCountGroupBy) sqlScan(ctx context.Context, root *LottoUserCountQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lucgb.fns))
	for _, fn := range lucgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lucgb.flds)+len(lucgb.fns))
		for _, f := range *lucgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lucgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lucgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LottoUserCountSelect is the builder for selecting fields of LottoUserCount entities.
type LottoUserCountSelect struct {
	*LottoUserCountQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lucs *LottoUserCountSelect) Aggregate(fns ...AggregateFunc) *LottoUserCountSelect {
	lucs.fns = append(lucs.fns, fns...)
	return lucs
}

// Scan applies the selector query and scans the result into the given value.
func (lucs *LottoUserCountSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lucs.ctx, "Select")
	if err := lucs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LottoUserCountQuery, *LottoUserCountSelect](ctx, lucs.LottoUserCountQuery, lucs, lucs.inters, v)
}

func (lucs *LottoUserCountSelect) sqlScan(ctx context.Context, root *LottoUserCountQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lucs.fns))
	for _, fn := range lucs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lucs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lucs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lucs *LottoUserCountSelect) Modify(modifiers ...func(s *sql.Selector)) *LottoUserCountSelect {
	lucs.modifiers = append(lucs.modifiers, modifiers...)
	return lucs
}

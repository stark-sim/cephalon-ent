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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/lottochancerule"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// LottoChanceRuleQuery is the builder for querying LottoChanceRule entities.
type LottoChanceRuleQuery struct {
	config
	ctx        *QueryContext
	order      []lottochancerule.OrderOption
	inters     []Interceptor
	predicates []predicate.LottoChanceRule
	withLotto  *LottoQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LottoChanceRuleQuery builder.
func (lcrq *LottoChanceRuleQuery) Where(ps ...predicate.LottoChanceRule) *LottoChanceRuleQuery {
	lcrq.predicates = append(lcrq.predicates, ps...)
	return lcrq
}

// Limit the number of records to be returned by this query.
func (lcrq *LottoChanceRuleQuery) Limit(limit int) *LottoChanceRuleQuery {
	lcrq.ctx.Limit = &limit
	return lcrq
}

// Offset to start from.
func (lcrq *LottoChanceRuleQuery) Offset(offset int) *LottoChanceRuleQuery {
	lcrq.ctx.Offset = &offset
	return lcrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lcrq *LottoChanceRuleQuery) Unique(unique bool) *LottoChanceRuleQuery {
	lcrq.ctx.Unique = &unique
	return lcrq
}

// Order specifies how the records should be ordered.
func (lcrq *LottoChanceRuleQuery) Order(o ...lottochancerule.OrderOption) *LottoChanceRuleQuery {
	lcrq.order = append(lcrq.order, o...)
	return lcrq
}

// QueryLotto chains the current query on the "lotto" edge.
func (lcrq *LottoChanceRuleQuery) QueryLotto() *LottoQuery {
	query := (&LottoClient{config: lcrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lcrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lcrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lottochancerule.Table, lottochancerule.FieldID, selector),
			sqlgraph.To(lotto.Table, lotto.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, lottochancerule.LottoTable, lottochancerule.LottoColumn),
		)
		fromU = sqlgraph.SetNeighbors(lcrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LottoChanceRule entity from the query.
// Returns a *NotFoundError when no LottoChanceRule was found.
func (lcrq *LottoChanceRuleQuery) First(ctx context.Context) (*LottoChanceRule, error) {
	nodes, err := lcrq.Limit(1).All(setContextOp(ctx, lcrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{lottochancerule.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lcrq *LottoChanceRuleQuery) FirstX(ctx context.Context) *LottoChanceRule {
	node, err := lcrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LottoChanceRule ID from the query.
// Returns a *NotFoundError when no LottoChanceRule ID was found.
func (lcrq *LottoChanceRuleQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lcrq.Limit(1).IDs(setContextOp(ctx, lcrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{lottochancerule.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lcrq *LottoChanceRuleQuery) FirstIDX(ctx context.Context) int64 {
	id, err := lcrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LottoChanceRule entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LottoChanceRule entity is found.
// Returns a *NotFoundError when no LottoChanceRule entities are found.
func (lcrq *LottoChanceRuleQuery) Only(ctx context.Context) (*LottoChanceRule, error) {
	nodes, err := lcrq.Limit(2).All(setContextOp(ctx, lcrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{lottochancerule.Label}
	default:
		return nil, &NotSingularError{lottochancerule.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lcrq *LottoChanceRuleQuery) OnlyX(ctx context.Context) *LottoChanceRule {
	node, err := lcrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LottoChanceRule ID in the query.
// Returns a *NotSingularError when more than one LottoChanceRule ID is found.
// Returns a *NotFoundError when no entities are found.
func (lcrq *LottoChanceRuleQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lcrq.Limit(2).IDs(setContextOp(ctx, lcrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{lottochancerule.Label}
	default:
		err = &NotSingularError{lottochancerule.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lcrq *LottoChanceRuleQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := lcrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LottoChanceRules.
func (lcrq *LottoChanceRuleQuery) All(ctx context.Context) ([]*LottoChanceRule, error) {
	ctx = setContextOp(ctx, lcrq.ctx, "All")
	if err := lcrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LottoChanceRule, *LottoChanceRuleQuery]()
	return withInterceptors[[]*LottoChanceRule](ctx, lcrq, qr, lcrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lcrq *LottoChanceRuleQuery) AllX(ctx context.Context) []*LottoChanceRule {
	nodes, err := lcrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LottoChanceRule IDs.
func (lcrq *LottoChanceRuleQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if lcrq.ctx.Unique == nil && lcrq.path != nil {
		lcrq.Unique(true)
	}
	ctx = setContextOp(ctx, lcrq.ctx, "IDs")
	if err = lcrq.Select(lottochancerule.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lcrq *LottoChanceRuleQuery) IDsX(ctx context.Context) []int64 {
	ids, err := lcrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lcrq *LottoChanceRuleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lcrq.ctx, "Count")
	if err := lcrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lcrq, querierCount[*LottoChanceRuleQuery](), lcrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lcrq *LottoChanceRuleQuery) CountX(ctx context.Context) int {
	count, err := lcrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lcrq *LottoChanceRuleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lcrq.ctx, "Exist")
	switch _, err := lcrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lcrq *LottoChanceRuleQuery) ExistX(ctx context.Context) bool {
	exist, err := lcrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LottoChanceRuleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lcrq *LottoChanceRuleQuery) Clone() *LottoChanceRuleQuery {
	if lcrq == nil {
		return nil
	}
	return &LottoChanceRuleQuery{
		config:     lcrq.config,
		ctx:        lcrq.ctx.Clone(),
		order:      append([]lottochancerule.OrderOption{}, lcrq.order...),
		inters:     append([]Interceptor{}, lcrq.inters...),
		predicates: append([]predicate.LottoChanceRule{}, lcrq.predicates...),
		withLotto:  lcrq.withLotto.Clone(),
		// clone intermediate query.
		sql:  lcrq.sql.Clone(),
		path: lcrq.path,
	}
}

// WithLotto tells the query-builder to eager-load the nodes that are connected to
// the "lotto" edge. The optional arguments are used to configure the query builder of the edge.
func (lcrq *LottoChanceRuleQuery) WithLotto(opts ...func(*LottoQuery)) *LottoChanceRuleQuery {
	query := (&LottoClient{config: lcrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lcrq.withLotto = query
	return lcrq
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
//	client.LottoChanceRule.Query().
//		GroupBy(lottochancerule.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (lcrq *LottoChanceRuleQuery) GroupBy(field string, fields ...string) *LottoChanceRuleGroupBy {
	lcrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LottoChanceRuleGroupBy{build: lcrq}
	grbuild.flds = &lcrq.ctx.Fields
	grbuild.label = lottochancerule.Label
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
//	client.LottoChanceRule.Query().
//		Select(lottochancerule.FieldCreatedBy).
//		Scan(ctx, &v)
func (lcrq *LottoChanceRuleQuery) Select(fields ...string) *LottoChanceRuleSelect {
	lcrq.ctx.Fields = append(lcrq.ctx.Fields, fields...)
	sbuild := &LottoChanceRuleSelect{LottoChanceRuleQuery: lcrq}
	sbuild.label = lottochancerule.Label
	sbuild.flds, sbuild.scan = &lcrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LottoChanceRuleSelect configured with the given aggregations.
func (lcrq *LottoChanceRuleQuery) Aggregate(fns ...AggregateFunc) *LottoChanceRuleSelect {
	return lcrq.Select().Aggregate(fns...)
}

func (lcrq *LottoChanceRuleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lcrq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lcrq); err != nil {
				return err
			}
		}
	}
	for _, f := range lcrq.ctx.Fields {
		if !lottochancerule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if lcrq.path != nil {
		prev, err := lcrq.path(ctx)
		if err != nil {
			return err
		}
		lcrq.sql = prev
	}
	return nil
}

func (lcrq *LottoChanceRuleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LottoChanceRule, error) {
	var (
		nodes       = []*LottoChanceRule{}
		_spec       = lcrq.querySpec()
		loadedTypes = [1]bool{
			lcrq.withLotto != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LottoChanceRule).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LottoChanceRule{config: lcrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(lcrq.modifiers) > 0 {
		_spec.Modifiers = lcrq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lcrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lcrq.withLotto; query != nil {
		if err := lcrq.loadLotto(ctx, query, nodes, nil,
			func(n *LottoChanceRule, e *Lotto) { n.Edges.Lotto = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lcrq *LottoChanceRuleQuery) loadLotto(ctx context.Context, query *LottoQuery, nodes []*LottoChanceRule, init func(*LottoChanceRule), assign func(*LottoChanceRule, *Lotto)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*LottoChanceRule)
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

func (lcrq *LottoChanceRuleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lcrq.querySpec()
	if len(lcrq.modifiers) > 0 {
		_spec.Modifiers = lcrq.modifiers
	}
	_spec.Node.Columns = lcrq.ctx.Fields
	if len(lcrq.ctx.Fields) > 0 {
		_spec.Unique = lcrq.ctx.Unique != nil && *lcrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lcrq.driver, _spec)
}

func (lcrq *LottoChanceRuleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(lottochancerule.Table, lottochancerule.Columns, sqlgraph.NewFieldSpec(lottochancerule.FieldID, field.TypeInt64))
	_spec.From = lcrq.sql
	if unique := lcrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lcrq.path != nil {
		_spec.Unique = true
	}
	if fields := lcrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lottochancerule.FieldID)
		for i := range fields {
			if fields[i] != lottochancerule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if lcrq.withLotto != nil {
			_spec.Node.AddColumnOnce(lottochancerule.FieldLottoID)
		}
	}
	if ps := lcrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lcrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lcrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lcrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lcrq *LottoChanceRuleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lcrq.driver.Dialect())
	t1 := builder.Table(lottochancerule.Table)
	columns := lcrq.ctx.Fields
	if len(columns) == 0 {
		columns = lottochancerule.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lcrq.sql != nil {
		selector = lcrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lcrq.ctx.Unique != nil && *lcrq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range lcrq.modifiers {
		m(selector)
	}
	for _, p := range lcrq.predicates {
		p(selector)
	}
	for _, p := range lcrq.order {
		p(selector)
	}
	if offset := lcrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lcrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lcrq *LottoChanceRuleQuery) Modify(modifiers ...func(s *sql.Selector)) *LottoChanceRuleSelect {
	lcrq.modifiers = append(lcrq.modifiers, modifiers...)
	return lcrq.Select()
}

// LottoChanceRuleGroupBy is the group-by builder for LottoChanceRule entities.
type LottoChanceRuleGroupBy struct {
	selector
	build *LottoChanceRuleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lcrgb *LottoChanceRuleGroupBy) Aggregate(fns ...AggregateFunc) *LottoChanceRuleGroupBy {
	lcrgb.fns = append(lcrgb.fns, fns...)
	return lcrgb
}

// Scan applies the selector query and scans the result into the given value.
func (lcrgb *LottoChanceRuleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lcrgb.build.ctx, "GroupBy")
	if err := lcrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LottoChanceRuleQuery, *LottoChanceRuleGroupBy](ctx, lcrgb.build, lcrgb, lcrgb.build.inters, v)
}

func (lcrgb *LottoChanceRuleGroupBy) sqlScan(ctx context.Context, root *LottoChanceRuleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lcrgb.fns))
	for _, fn := range lcrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lcrgb.flds)+len(lcrgb.fns))
		for _, f := range *lcrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lcrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lcrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LottoChanceRuleSelect is the builder for selecting fields of LottoChanceRule entities.
type LottoChanceRuleSelect struct {
	*LottoChanceRuleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lcrs *LottoChanceRuleSelect) Aggregate(fns ...AggregateFunc) *LottoChanceRuleSelect {
	lcrs.fns = append(lcrs.fns, fns...)
	return lcrs
}

// Scan applies the selector query and scans the result into the given value.
func (lcrs *LottoChanceRuleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lcrs.ctx, "Select")
	if err := lcrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LottoChanceRuleQuery, *LottoChanceRuleSelect](ctx, lcrs.LottoChanceRuleQuery, lcrs, lcrs.inters, v)
}

func (lcrs *LottoChanceRuleSelect) sqlScan(ctx context.Context, root *LottoChanceRuleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lcrs.fns))
	for _, fn := range lcrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lcrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lcrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lcrs *LottoChanceRuleSelect) Modify(modifiers ...func(s *sql.Selector)) *LottoChanceRuleSelect {
	lcrs.modifiers = append(lcrs.modifiers, modifiers...)
	return lcrs
}

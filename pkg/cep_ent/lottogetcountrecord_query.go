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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/lottogetcountrecord"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// LottoGetCountRecordQuery is the builder for querying LottoGetCountRecord entities.
type LottoGetCountRecordQuery struct {
	config
	ctx        *QueryContext
	order      []lottogetcountrecord.OrderOption
	inters     []Interceptor
	predicates []predicate.LottoGetCountRecord
	withUser   *UserQuery
	withLotto  *LottoQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LottoGetCountRecordQuery builder.
func (lgcrq *LottoGetCountRecordQuery) Where(ps ...predicate.LottoGetCountRecord) *LottoGetCountRecordQuery {
	lgcrq.predicates = append(lgcrq.predicates, ps...)
	return lgcrq
}

// Limit the number of records to be returned by this query.
func (lgcrq *LottoGetCountRecordQuery) Limit(limit int) *LottoGetCountRecordQuery {
	lgcrq.ctx.Limit = &limit
	return lgcrq
}

// Offset to start from.
func (lgcrq *LottoGetCountRecordQuery) Offset(offset int) *LottoGetCountRecordQuery {
	lgcrq.ctx.Offset = &offset
	return lgcrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lgcrq *LottoGetCountRecordQuery) Unique(unique bool) *LottoGetCountRecordQuery {
	lgcrq.ctx.Unique = &unique
	return lgcrq
}

// Order specifies how the records should be ordered.
func (lgcrq *LottoGetCountRecordQuery) Order(o ...lottogetcountrecord.OrderOption) *LottoGetCountRecordQuery {
	lgcrq.order = append(lgcrq.order, o...)
	return lgcrq
}

// QueryUser chains the current query on the "user" edge.
func (lgcrq *LottoGetCountRecordQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: lgcrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lgcrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lgcrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lottogetcountrecord.Table, lottogetcountrecord.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, lottogetcountrecord.UserTable, lottogetcountrecord.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(lgcrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLotto chains the current query on the "lotto" edge.
func (lgcrq *LottoGetCountRecordQuery) QueryLotto() *LottoQuery {
	query := (&LottoClient{config: lgcrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lgcrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lgcrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lottogetcountrecord.Table, lottogetcountrecord.FieldID, selector),
			sqlgraph.To(lotto.Table, lotto.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, lottogetcountrecord.LottoTable, lottogetcountrecord.LottoColumn),
		)
		fromU = sqlgraph.SetNeighbors(lgcrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LottoGetCountRecord entity from the query.
// Returns a *NotFoundError when no LottoGetCountRecord was found.
func (lgcrq *LottoGetCountRecordQuery) First(ctx context.Context) (*LottoGetCountRecord, error) {
	nodes, err := lgcrq.Limit(1).All(setContextOp(ctx, lgcrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{lottogetcountrecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lgcrq *LottoGetCountRecordQuery) FirstX(ctx context.Context) *LottoGetCountRecord {
	node, err := lgcrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LottoGetCountRecord ID from the query.
// Returns a *NotFoundError when no LottoGetCountRecord ID was found.
func (lgcrq *LottoGetCountRecordQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lgcrq.Limit(1).IDs(setContextOp(ctx, lgcrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{lottogetcountrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lgcrq *LottoGetCountRecordQuery) FirstIDX(ctx context.Context) int64 {
	id, err := lgcrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LottoGetCountRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LottoGetCountRecord entity is found.
// Returns a *NotFoundError when no LottoGetCountRecord entities are found.
func (lgcrq *LottoGetCountRecordQuery) Only(ctx context.Context) (*LottoGetCountRecord, error) {
	nodes, err := lgcrq.Limit(2).All(setContextOp(ctx, lgcrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{lottogetcountrecord.Label}
	default:
		return nil, &NotSingularError{lottogetcountrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lgcrq *LottoGetCountRecordQuery) OnlyX(ctx context.Context) *LottoGetCountRecord {
	node, err := lgcrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LottoGetCountRecord ID in the query.
// Returns a *NotSingularError when more than one LottoGetCountRecord ID is found.
// Returns a *NotFoundError when no entities are found.
func (lgcrq *LottoGetCountRecordQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lgcrq.Limit(2).IDs(setContextOp(ctx, lgcrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{lottogetcountrecord.Label}
	default:
		err = &NotSingularError{lottogetcountrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lgcrq *LottoGetCountRecordQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := lgcrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LottoGetCountRecords.
func (lgcrq *LottoGetCountRecordQuery) All(ctx context.Context) ([]*LottoGetCountRecord, error) {
	ctx = setContextOp(ctx, lgcrq.ctx, "All")
	if err := lgcrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LottoGetCountRecord, *LottoGetCountRecordQuery]()
	return withInterceptors[[]*LottoGetCountRecord](ctx, lgcrq, qr, lgcrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lgcrq *LottoGetCountRecordQuery) AllX(ctx context.Context) []*LottoGetCountRecord {
	nodes, err := lgcrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LottoGetCountRecord IDs.
func (lgcrq *LottoGetCountRecordQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if lgcrq.ctx.Unique == nil && lgcrq.path != nil {
		lgcrq.Unique(true)
	}
	ctx = setContextOp(ctx, lgcrq.ctx, "IDs")
	if err = lgcrq.Select(lottogetcountrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lgcrq *LottoGetCountRecordQuery) IDsX(ctx context.Context) []int64 {
	ids, err := lgcrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lgcrq *LottoGetCountRecordQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lgcrq.ctx, "Count")
	if err := lgcrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lgcrq, querierCount[*LottoGetCountRecordQuery](), lgcrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lgcrq *LottoGetCountRecordQuery) CountX(ctx context.Context) int {
	count, err := lgcrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lgcrq *LottoGetCountRecordQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lgcrq.ctx, "Exist")
	switch _, err := lgcrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lgcrq *LottoGetCountRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := lgcrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LottoGetCountRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lgcrq *LottoGetCountRecordQuery) Clone() *LottoGetCountRecordQuery {
	if lgcrq == nil {
		return nil
	}
	return &LottoGetCountRecordQuery{
		config:     lgcrq.config,
		ctx:        lgcrq.ctx.Clone(),
		order:      append([]lottogetcountrecord.OrderOption{}, lgcrq.order...),
		inters:     append([]Interceptor{}, lgcrq.inters...),
		predicates: append([]predicate.LottoGetCountRecord{}, lgcrq.predicates...),
		withUser:   lgcrq.withUser.Clone(),
		withLotto:  lgcrq.withLotto.Clone(),
		// clone intermediate query.
		sql:  lgcrq.sql.Clone(),
		path: lgcrq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (lgcrq *LottoGetCountRecordQuery) WithUser(opts ...func(*UserQuery)) *LottoGetCountRecordQuery {
	query := (&UserClient{config: lgcrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lgcrq.withUser = query
	return lgcrq
}

// WithLotto tells the query-builder to eager-load the nodes that are connected to
// the "lotto" edge. The optional arguments are used to configure the query builder of the edge.
func (lgcrq *LottoGetCountRecordQuery) WithLotto(opts ...func(*LottoQuery)) *LottoGetCountRecordQuery {
	query := (&LottoClient{config: lgcrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lgcrq.withLotto = query
	return lgcrq
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
//	client.LottoGetCountRecord.Query().
//		GroupBy(lottogetcountrecord.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (lgcrq *LottoGetCountRecordQuery) GroupBy(field string, fields ...string) *LottoGetCountRecordGroupBy {
	lgcrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LottoGetCountRecordGroupBy{build: lgcrq}
	grbuild.flds = &lgcrq.ctx.Fields
	grbuild.label = lottogetcountrecord.Label
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
//	client.LottoGetCountRecord.Query().
//		Select(lottogetcountrecord.FieldCreatedBy).
//		Scan(ctx, &v)
func (lgcrq *LottoGetCountRecordQuery) Select(fields ...string) *LottoGetCountRecordSelect {
	lgcrq.ctx.Fields = append(lgcrq.ctx.Fields, fields...)
	sbuild := &LottoGetCountRecordSelect{LottoGetCountRecordQuery: lgcrq}
	sbuild.label = lottogetcountrecord.Label
	sbuild.flds, sbuild.scan = &lgcrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LottoGetCountRecordSelect configured with the given aggregations.
func (lgcrq *LottoGetCountRecordQuery) Aggregate(fns ...AggregateFunc) *LottoGetCountRecordSelect {
	return lgcrq.Select().Aggregate(fns...)
}

func (lgcrq *LottoGetCountRecordQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lgcrq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lgcrq); err != nil {
				return err
			}
		}
	}
	for _, f := range lgcrq.ctx.Fields {
		if !lottogetcountrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if lgcrq.path != nil {
		prev, err := lgcrq.path(ctx)
		if err != nil {
			return err
		}
		lgcrq.sql = prev
	}
	return nil
}

func (lgcrq *LottoGetCountRecordQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LottoGetCountRecord, error) {
	var (
		nodes       = []*LottoGetCountRecord{}
		_spec       = lgcrq.querySpec()
		loadedTypes = [2]bool{
			lgcrq.withUser != nil,
			lgcrq.withLotto != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LottoGetCountRecord).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LottoGetCountRecord{config: lgcrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(lgcrq.modifiers) > 0 {
		_spec.Modifiers = lgcrq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lgcrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lgcrq.withUser; query != nil {
		if err := lgcrq.loadUser(ctx, query, nodes, nil,
			func(n *LottoGetCountRecord, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := lgcrq.withLotto; query != nil {
		if err := lgcrq.loadLotto(ctx, query, nodes, nil,
			func(n *LottoGetCountRecord, e *Lotto) { n.Edges.Lotto = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lgcrq *LottoGetCountRecordQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*LottoGetCountRecord, init func(*LottoGetCountRecord), assign func(*LottoGetCountRecord, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*LottoGetCountRecord)
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
func (lgcrq *LottoGetCountRecordQuery) loadLotto(ctx context.Context, query *LottoQuery, nodes []*LottoGetCountRecord, init func(*LottoGetCountRecord), assign func(*LottoGetCountRecord, *Lotto)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*LottoGetCountRecord)
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

func (lgcrq *LottoGetCountRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lgcrq.querySpec()
	if len(lgcrq.modifiers) > 0 {
		_spec.Modifiers = lgcrq.modifiers
	}
	_spec.Node.Columns = lgcrq.ctx.Fields
	if len(lgcrq.ctx.Fields) > 0 {
		_spec.Unique = lgcrq.ctx.Unique != nil && *lgcrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lgcrq.driver, _spec)
}

func (lgcrq *LottoGetCountRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(lottogetcountrecord.Table, lottogetcountrecord.Columns, sqlgraph.NewFieldSpec(lottogetcountrecord.FieldID, field.TypeInt64))
	_spec.From = lgcrq.sql
	if unique := lgcrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lgcrq.path != nil {
		_spec.Unique = true
	}
	if fields := lgcrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lottogetcountrecord.FieldID)
		for i := range fields {
			if fields[i] != lottogetcountrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if lgcrq.withUser != nil {
			_spec.Node.AddColumnOnce(lottogetcountrecord.FieldUserID)
		}
		if lgcrq.withLotto != nil {
			_spec.Node.AddColumnOnce(lottogetcountrecord.FieldLottoID)
		}
	}
	if ps := lgcrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lgcrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lgcrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lgcrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lgcrq *LottoGetCountRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lgcrq.driver.Dialect())
	t1 := builder.Table(lottogetcountrecord.Table)
	columns := lgcrq.ctx.Fields
	if len(columns) == 0 {
		columns = lottogetcountrecord.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lgcrq.sql != nil {
		selector = lgcrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lgcrq.ctx.Unique != nil && *lgcrq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range lgcrq.modifiers {
		m(selector)
	}
	for _, p := range lgcrq.predicates {
		p(selector)
	}
	for _, p := range lgcrq.order {
		p(selector)
	}
	if offset := lgcrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lgcrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lgcrq *LottoGetCountRecordQuery) Modify(modifiers ...func(s *sql.Selector)) *LottoGetCountRecordSelect {
	lgcrq.modifiers = append(lgcrq.modifiers, modifiers...)
	return lgcrq.Select()
}

// LottoGetCountRecordGroupBy is the group-by builder for LottoGetCountRecord entities.
type LottoGetCountRecordGroupBy struct {
	selector
	build *LottoGetCountRecordQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgcrgb *LottoGetCountRecordGroupBy) Aggregate(fns ...AggregateFunc) *LottoGetCountRecordGroupBy {
	lgcrgb.fns = append(lgcrgb.fns, fns...)
	return lgcrgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgcrgb *LottoGetCountRecordGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgcrgb.build.ctx, "GroupBy")
	if err := lgcrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LottoGetCountRecordQuery, *LottoGetCountRecordGroupBy](ctx, lgcrgb.build, lgcrgb, lgcrgb.build.inters, v)
}

func (lgcrgb *LottoGetCountRecordGroupBy) sqlScan(ctx context.Context, root *LottoGetCountRecordQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgcrgb.fns))
	for _, fn := range lgcrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgcrgb.flds)+len(lgcrgb.fns))
		for _, f := range *lgcrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgcrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgcrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LottoGetCountRecordSelect is the builder for selecting fields of LottoGetCountRecord entities.
type LottoGetCountRecordSelect struct {
	*LottoGetCountRecordQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lgcrs *LottoGetCountRecordSelect) Aggregate(fns ...AggregateFunc) *LottoGetCountRecordSelect {
	lgcrs.fns = append(lgcrs.fns, fns...)
	return lgcrs
}

// Scan applies the selector query and scans the result into the given value.
func (lgcrs *LottoGetCountRecordSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgcrs.ctx, "Select")
	if err := lgcrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LottoGetCountRecordQuery, *LottoGetCountRecordSelect](ctx, lgcrs.LottoGetCountRecordQuery, lgcrs, lgcrs.inters, v)
}

func (lgcrs *LottoGetCountRecordSelect) sqlScan(ctx context.Context, root *LottoGetCountRecordQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lgcrs.fns))
	for _, fn := range lgcrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lgcrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgcrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lgcrs *LottoGetCountRecordSelect) Modify(modifiers ...func(s *sql.Selector)) *LottoGetCountRecordSelect {
	lgcrs.modifiers = append(lgcrs.modifiers, modifiers...)
	return lgcrs
}

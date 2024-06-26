// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/cdkinfo"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// CDKInfoQuery is the builder for querying CDKInfo entities.
type CDKInfoQuery struct {
	config
	ctx           *QueryContext
	order         []cdkinfo.OrderOption
	inters        []Interceptor
	predicates    []predicate.CDKInfo
	withIssueUser *UserQuery
	withUseUser   *UserQuery
	modifiers     []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CDKInfoQuery builder.
func (ciq *CDKInfoQuery) Where(ps ...predicate.CDKInfo) *CDKInfoQuery {
	ciq.predicates = append(ciq.predicates, ps...)
	return ciq
}

// Limit the number of records to be returned by this query.
func (ciq *CDKInfoQuery) Limit(limit int) *CDKInfoQuery {
	ciq.ctx.Limit = &limit
	return ciq
}

// Offset to start from.
func (ciq *CDKInfoQuery) Offset(offset int) *CDKInfoQuery {
	ciq.ctx.Offset = &offset
	return ciq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ciq *CDKInfoQuery) Unique(unique bool) *CDKInfoQuery {
	ciq.ctx.Unique = &unique
	return ciq
}

// Order specifies how the records should be ordered.
func (ciq *CDKInfoQuery) Order(o ...cdkinfo.OrderOption) *CDKInfoQuery {
	ciq.order = append(ciq.order, o...)
	return ciq
}

// QueryIssueUser chains the current query on the "issue_user" edge.
func (ciq *CDKInfoQuery) QueryIssueUser() *UserQuery {
	query := (&UserClient{config: ciq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cdkinfo.Table, cdkinfo.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cdkinfo.IssueUserTable, cdkinfo.IssueUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUseUser chains the current query on the "use_user" edge.
func (ciq *CDKInfoQuery) QueryUseUser() *UserQuery {
	query := (&UserClient{config: ciq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cdkinfo.Table, cdkinfo.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cdkinfo.UseUserTable, cdkinfo.UseUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CDKInfo entity from the query.
// Returns a *NotFoundError when no CDKInfo was found.
func (ciq *CDKInfoQuery) First(ctx context.Context) (*CDKInfo, error) {
	nodes, err := ciq.Limit(1).All(setContextOp(ctx, ciq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cdkinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ciq *CDKInfoQuery) FirstX(ctx context.Context) *CDKInfo {
	node, err := ciq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CDKInfo ID from the query.
// Returns a *NotFoundError when no CDKInfo ID was found.
func (ciq *CDKInfoQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ciq.Limit(1).IDs(setContextOp(ctx, ciq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cdkinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ciq *CDKInfoQuery) FirstIDX(ctx context.Context) int64 {
	id, err := ciq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CDKInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CDKInfo entity is found.
// Returns a *NotFoundError when no CDKInfo entities are found.
func (ciq *CDKInfoQuery) Only(ctx context.Context) (*CDKInfo, error) {
	nodes, err := ciq.Limit(2).All(setContextOp(ctx, ciq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cdkinfo.Label}
	default:
		return nil, &NotSingularError{cdkinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ciq *CDKInfoQuery) OnlyX(ctx context.Context) *CDKInfo {
	node, err := ciq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CDKInfo ID in the query.
// Returns a *NotSingularError when more than one CDKInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (ciq *CDKInfoQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ciq.Limit(2).IDs(setContextOp(ctx, ciq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cdkinfo.Label}
	default:
		err = &NotSingularError{cdkinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ciq *CDKInfoQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := ciq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CDKInfos.
func (ciq *CDKInfoQuery) All(ctx context.Context) ([]*CDKInfo, error) {
	ctx = setContextOp(ctx, ciq.ctx, "All")
	if err := ciq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CDKInfo, *CDKInfoQuery]()
	return withInterceptors[[]*CDKInfo](ctx, ciq, qr, ciq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ciq *CDKInfoQuery) AllX(ctx context.Context) []*CDKInfo {
	nodes, err := ciq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CDKInfo IDs.
func (ciq *CDKInfoQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if ciq.ctx.Unique == nil && ciq.path != nil {
		ciq.Unique(true)
	}
	ctx = setContextOp(ctx, ciq.ctx, "IDs")
	if err = ciq.Select(cdkinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ciq *CDKInfoQuery) IDsX(ctx context.Context) []int64 {
	ids, err := ciq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ciq *CDKInfoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ciq.ctx, "Count")
	if err := ciq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ciq, querierCount[*CDKInfoQuery](), ciq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ciq *CDKInfoQuery) CountX(ctx context.Context) int {
	count, err := ciq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ciq *CDKInfoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ciq.ctx, "Exist")
	switch _, err := ciq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ciq *CDKInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := ciq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CDKInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ciq *CDKInfoQuery) Clone() *CDKInfoQuery {
	if ciq == nil {
		return nil
	}
	return &CDKInfoQuery{
		config:        ciq.config,
		ctx:           ciq.ctx.Clone(),
		order:         append([]cdkinfo.OrderOption{}, ciq.order...),
		inters:        append([]Interceptor{}, ciq.inters...),
		predicates:    append([]predicate.CDKInfo{}, ciq.predicates...),
		withIssueUser: ciq.withIssueUser.Clone(),
		withUseUser:   ciq.withUseUser.Clone(),
		// clone intermediate query.
		sql:  ciq.sql.Clone(),
		path: ciq.path,
	}
}

// WithIssueUser tells the query-builder to eager-load the nodes that are connected to
// the "issue_user" edge. The optional arguments are used to configure the query builder of the edge.
func (ciq *CDKInfoQuery) WithIssueUser(opts ...func(*UserQuery)) *CDKInfoQuery {
	query := (&UserClient{config: ciq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ciq.withIssueUser = query
	return ciq
}

// WithUseUser tells the query-builder to eager-load the nodes that are connected to
// the "use_user" edge. The optional arguments are used to configure the query builder of the edge.
func (ciq *CDKInfoQuery) WithUseUser(opts ...func(*UserQuery)) *CDKInfoQuery {
	query := (&UserClient{config: ciq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ciq.withUseUser = query
	return ciq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by,string"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CDKInfo.Query().
//		GroupBy(cdkinfo.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (ciq *CDKInfoQuery) GroupBy(field string, fields ...string) *CDKInfoGroupBy {
	ciq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CDKInfoGroupBy{build: ciq}
	grbuild.flds = &ciq.ctx.Fields
	grbuild.label = cdkinfo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by,string"`
//	}
//
//	client.CDKInfo.Query().
//		Select(cdkinfo.FieldCreatedBy).
//		Scan(ctx, &v)
func (ciq *CDKInfoQuery) Select(fields ...string) *CDKInfoSelect {
	ciq.ctx.Fields = append(ciq.ctx.Fields, fields...)
	sbuild := &CDKInfoSelect{CDKInfoQuery: ciq}
	sbuild.label = cdkinfo.Label
	sbuild.flds, sbuild.scan = &ciq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CDKInfoSelect configured with the given aggregations.
func (ciq *CDKInfoQuery) Aggregate(fns ...AggregateFunc) *CDKInfoSelect {
	return ciq.Select().Aggregate(fns...)
}

func (ciq *CDKInfoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ciq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ciq); err != nil {
				return err
			}
		}
	}
	for _, f := range ciq.ctx.Fields {
		if !cdkinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if ciq.path != nil {
		prev, err := ciq.path(ctx)
		if err != nil {
			return err
		}
		ciq.sql = prev
	}
	return nil
}

func (ciq *CDKInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CDKInfo, error) {
	var (
		nodes       = []*CDKInfo{}
		_spec       = ciq.querySpec()
		loadedTypes = [2]bool{
			ciq.withIssueUser != nil,
			ciq.withUseUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CDKInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CDKInfo{config: ciq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ciq.modifiers) > 0 {
		_spec.Modifiers = ciq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ciq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ciq.withIssueUser; query != nil {
		if err := ciq.loadIssueUser(ctx, query, nodes, nil,
			func(n *CDKInfo, e *User) { n.Edges.IssueUser = e }); err != nil {
			return nil, err
		}
	}
	if query := ciq.withUseUser; query != nil {
		if err := ciq.loadUseUser(ctx, query, nodes, nil,
			func(n *CDKInfo, e *User) { n.Edges.UseUser = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ciq *CDKInfoQuery) loadIssueUser(ctx context.Context, query *UserQuery, nodes []*CDKInfo, init func(*CDKInfo), assign func(*CDKInfo, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*CDKInfo)
	for i := range nodes {
		fk := nodes[i].IssueUserID
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
			return fmt.Errorf(`unexpected foreign-key "issue_user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ciq *CDKInfoQuery) loadUseUser(ctx context.Context, query *UserQuery, nodes []*CDKInfo, init func(*CDKInfo), assign func(*CDKInfo, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*CDKInfo)
	for i := range nodes {
		fk := nodes[i].UseUserID
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
			return fmt.Errorf(`unexpected foreign-key "use_user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ciq *CDKInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ciq.querySpec()
	if len(ciq.modifiers) > 0 {
		_spec.Modifiers = ciq.modifiers
	}
	_spec.Node.Columns = ciq.ctx.Fields
	if len(ciq.ctx.Fields) > 0 {
		_spec.Unique = ciq.ctx.Unique != nil && *ciq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ciq.driver, _spec)
}

func (ciq *CDKInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(cdkinfo.Table, cdkinfo.Columns, sqlgraph.NewFieldSpec(cdkinfo.FieldID, field.TypeInt64))
	_spec.From = ciq.sql
	if unique := ciq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ciq.path != nil {
		_spec.Unique = true
	}
	if fields := ciq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cdkinfo.FieldID)
		for i := range fields {
			if fields[i] != cdkinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ciq.withIssueUser != nil {
			_spec.Node.AddColumnOnce(cdkinfo.FieldIssueUserID)
		}
		if ciq.withUseUser != nil {
			_spec.Node.AddColumnOnce(cdkinfo.FieldUseUserID)
		}
	}
	if ps := ciq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ciq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ciq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ciq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ciq *CDKInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ciq.driver.Dialect())
	t1 := builder.Table(cdkinfo.Table)
	columns := ciq.ctx.Fields
	if len(columns) == 0 {
		columns = cdkinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ciq.sql != nil {
		selector = ciq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ciq.ctx.Unique != nil && *ciq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ciq.modifiers {
		m(selector)
	}
	for _, p := range ciq.predicates {
		p(selector)
	}
	for _, p := range ciq.order {
		p(selector)
	}
	if offset := ciq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ciq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ciq *CDKInfoQuery) Modify(modifiers ...func(s *sql.Selector)) *CDKInfoSelect {
	ciq.modifiers = append(ciq.modifiers, modifiers...)
	return ciq.Select()
}

// CDKInfoGroupBy is the group-by builder for CDKInfo entities.
type CDKInfoGroupBy struct {
	selector
	build *CDKInfoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cigb *CDKInfoGroupBy) Aggregate(fns ...AggregateFunc) *CDKInfoGroupBy {
	cigb.fns = append(cigb.fns, fns...)
	return cigb
}

// Scan applies the selector query and scans the result into the given value.
func (cigb *CDKInfoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cigb.build.ctx, "GroupBy")
	if err := cigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CDKInfoQuery, *CDKInfoGroupBy](ctx, cigb.build, cigb, cigb.build.inters, v)
}

func (cigb *CDKInfoGroupBy) sqlScan(ctx context.Context, root *CDKInfoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cigb.fns))
	for _, fn := range cigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cigb.flds)+len(cigb.fns))
		for _, f := range *cigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CDKInfoSelect is the builder for selecting fields of CDKInfo entities.
type CDKInfoSelect struct {
	*CDKInfoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cis *CDKInfoSelect) Aggregate(fns ...AggregateFunc) *CDKInfoSelect {
	cis.fns = append(cis.fns, fns...)
	return cis
}

// Scan applies the selector query and scans the result into the given value.
func (cis *CDKInfoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cis.ctx, "Select")
	if err := cis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CDKInfoQuery, *CDKInfoSelect](ctx, cis.CDKInfoQuery, cis, cis.inters, v)
}

func (cis *CDKInfoSelect) sqlScan(ctx context.Context, root *CDKInfoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cis.fns))
	for _, fn := range cis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cis *CDKInfoSelect) Modify(modifiers ...func(s *sql.Selector)) *CDKInfoSelect {
	cis.modifiers = append(cis.modifiers, modifiers...)
	return cis
}

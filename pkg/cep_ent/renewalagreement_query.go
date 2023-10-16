// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/mission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/renewalagreement"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// RenewalAgreementQuery is the builder for querying RenewalAgreement entities.
type RenewalAgreementQuery struct {
	config
	ctx         *QueryContext
	order       []renewalagreement.OrderOption
	inters      []Interceptor
	predicates  []predicate.RenewalAgreement
	withUser    *UserQuery
	withMission *MissionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RenewalAgreementQuery builder.
func (raq *RenewalAgreementQuery) Where(ps ...predicate.RenewalAgreement) *RenewalAgreementQuery {
	raq.predicates = append(raq.predicates, ps...)
	return raq
}

// Limit the number of records to be returned by this query.
func (raq *RenewalAgreementQuery) Limit(limit int) *RenewalAgreementQuery {
	raq.ctx.Limit = &limit
	return raq
}

// Offset to start from.
func (raq *RenewalAgreementQuery) Offset(offset int) *RenewalAgreementQuery {
	raq.ctx.Offset = &offset
	return raq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (raq *RenewalAgreementQuery) Unique(unique bool) *RenewalAgreementQuery {
	raq.ctx.Unique = &unique
	return raq
}

// Order specifies how the records should be ordered.
func (raq *RenewalAgreementQuery) Order(o ...renewalagreement.OrderOption) *RenewalAgreementQuery {
	raq.order = append(raq.order, o...)
	return raq
}

// QueryUser chains the current query on the "user" edge.
func (raq *RenewalAgreementQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: raq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := raq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := raq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(renewalagreement.Table, renewalagreement.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, renewalagreement.UserTable, renewalagreement.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(raq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMission chains the current query on the "mission" edge.
func (raq *RenewalAgreementQuery) QueryMission() *MissionQuery {
	query := (&MissionClient{config: raq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := raq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := raq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(renewalagreement.Table, renewalagreement.FieldID, selector),
			sqlgraph.To(mission.Table, mission.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, renewalagreement.MissionTable, renewalagreement.MissionColumn),
		)
		fromU = sqlgraph.SetNeighbors(raq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RenewalAgreement entity from the query.
// Returns a *NotFoundError when no RenewalAgreement was found.
func (raq *RenewalAgreementQuery) First(ctx context.Context) (*RenewalAgreement, error) {
	nodes, err := raq.Limit(1).All(setContextOp(ctx, raq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{renewalagreement.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (raq *RenewalAgreementQuery) FirstX(ctx context.Context) *RenewalAgreement {
	node, err := raq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RenewalAgreement ID from the query.
// Returns a *NotFoundError when no RenewalAgreement ID was found.
func (raq *RenewalAgreementQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = raq.Limit(1).IDs(setContextOp(ctx, raq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{renewalagreement.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (raq *RenewalAgreementQuery) FirstIDX(ctx context.Context) int64 {
	id, err := raq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RenewalAgreement entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RenewalAgreement entity is found.
// Returns a *NotFoundError when no RenewalAgreement entities are found.
func (raq *RenewalAgreementQuery) Only(ctx context.Context) (*RenewalAgreement, error) {
	nodes, err := raq.Limit(2).All(setContextOp(ctx, raq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{renewalagreement.Label}
	default:
		return nil, &NotSingularError{renewalagreement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (raq *RenewalAgreementQuery) OnlyX(ctx context.Context) *RenewalAgreement {
	node, err := raq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RenewalAgreement ID in the query.
// Returns a *NotSingularError when more than one RenewalAgreement ID is found.
// Returns a *NotFoundError when no entities are found.
func (raq *RenewalAgreementQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = raq.Limit(2).IDs(setContextOp(ctx, raq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{renewalagreement.Label}
	default:
		err = &NotSingularError{renewalagreement.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (raq *RenewalAgreementQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := raq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RenewalAgreements.
func (raq *RenewalAgreementQuery) All(ctx context.Context) ([]*RenewalAgreement, error) {
	ctx = setContextOp(ctx, raq.ctx, "All")
	if err := raq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RenewalAgreement, *RenewalAgreementQuery]()
	return withInterceptors[[]*RenewalAgreement](ctx, raq, qr, raq.inters)
}

// AllX is like All, but panics if an error occurs.
func (raq *RenewalAgreementQuery) AllX(ctx context.Context) []*RenewalAgreement {
	nodes, err := raq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RenewalAgreement IDs.
func (raq *RenewalAgreementQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if raq.ctx.Unique == nil && raq.path != nil {
		raq.Unique(true)
	}
	ctx = setContextOp(ctx, raq.ctx, "IDs")
	if err = raq.Select(renewalagreement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (raq *RenewalAgreementQuery) IDsX(ctx context.Context) []int64 {
	ids, err := raq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (raq *RenewalAgreementQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, raq.ctx, "Count")
	if err := raq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, raq, querierCount[*RenewalAgreementQuery](), raq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (raq *RenewalAgreementQuery) CountX(ctx context.Context) int {
	count, err := raq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (raq *RenewalAgreementQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, raq.ctx, "Exist")
	switch _, err := raq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (raq *RenewalAgreementQuery) ExistX(ctx context.Context) bool {
	exist, err := raq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RenewalAgreementQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (raq *RenewalAgreementQuery) Clone() *RenewalAgreementQuery {
	if raq == nil {
		return nil
	}
	return &RenewalAgreementQuery{
		config:      raq.config,
		ctx:         raq.ctx.Clone(),
		order:       append([]renewalagreement.OrderOption{}, raq.order...),
		inters:      append([]Interceptor{}, raq.inters...),
		predicates:  append([]predicate.RenewalAgreement{}, raq.predicates...),
		withUser:    raq.withUser.Clone(),
		withMission: raq.withMission.Clone(),
		// clone intermediate query.
		sql:  raq.sql.Clone(),
		path: raq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (raq *RenewalAgreementQuery) WithUser(opts ...func(*UserQuery)) *RenewalAgreementQuery {
	query := (&UserClient{config: raq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	raq.withUser = query
	return raq
}

// WithMission tells the query-builder to eager-load the nodes that are connected to
// the "mission" edge. The optional arguments are used to configure the query builder of the edge.
func (raq *RenewalAgreementQuery) WithMission(opts ...func(*MissionQuery)) *RenewalAgreementQuery {
	query := (&MissionClient{config: raq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	raq.withMission = query
	return raq
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
//	client.RenewalAgreement.Query().
//		GroupBy(renewalagreement.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (raq *RenewalAgreementQuery) GroupBy(field string, fields ...string) *RenewalAgreementGroupBy {
	raq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RenewalAgreementGroupBy{build: raq}
	grbuild.flds = &raq.ctx.Fields
	grbuild.label = renewalagreement.Label
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
//	client.RenewalAgreement.Query().
//		Select(renewalagreement.FieldCreatedBy).
//		Scan(ctx, &v)
func (raq *RenewalAgreementQuery) Select(fields ...string) *RenewalAgreementSelect {
	raq.ctx.Fields = append(raq.ctx.Fields, fields...)
	sbuild := &RenewalAgreementSelect{RenewalAgreementQuery: raq}
	sbuild.label = renewalagreement.Label
	sbuild.flds, sbuild.scan = &raq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RenewalAgreementSelect configured with the given aggregations.
func (raq *RenewalAgreementQuery) Aggregate(fns ...AggregateFunc) *RenewalAgreementSelect {
	return raq.Select().Aggregate(fns...)
}

func (raq *RenewalAgreementQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range raq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, raq); err != nil {
				return err
			}
		}
	}
	for _, f := range raq.ctx.Fields {
		if !renewalagreement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if raq.path != nil {
		prev, err := raq.path(ctx)
		if err != nil {
			return err
		}
		raq.sql = prev
	}
	return nil
}

func (raq *RenewalAgreementQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RenewalAgreement, error) {
	var (
		nodes       = []*RenewalAgreement{}
		_spec       = raq.querySpec()
		loadedTypes = [2]bool{
			raq.withUser != nil,
			raq.withMission != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RenewalAgreement).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RenewalAgreement{config: raq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, raq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := raq.withUser; query != nil {
		if err := raq.loadUser(ctx, query, nodes, nil,
			func(n *RenewalAgreement, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := raq.withMission; query != nil {
		if err := raq.loadMission(ctx, query, nodes, nil,
			func(n *RenewalAgreement, e *Mission) { n.Edges.Mission = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (raq *RenewalAgreementQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*RenewalAgreement, init func(*RenewalAgreement), assign func(*RenewalAgreement, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*RenewalAgreement)
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
func (raq *RenewalAgreementQuery) loadMission(ctx context.Context, query *MissionQuery, nodes []*RenewalAgreement, init func(*RenewalAgreement), assign func(*RenewalAgreement, *Mission)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*RenewalAgreement)
	for i := range nodes {
		fk := nodes[i].MissionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(mission.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "mission_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (raq *RenewalAgreementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := raq.querySpec()
	_spec.Node.Columns = raq.ctx.Fields
	if len(raq.ctx.Fields) > 0 {
		_spec.Unique = raq.ctx.Unique != nil && *raq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, raq.driver, _spec)
}

func (raq *RenewalAgreementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(renewalagreement.Table, renewalagreement.Columns, sqlgraph.NewFieldSpec(renewalagreement.FieldID, field.TypeInt64))
	_spec.From = raq.sql
	if unique := raq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if raq.path != nil {
		_spec.Unique = true
	}
	if fields := raq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, renewalagreement.FieldID)
		for i := range fields {
			if fields[i] != renewalagreement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if raq.withUser != nil {
			_spec.Node.AddColumnOnce(renewalagreement.FieldUserID)
		}
		if raq.withMission != nil {
			_spec.Node.AddColumnOnce(renewalagreement.FieldMissionID)
		}
	}
	if ps := raq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := raq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := raq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := raq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (raq *RenewalAgreementQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(raq.driver.Dialect())
	t1 := builder.Table(renewalagreement.Table)
	columns := raq.ctx.Fields
	if len(columns) == 0 {
		columns = renewalagreement.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if raq.sql != nil {
		selector = raq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if raq.ctx.Unique != nil && *raq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range raq.predicates {
		p(selector)
	}
	for _, p := range raq.order {
		p(selector)
	}
	if offset := raq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := raq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RenewalAgreementGroupBy is the group-by builder for RenewalAgreement entities.
type RenewalAgreementGroupBy struct {
	selector
	build *RenewalAgreementQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ragb *RenewalAgreementGroupBy) Aggregate(fns ...AggregateFunc) *RenewalAgreementGroupBy {
	ragb.fns = append(ragb.fns, fns...)
	return ragb
}

// Scan applies the selector query and scans the result into the given value.
func (ragb *RenewalAgreementGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ragb.build.ctx, "GroupBy")
	if err := ragb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RenewalAgreementQuery, *RenewalAgreementGroupBy](ctx, ragb.build, ragb, ragb.build.inters, v)
}

func (ragb *RenewalAgreementGroupBy) sqlScan(ctx context.Context, root *RenewalAgreementQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ragb.fns))
	for _, fn := range ragb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ragb.flds)+len(ragb.fns))
		for _, f := range *ragb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ragb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ragb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RenewalAgreementSelect is the builder for selecting fields of RenewalAgreement entities.
type RenewalAgreementSelect struct {
	*RenewalAgreementQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ras *RenewalAgreementSelect) Aggregate(fns ...AggregateFunc) *RenewalAgreementSelect {
	ras.fns = append(ras.fns, fns...)
	return ras
}

// Scan applies the selector query and scans the result into the given value.
func (ras *RenewalAgreementSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ras.ctx, "Select")
	if err := ras.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RenewalAgreementQuery, *RenewalAgreementSelect](ctx, ras.RenewalAgreementQuery, ras, ras.inters, v)
}

func (ras *RenewalAgreementSelect) sqlScan(ctx context.Context, root *RenewalAgreementQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ras.fns))
	for _, fn := range ras.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ras.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ras.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

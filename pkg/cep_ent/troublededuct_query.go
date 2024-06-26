// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/device"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/troublededuct"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// TroubleDeductQuery is the builder for querying TroubleDeduct entities.
type TroubleDeductQuery struct {
	config
	ctx        *QueryContext
	order      []troublededuct.OrderOption
	inters     []Interceptor
	predicates []predicate.TroubleDeduct
	withUser   *UserQuery
	withDevice *DeviceQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TroubleDeductQuery builder.
func (tdq *TroubleDeductQuery) Where(ps ...predicate.TroubleDeduct) *TroubleDeductQuery {
	tdq.predicates = append(tdq.predicates, ps...)
	return tdq
}

// Limit the number of records to be returned by this query.
func (tdq *TroubleDeductQuery) Limit(limit int) *TroubleDeductQuery {
	tdq.ctx.Limit = &limit
	return tdq
}

// Offset to start from.
func (tdq *TroubleDeductQuery) Offset(offset int) *TroubleDeductQuery {
	tdq.ctx.Offset = &offset
	return tdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tdq *TroubleDeductQuery) Unique(unique bool) *TroubleDeductQuery {
	tdq.ctx.Unique = &unique
	return tdq
}

// Order specifies how the records should be ordered.
func (tdq *TroubleDeductQuery) Order(o ...troublededuct.OrderOption) *TroubleDeductQuery {
	tdq.order = append(tdq.order, o...)
	return tdq
}

// QueryUser chains the current query on the "user" edge.
func (tdq *TroubleDeductQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: tdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(troublededuct.Table, troublededuct.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, troublededuct.UserTable, troublededuct.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(tdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDevice chains the current query on the "device" edge.
func (tdq *TroubleDeductQuery) QueryDevice() *DeviceQuery {
	query := (&DeviceClient{config: tdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(troublededuct.Table, troublededuct.FieldID, selector),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, troublededuct.DeviceTable, troublededuct.DeviceColumn),
		)
		fromU = sqlgraph.SetNeighbors(tdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TroubleDeduct entity from the query.
// Returns a *NotFoundError when no TroubleDeduct was found.
func (tdq *TroubleDeductQuery) First(ctx context.Context) (*TroubleDeduct, error) {
	nodes, err := tdq.Limit(1).All(setContextOp(ctx, tdq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{troublededuct.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tdq *TroubleDeductQuery) FirstX(ctx context.Context) *TroubleDeduct {
	node, err := tdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TroubleDeduct ID from the query.
// Returns a *NotFoundError when no TroubleDeduct ID was found.
func (tdq *TroubleDeductQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = tdq.Limit(1).IDs(setContextOp(ctx, tdq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{troublededuct.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tdq *TroubleDeductQuery) FirstIDX(ctx context.Context) int64 {
	id, err := tdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TroubleDeduct entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TroubleDeduct entity is found.
// Returns a *NotFoundError when no TroubleDeduct entities are found.
func (tdq *TroubleDeductQuery) Only(ctx context.Context) (*TroubleDeduct, error) {
	nodes, err := tdq.Limit(2).All(setContextOp(ctx, tdq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{troublededuct.Label}
	default:
		return nil, &NotSingularError{troublededuct.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tdq *TroubleDeductQuery) OnlyX(ctx context.Context) *TroubleDeduct {
	node, err := tdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TroubleDeduct ID in the query.
// Returns a *NotSingularError when more than one TroubleDeduct ID is found.
// Returns a *NotFoundError when no entities are found.
func (tdq *TroubleDeductQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = tdq.Limit(2).IDs(setContextOp(ctx, tdq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{troublededuct.Label}
	default:
		err = &NotSingularError{troublededuct.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tdq *TroubleDeductQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := tdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TroubleDeducts.
func (tdq *TroubleDeductQuery) All(ctx context.Context) ([]*TroubleDeduct, error) {
	ctx = setContextOp(ctx, tdq.ctx, "All")
	if err := tdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TroubleDeduct, *TroubleDeductQuery]()
	return withInterceptors[[]*TroubleDeduct](ctx, tdq, qr, tdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tdq *TroubleDeductQuery) AllX(ctx context.Context) []*TroubleDeduct {
	nodes, err := tdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TroubleDeduct IDs.
func (tdq *TroubleDeductQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if tdq.ctx.Unique == nil && tdq.path != nil {
		tdq.Unique(true)
	}
	ctx = setContextOp(ctx, tdq.ctx, "IDs")
	if err = tdq.Select(troublededuct.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tdq *TroubleDeductQuery) IDsX(ctx context.Context) []int64 {
	ids, err := tdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tdq *TroubleDeductQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tdq.ctx, "Count")
	if err := tdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tdq, querierCount[*TroubleDeductQuery](), tdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tdq *TroubleDeductQuery) CountX(ctx context.Context) int {
	count, err := tdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tdq *TroubleDeductQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tdq.ctx, "Exist")
	switch _, err := tdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tdq *TroubleDeductQuery) ExistX(ctx context.Context) bool {
	exist, err := tdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TroubleDeductQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tdq *TroubleDeductQuery) Clone() *TroubleDeductQuery {
	if tdq == nil {
		return nil
	}
	return &TroubleDeductQuery{
		config:     tdq.config,
		ctx:        tdq.ctx.Clone(),
		order:      append([]troublededuct.OrderOption{}, tdq.order...),
		inters:     append([]Interceptor{}, tdq.inters...),
		predicates: append([]predicate.TroubleDeduct{}, tdq.predicates...),
		withUser:   tdq.withUser.Clone(),
		withDevice: tdq.withDevice.Clone(),
		// clone intermediate query.
		sql:  tdq.sql.Clone(),
		path: tdq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (tdq *TroubleDeductQuery) WithUser(opts ...func(*UserQuery)) *TroubleDeductQuery {
	query := (&UserClient{config: tdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tdq.withUser = query
	return tdq
}

// WithDevice tells the query-builder to eager-load the nodes that are connected to
// the "device" edge. The optional arguments are used to configure the query builder of the edge.
func (tdq *TroubleDeductQuery) WithDevice(opts ...func(*DeviceQuery)) *TroubleDeductQuery {
	query := (&DeviceClient{config: tdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tdq.withDevice = query
	return tdq
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
//	client.TroubleDeduct.Query().
//		GroupBy(troublededuct.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (tdq *TroubleDeductQuery) GroupBy(field string, fields ...string) *TroubleDeductGroupBy {
	tdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TroubleDeductGroupBy{build: tdq}
	grbuild.flds = &tdq.ctx.Fields
	grbuild.label = troublededuct.Label
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
//	client.TroubleDeduct.Query().
//		Select(troublededuct.FieldCreatedBy).
//		Scan(ctx, &v)
func (tdq *TroubleDeductQuery) Select(fields ...string) *TroubleDeductSelect {
	tdq.ctx.Fields = append(tdq.ctx.Fields, fields...)
	sbuild := &TroubleDeductSelect{TroubleDeductQuery: tdq}
	sbuild.label = troublededuct.Label
	sbuild.flds, sbuild.scan = &tdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TroubleDeductSelect configured with the given aggregations.
func (tdq *TroubleDeductQuery) Aggregate(fns ...AggregateFunc) *TroubleDeductSelect {
	return tdq.Select().Aggregate(fns...)
}

func (tdq *TroubleDeductQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tdq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tdq); err != nil {
				return err
			}
		}
	}
	for _, f := range tdq.ctx.Fields {
		if !troublededuct.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if tdq.path != nil {
		prev, err := tdq.path(ctx)
		if err != nil {
			return err
		}
		tdq.sql = prev
	}
	return nil
}

func (tdq *TroubleDeductQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TroubleDeduct, error) {
	var (
		nodes       = []*TroubleDeduct{}
		_spec       = tdq.querySpec()
		loadedTypes = [2]bool{
			tdq.withUser != nil,
			tdq.withDevice != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TroubleDeduct).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TroubleDeduct{config: tdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(tdq.modifiers) > 0 {
		_spec.Modifiers = tdq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tdq.withUser; query != nil {
		if err := tdq.loadUser(ctx, query, nodes, nil,
			func(n *TroubleDeduct, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := tdq.withDevice; query != nil {
		if err := tdq.loadDevice(ctx, query, nodes, nil,
			func(n *TroubleDeduct, e *Device) { n.Edges.Device = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tdq *TroubleDeductQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*TroubleDeduct, init func(*TroubleDeduct), assign func(*TroubleDeduct, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*TroubleDeduct)
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
func (tdq *TroubleDeductQuery) loadDevice(ctx context.Context, query *DeviceQuery, nodes []*TroubleDeduct, init func(*TroubleDeduct), assign func(*TroubleDeduct, *Device)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*TroubleDeduct)
	for i := range nodes {
		fk := nodes[i].DeviceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(device.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "device_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tdq *TroubleDeductQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tdq.querySpec()
	if len(tdq.modifiers) > 0 {
		_spec.Modifiers = tdq.modifiers
	}
	_spec.Node.Columns = tdq.ctx.Fields
	if len(tdq.ctx.Fields) > 0 {
		_spec.Unique = tdq.ctx.Unique != nil && *tdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tdq.driver, _spec)
}

func (tdq *TroubleDeductQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(troublededuct.Table, troublededuct.Columns, sqlgraph.NewFieldSpec(troublededuct.FieldID, field.TypeInt64))
	_spec.From = tdq.sql
	if unique := tdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tdq.path != nil {
		_spec.Unique = true
	}
	if fields := tdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, troublededuct.FieldID)
		for i := range fields {
			if fields[i] != troublededuct.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if tdq.withUser != nil {
			_spec.Node.AddColumnOnce(troublededuct.FieldUserID)
		}
		if tdq.withDevice != nil {
			_spec.Node.AddColumnOnce(troublededuct.FieldDeviceID)
		}
	}
	if ps := tdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tdq *TroubleDeductQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tdq.driver.Dialect())
	t1 := builder.Table(troublededuct.Table)
	columns := tdq.ctx.Fields
	if len(columns) == 0 {
		columns = troublededuct.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tdq.sql != nil {
		selector = tdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tdq.ctx.Unique != nil && *tdq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range tdq.modifiers {
		m(selector)
	}
	for _, p := range tdq.predicates {
		p(selector)
	}
	for _, p := range tdq.order {
		p(selector)
	}
	if offset := tdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tdq *TroubleDeductQuery) Modify(modifiers ...func(s *sql.Selector)) *TroubleDeductSelect {
	tdq.modifiers = append(tdq.modifiers, modifiers...)
	return tdq.Select()
}

// TroubleDeductGroupBy is the group-by builder for TroubleDeduct entities.
type TroubleDeductGroupBy struct {
	selector
	build *TroubleDeductQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tdgb *TroubleDeductGroupBy) Aggregate(fns ...AggregateFunc) *TroubleDeductGroupBy {
	tdgb.fns = append(tdgb.fns, fns...)
	return tdgb
}

// Scan applies the selector query and scans the result into the given value.
func (tdgb *TroubleDeductGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tdgb.build.ctx, "GroupBy")
	if err := tdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TroubleDeductQuery, *TroubleDeductGroupBy](ctx, tdgb.build, tdgb, tdgb.build.inters, v)
}

func (tdgb *TroubleDeductGroupBy) sqlScan(ctx context.Context, root *TroubleDeductQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tdgb.fns))
	for _, fn := range tdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tdgb.flds)+len(tdgb.fns))
		for _, f := range *tdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TroubleDeductSelect is the builder for selecting fields of TroubleDeduct entities.
type TroubleDeductSelect struct {
	*TroubleDeductQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tds *TroubleDeductSelect) Aggregate(fns ...AggregateFunc) *TroubleDeductSelect {
	tds.fns = append(tds.fns, fns...)
	return tds
}

// Scan applies the selector query and scans the result into the given value.
func (tds *TroubleDeductSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tds.ctx, "Select")
	if err := tds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TroubleDeductQuery, *TroubleDeductSelect](ctx, tds.TroubleDeductQuery, tds, tds.inters, v)
}

func (tds *TroubleDeductSelect) sqlScan(ctx context.Context, root *TroubleDeductQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tds.fns))
	for _, fn := range tds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tds *TroubleDeductSelect) Modify(modifiers ...func(s *sql.Selector)) *TroubleDeductSelect {
	tds.modifiers = append(tds.modifiers, modifiers...)
	return tds
}

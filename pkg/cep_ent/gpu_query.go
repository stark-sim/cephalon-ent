// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/devicegpumission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/gpu"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/price"
)

// GpuQuery is the builder for querying Gpu entities.
type GpuQuery struct {
	config
	ctx                   *QueryContext
	order                 []gpu.OrderOption
	inters                []Interceptor
	predicates            []predicate.Gpu
	withDeviceGpuMissions *DeviceGpuMissionQuery
	withPrices            *PriceQuery
	modifiers             []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GpuQuery builder.
func (gq *GpuQuery) Where(ps ...predicate.Gpu) *GpuQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit the number of records to be returned by this query.
func (gq *GpuQuery) Limit(limit int) *GpuQuery {
	gq.ctx.Limit = &limit
	return gq
}

// Offset to start from.
func (gq *GpuQuery) Offset(offset int) *GpuQuery {
	gq.ctx.Offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GpuQuery) Unique(unique bool) *GpuQuery {
	gq.ctx.Unique = &unique
	return gq
}

// Order specifies how the records should be ordered.
func (gq *GpuQuery) Order(o ...gpu.OrderOption) *GpuQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// QueryDeviceGpuMissions chains the current query on the "device_gpu_missions" edge.
func (gq *GpuQuery) QueryDeviceGpuMissions() *DeviceGpuMissionQuery {
	query := (&DeviceGpuMissionClient{config: gq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(gpu.Table, gpu.FieldID, selector),
			sqlgraph.To(devicegpumission.Table, devicegpumission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, gpu.DeviceGpuMissionsTable, gpu.DeviceGpuMissionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPrices chains the current query on the "prices" edge.
func (gq *GpuQuery) QueryPrices() *PriceQuery {
	query := (&PriceClient{config: gq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(gpu.Table, gpu.FieldID, selector),
			sqlgraph.To(price.Table, price.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, gpu.PricesTable, gpu.PricesColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Gpu entity from the query.
// Returns a *NotFoundError when no Gpu was found.
func (gq *GpuQuery) First(ctx context.Context) (*Gpu, error) {
	nodes, err := gq.Limit(1).All(setContextOp(ctx, gq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{gpu.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GpuQuery) FirstX(ctx context.Context) *Gpu {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Gpu ID from the query.
// Returns a *NotFoundError when no Gpu ID was found.
func (gq *GpuQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = gq.Limit(1).IDs(setContextOp(ctx, gq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{gpu.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GpuQuery) FirstIDX(ctx context.Context) int64 {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Gpu entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Gpu entity is found.
// Returns a *NotFoundError when no Gpu entities are found.
func (gq *GpuQuery) Only(ctx context.Context) (*Gpu, error) {
	nodes, err := gq.Limit(2).All(setContextOp(ctx, gq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{gpu.Label}
	default:
		return nil, &NotSingularError{gpu.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GpuQuery) OnlyX(ctx context.Context) *Gpu {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Gpu ID in the query.
// Returns a *NotSingularError when more than one Gpu ID is found.
// Returns a *NotFoundError when no entities are found.
func (gq *GpuQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = gq.Limit(2).IDs(setContextOp(ctx, gq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{gpu.Label}
	default:
		err = &NotSingularError{gpu.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GpuQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Gpus.
func (gq *GpuQuery) All(ctx context.Context) ([]*Gpu, error) {
	ctx = setContextOp(ctx, gq.ctx, "All")
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Gpu, *GpuQuery]()
	return withInterceptors[[]*Gpu](ctx, gq, qr, gq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gq *GpuQuery) AllX(ctx context.Context) []*Gpu {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Gpu IDs.
func (gq *GpuQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if gq.ctx.Unique == nil && gq.path != nil {
		gq.Unique(true)
	}
	ctx = setContextOp(ctx, gq.ctx, "IDs")
	if err = gq.Select(gpu.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GpuQuery) IDsX(ctx context.Context) []int64 {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GpuQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gq.ctx, "Count")
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gq, querierCount[*GpuQuery](), gq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GpuQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GpuQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gq.ctx, "Exist")
	switch _, err := gq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GpuQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GpuQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GpuQuery) Clone() *GpuQuery {
	if gq == nil {
		return nil
	}
	return &GpuQuery{
		config:                gq.config,
		ctx:                   gq.ctx.Clone(),
		order:                 append([]gpu.OrderOption{}, gq.order...),
		inters:                append([]Interceptor{}, gq.inters...),
		predicates:            append([]predicate.Gpu{}, gq.predicates...),
		withDeviceGpuMissions: gq.withDeviceGpuMissions.Clone(),
		withPrices:            gq.withPrices.Clone(),
		// clone intermediate query.
		sql:  gq.sql.Clone(),
		path: gq.path,
	}
}

// WithDeviceGpuMissions tells the query-builder to eager-load the nodes that are connected to
// the "device_gpu_missions" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GpuQuery) WithDeviceGpuMissions(opts ...func(*DeviceGpuMissionQuery)) *GpuQuery {
	query := (&DeviceGpuMissionClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gq.withDeviceGpuMissions = query
	return gq
}

// WithPrices tells the query-builder to eager-load the nodes that are connected to
// the "prices" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GpuQuery) WithPrices(opts ...func(*PriceQuery)) *GpuQuery {
	query := (&PriceClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gq.withPrices = query
	return gq
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
//	client.Gpu.Query().
//		GroupBy(gpu.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (gq *GpuQuery) GroupBy(field string, fields ...string) *GpuGroupBy {
	gq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GpuGroupBy{build: gq}
	grbuild.flds = &gq.ctx.Fields
	grbuild.label = gpu.Label
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
//	client.Gpu.Query().
//		Select(gpu.FieldCreatedBy).
//		Scan(ctx, &v)
func (gq *GpuQuery) Select(fields ...string) *GpuSelect {
	gq.ctx.Fields = append(gq.ctx.Fields, fields...)
	sbuild := &GpuSelect{GpuQuery: gq}
	sbuild.label = gpu.Label
	sbuild.flds, sbuild.scan = &gq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GpuSelect configured with the given aggregations.
func (gq *GpuQuery) Aggregate(fns ...AggregateFunc) *GpuSelect {
	return gq.Select().Aggregate(fns...)
}

func (gq *GpuQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gq); err != nil {
				return err
			}
		}
	}
	for _, f := range gq.ctx.Fields {
		if !gpu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GpuQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Gpu, error) {
	var (
		nodes       = []*Gpu{}
		_spec       = gq.querySpec()
		loadedTypes = [2]bool{
			gq.withDeviceGpuMissions != nil,
			gq.withPrices != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Gpu).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Gpu{config: gq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gq.withDeviceGpuMissions; query != nil {
		if err := gq.loadDeviceGpuMissions(ctx, query, nodes,
			func(n *Gpu) { n.Edges.DeviceGpuMissions = []*DeviceGpuMission{} },
			func(n *Gpu, e *DeviceGpuMission) { n.Edges.DeviceGpuMissions = append(n.Edges.DeviceGpuMissions, e) }); err != nil {
			return nil, err
		}
	}
	if query := gq.withPrices; query != nil {
		if err := gq.loadPrices(ctx, query, nodes,
			func(n *Gpu) { n.Edges.Prices = []*Price{} },
			func(n *Gpu, e *Price) { n.Edges.Prices = append(n.Edges.Prices, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gq *GpuQuery) loadDeviceGpuMissions(ctx context.Context, query *DeviceGpuMissionQuery, nodes []*Gpu, init func(*Gpu), assign func(*Gpu, *DeviceGpuMission)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Gpu)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(devicegpumission.FieldGpuID)
	}
	query.Where(predicate.DeviceGpuMission(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(gpu.DeviceGpuMissionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.GpuID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "gpu_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (gq *GpuQuery) loadPrices(ctx context.Context, query *PriceQuery, nodes []*Gpu, init func(*Gpu), assign func(*Gpu, *Price)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Gpu)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(price.FieldGpuID)
	}
	query.Where(predicate.Price(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(gpu.PricesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.GpuID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "gpu_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (gq *GpuQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	_spec.Node.Columns = gq.ctx.Fields
	if len(gq.ctx.Fields) > 0 {
		_spec.Unique = gq.ctx.Unique != nil && *gq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GpuQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(gpu.Table, gpu.Columns, sqlgraph.NewFieldSpec(gpu.FieldID, field.TypeInt64))
	_spec.From = gq.sql
	if unique := gq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gq.path != nil {
		_spec.Unique = true
	}
	if fields := gq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gpu.FieldID)
		for i := range fields {
			if fields[i] != gpu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GpuQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(gpu.Table)
	columns := gq.ctx.Fields
	if len(columns) == 0 {
		columns = gpu.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gq.ctx.Unique != nil && *gq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range gq.modifiers {
		m(selector)
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gq *GpuQuery) Modify(modifiers ...func(s *sql.Selector)) *GpuSelect {
	gq.modifiers = append(gq.modifiers, modifiers...)
	return gq.Select()
}

// GpuGroupBy is the group-by builder for Gpu entities.
type GpuGroupBy struct {
	selector
	build *GpuQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GpuGroupBy) Aggregate(fns ...AggregateFunc) *GpuGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the selector query and scans the result into the given value.
func (ggb *GpuGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ggb.build.ctx, "GroupBy")
	if err := ggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GpuQuery, *GpuGroupBy](ctx, ggb.build, ggb, ggb.build.inters, v)
}

func (ggb *GpuGroupBy) sqlScan(ctx context.Context, root *GpuQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ggb.flds)+len(ggb.fns))
		for _, f := range *ggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GpuSelect is the builder for selecting fields of Gpu entities.
type GpuSelect struct {
	*GpuQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gs *GpuSelect) Aggregate(fns ...AggregateFunc) *GpuSelect {
	gs.fns = append(gs.fns, fns...)
	return gs
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GpuSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gs.ctx, "Select")
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GpuQuery, *GpuSelect](ctx, gs.GpuQuery, gs, gs.inters, v)
}

func (gs *GpuSelect) sqlScan(ctx context.Context, root *GpuQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gs.fns))
	for _, fn := range gs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gs *GpuSelect) Modify(modifiers ...func(s *sql.Selector)) *GpuSelect {
	gs.modifiers = append(gs.modifiers, modifiers...)
	return gs
}

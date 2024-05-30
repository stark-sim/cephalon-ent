// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/extraserviceorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/mission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionbatch"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/symbol"
)

// ExtraServiceOrderQuery is the builder for querying ExtraServiceOrder entities.
type ExtraServiceOrderQuery struct {
	config
	ctx              *QueryContext
	order            []extraserviceorder.OrderOption
	inters           []Interceptor
	predicates       []predicate.ExtraServiceOrder
	withMission      *MissionQuery
	withMissionOrder *MissionOrderQuery
	withSymbol       *SymbolQuery
	withMissionBatch *MissionBatchQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExtraServiceOrderQuery builder.
func (esoq *ExtraServiceOrderQuery) Where(ps ...predicate.ExtraServiceOrder) *ExtraServiceOrderQuery {
	esoq.predicates = append(esoq.predicates, ps...)
	return esoq
}

// Limit the number of records to be returned by this query.
func (esoq *ExtraServiceOrderQuery) Limit(limit int) *ExtraServiceOrderQuery {
	esoq.ctx.Limit = &limit
	return esoq
}

// Offset to start from.
func (esoq *ExtraServiceOrderQuery) Offset(offset int) *ExtraServiceOrderQuery {
	esoq.ctx.Offset = &offset
	return esoq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (esoq *ExtraServiceOrderQuery) Unique(unique bool) *ExtraServiceOrderQuery {
	esoq.ctx.Unique = &unique
	return esoq
}

// Order specifies how the records should be ordered.
func (esoq *ExtraServiceOrderQuery) Order(o ...extraserviceorder.OrderOption) *ExtraServiceOrderQuery {
	esoq.order = append(esoq.order, o...)
	return esoq
}

// QueryMission chains the current query on the "mission" edge.
func (esoq *ExtraServiceOrderQuery) QueryMission() *MissionQuery {
	query := (&MissionClient{config: esoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := esoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := esoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(extraserviceorder.Table, extraserviceorder.FieldID, selector),
			sqlgraph.To(mission.Table, mission.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, extraserviceorder.MissionTable, extraserviceorder.MissionColumn),
		)
		fromU = sqlgraph.SetNeighbors(esoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMissionOrder chains the current query on the "mission_order" edge.
func (esoq *ExtraServiceOrderQuery) QueryMissionOrder() *MissionOrderQuery {
	query := (&MissionOrderClient{config: esoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := esoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := esoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(extraserviceorder.Table, extraserviceorder.FieldID, selector),
			sqlgraph.To(missionorder.Table, missionorder.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, extraserviceorder.MissionOrderTable, extraserviceorder.MissionOrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(esoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySymbol chains the current query on the "symbol" edge.
func (esoq *ExtraServiceOrderQuery) QuerySymbol() *SymbolQuery {
	query := (&SymbolClient{config: esoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := esoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := esoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(extraserviceorder.Table, extraserviceorder.FieldID, selector),
			sqlgraph.To(symbol.Table, symbol.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, extraserviceorder.SymbolTable, extraserviceorder.SymbolColumn),
		)
		fromU = sqlgraph.SetNeighbors(esoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMissionBatch chains the current query on the "mission_batch" edge.
func (esoq *ExtraServiceOrderQuery) QueryMissionBatch() *MissionBatchQuery {
	query := (&MissionBatchClient{config: esoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := esoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := esoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(extraserviceorder.Table, extraserviceorder.FieldID, selector),
			sqlgraph.To(missionbatch.Table, missionbatch.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, extraserviceorder.MissionBatchTable, extraserviceorder.MissionBatchColumn),
		)
		fromU = sqlgraph.SetNeighbors(esoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ExtraServiceOrder entity from the query.
// Returns a *NotFoundError when no ExtraServiceOrder was found.
func (esoq *ExtraServiceOrderQuery) First(ctx context.Context) (*ExtraServiceOrder, error) {
	nodes, err := esoq.Limit(1).All(setContextOp(ctx, esoq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{extraserviceorder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (esoq *ExtraServiceOrderQuery) FirstX(ctx context.Context) *ExtraServiceOrder {
	node, err := esoq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ExtraServiceOrder ID from the query.
// Returns a *NotFoundError when no ExtraServiceOrder ID was found.
func (esoq *ExtraServiceOrderQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = esoq.Limit(1).IDs(setContextOp(ctx, esoq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{extraserviceorder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (esoq *ExtraServiceOrderQuery) FirstIDX(ctx context.Context) int64 {
	id, err := esoq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ExtraServiceOrder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ExtraServiceOrder entity is found.
// Returns a *NotFoundError when no ExtraServiceOrder entities are found.
func (esoq *ExtraServiceOrderQuery) Only(ctx context.Context) (*ExtraServiceOrder, error) {
	nodes, err := esoq.Limit(2).All(setContextOp(ctx, esoq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{extraserviceorder.Label}
	default:
		return nil, &NotSingularError{extraserviceorder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (esoq *ExtraServiceOrderQuery) OnlyX(ctx context.Context) *ExtraServiceOrder {
	node, err := esoq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ExtraServiceOrder ID in the query.
// Returns a *NotSingularError when more than one ExtraServiceOrder ID is found.
// Returns a *NotFoundError when no entities are found.
func (esoq *ExtraServiceOrderQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = esoq.Limit(2).IDs(setContextOp(ctx, esoq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{extraserviceorder.Label}
	default:
		err = &NotSingularError{extraserviceorder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (esoq *ExtraServiceOrderQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := esoq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ExtraServiceOrders.
func (esoq *ExtraServiceOrderQuery) All(ctx context.Context) ([]*ExtraServiceOrder, error) {
	ctx = setContextOp(ctx, esoq.ctx, "All")
	if err := esoq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ExtraServiceOrder, *ExtraServiceOrderQuery]()
	return withInterceptors[[]*ExtraServiceOrder](ctx, esoq, qr, esoq.inters)
}

// AllX is like All, but panics if an error occurs.
func (esoq *ExtraServiceOrderQuery) AllX(ctx context.Context) []*ExtraServiceOrder {
	nodes, err := esoq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ExtraServiceOrder IDs.
func (esoq *ExtraServiceOrderQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if esoq.ctx.Unique == nil && esoq.path != nil {
		esoq.Unique(true)
	}
	ctx = setContextOp(ctx, esoq.ctx, "IDs")
	if err = esoq.Select(extraserviceorder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (esoq *ExtraServiceOrderQuery) IDsX(ctx context.Context) []int64 {
	ids, err := esoq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (esoq *ExtraServiceOrderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, esoq.ctx, "Count")
	if err := esoq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, esoq, querierCount[*ExtraServiceOrderQuery](), esoq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (esoq *ExtraServiceOrderQuery) CountX(ctx context.Context) int {
	count, err := esoq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (esoq *ExtraServiceOrderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, esoq.ctx, "Exist")
	switch _, err := esoq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (esoq *ExtraServiceOrderQuery) ExistX(ctx context.Context) bool {
	exist, err := esoq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExtraServiceOrderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (esoq *ExtraServiceOrderQuery) Clone() *ExtraServiceOrderQuery {
	if esoq == nil {
		return nil
	}
	return &ExtraServiceOrderQuery{
		config:           esoq.config,
		ctx:              esoq.ctx.Clone(),
		order:            append([]extraserviceorder.OrderOption{}, esoq.order...),
		inters:           append([]Interceptor{}, esoq.inters...),
		predicates:       append([]predicate.ExtraServiceOrder{}, esoq.predicates...),
		withMission:      esoq.withMission.Clone(),
		withMissionOrder: esoq.withMissionOrder.Clone(),
		withSymbol:       esoq.withSymbol.Clone(),
		withMissionBatch: esoq.withMissionBatch.Clone(),
		// clone intermediate query.
		sql:  esoq.sql.Clone(),
		path: esoq.path,
	}
}

// WithMission tells the query-builder to eager-load the nodes that are connected to
// the "mission" edge. The optional arguments are used to configure the query builder of the edge.
func (esoq *ExtraServiceOrderQuery) WithMission(opts ...func(*MissionQuery)) *ExtraServiceOrderQuery {
	query := (&MissionClient{config: esoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	esoq.withMission = query
	return esoq
}

// WithMissionOrder tells the query-builder to eager-load the nodes that are connected to
// the "mission_order" edge. The optional arguments are used to configure the query builder of the edge.
func (esoq *ExtraServiceOrderQuery) WithMissionOrder(opts ...func(*MissionOrderQuery)) *ExtraServiceOrderQuery {
	query := (&MissionOrderClient{config: esoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	esoq.withMissionOrder = query
	return esoq
}

// WithSymbol tells the query-builder to eager-load the nodes that are connected to
// the "symbol" edge. The optional arguments are used to configure the query builder of the edge.
func (esoq *ExtraServiceOrderQuery) WithSymbol(opts ...func(*SymbolQuery)) *ExtraServiceOrderQuery {
	query := (&SymbolClient{config: esoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	esoq.withSymbol = query
	return esoq
}

// WithMissionBatch tells the query-builder to eager-load the nodes that are connected to
// the "mission_batch" edge. The optional arguments are used to configure the query builder of the edge.
func (esoq *ExtraServiceOrderQuery) WithMissionBatch(opts ...func(*MissionBatchQuery)) *ExtraServiceOrderQuery {
	query := (&MissionBatchClient{config: esoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	esoq.withMissionBatch = query
	return esoq
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
//	client.ExtraServiceOrder.Query().
//		GroupBy(extraserviceorder.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (esoq *ExtraServiceOrderQuery) GroupBy(field string, fields ...string) *ExtraServiceOrderGroupBy {
	esoq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ExtraServiceOrderGroupBy{build: esoq}
	grbuild.flds = &esoq.ctx.Fields
	grbuild.label = extraserviceorder.Label
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
//	client.ExtraServiceOrder.Query().
//		Select(extraserviceorder.FieldCreatedBy).
//		Scan(ctx, &v)
func (esoq *ExtraServiceOrderQuery) Select(fields ...string) *ExtraServiceOrderSelect {
	esoq.ctx.Fields = append(esoq.ctx.Fields, fields...)
	sbuild := &ExtraServiceOrderSelect{ExtraServiceOrderQuery: esoq}
	sbuild.label = extraserviceorder.Label
	sbuild.flds, sbuild.scan = &esoq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ExtraServiceOrderSelect configured with the given aggregations.
func (esoq *ExtraServiceOrderQuery) Aggregate(fns ...AggregateFunc) *ExtraServiceOrderSelect {
	return esoq.Select().Aggregate(fns...)
}

func (esoq *ExtraServiceOrderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range esoq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, esoq); err != nil {
				return err
			}
		}
	}
	for _, f := range esoq.ctx.Fields {
		if !extraserviceorder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if esoq.path != nil {
		prev, err := esoq.path(ctx)
		if err != nil {
			return err
		}
		esoq.sql = prev
	}
	return nil
}

func (esoq *ExtraServiceOrderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ExtraServiceOrder, error) {
	var (
		nodes       = []*ExtraServiceOrder{}
		_spec       = esoq.querySpec()
		loadedTypes = [4]bool{
			esoq.withMission != nil,
			esoq.withMissionOrder != nil,
			esoq.withSymbol != nil,
			esoq.withMissionBatch != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ExtraServiceOrder).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ExtraServiceOrder{config: esoq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(esoq.modifiers) > 0 {
		_spec.Modifiers = esoq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, esoq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := esoq.withMission; query != nil {
		if err := esoq.loadMission(ctx, query, nodes, nil,
			func(n *ExtraServiceOrder, e *Mission) { n.Edges.Mission = e }); err != nil {
			return nil, err
		}
	}
	if query := esoq.withMissionOrder; query != nil {
		if err := esoq.loadMissionOrder(ctx, query, nodes, nil,
			func(n *ExtraServiceOrder, e *MissionOrder) { n.Edges.MissionOrder = e }); err != nil {
			return nil, err
		}
	}
	if query := esoq.withSymbol; query != nil {
		if err := esoq.loadSymbol(ctx, query, nodes, nil,
			func(n *ExtraServiceOrder, e *Symbol) { n.Edges.Symbol = e }); err != nil {
			return nil, err
		}
	}
	if query := esoq.withMissionBatch; query != nil {
		if err := esoq.loadMissionBatch(ctx, query, nodes, nil,
			func(n *ExtraServiceOrder, e *MissionBatch) { n.Edges.MissionBatch = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (esoq *ExtraServiceOrderQuery) loadMission(ctx context.Context, query *MissionQuery, nodes []*ExtraServiceOrder, init func(*ExtraServiceOrder), assign func(*ExtraServiceOrder, *Mission)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*ExtraServiceOrder)
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
func (esoq *ExtraServiceOrderQuery) loadMissionOrder(ctx context.Context, query *MissionOrderQuery, nodes []*ExtraServiceOrder, init func(*ExtraServiceOrder), assign func(*ExtraServiceOrder, *MissionOrder)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*ExtraServiceOrder)
	for i := range nodes {
		fk := nodes[i].MissionOrderID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(missionorder.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "mission_order_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (esoq *ExtraServiceOrderQuery) loadSymbol(ctx context.Context, query *SymbolQuery, nodes []*ExtraServiceOrder, init func(*ExtraServiceOrder), assign func(*ExtraServiceOrder, *Symbol)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*ExtraServiceOrder)
	for i := range nodes {
		fk := nodes[i].SymbolID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(symbol.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "symbol_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (esoq *ExtraServiceOrderQuery) loadMissionBatch(ctx context.Context, query *MissionBatchQuery, nodes []*ExtraServiceOrder, init func(*ExtraServiceOrder), assign func(*ExtraServiceOrder, *MissionBatch)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*ExtraServiceOrder)
	for i := range nodes {
		fk := nodes[i].MissionBatchID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(missionbatch.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "mission_batch_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (esoq *ExtraServiceOrderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := esoq.querySpec()
	if len(esoq.modifiers) > 0 {
		_spec.Modifiers = esoq.modifiers
	}
	_spec.Node.Columns = esoq.ctx.Fields
	if len(esoq.ctx.Fields) > 0 {
		_spec.Unique = esoq.ctx.Unique != nil && *esoq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, esoq.driver, _spec)
}

func (esoq *ExtraServiceOrderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(extraserviceorder.Table, extraserviceorder.Columns, sqlgraph.NewFieldSpec(extraserviceorder.FieldID, field.TypeInt64))
	_spec.From = esoq.sql
	if unique := esoq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if esoq.path != nil {
		_spec.Unique = true
	}
	if fields := esoq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, extraserviceorder.FieldID)
		for i := range fields {
			if fields[i] != extraserviceorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if esoq.withMission != nil {
			_spec.Node.AddColumnOnce(extraserviceorder.FieldMissionID)
		}
		if esoq.withMissionOrder != nil {
			_spec.Node.AddColumnOnce(extraserviceorder.FieldMissionOrderID)
		}
		if esoq.withSymbol != nil {
			_spec.Node.AddColumnOnce(extraserviceorder.FieldSymbolID)
		}
		if esoq.withMissionBatch != nil {
			_spec.Node.AddColumnOnce(extraserviceorder.FieldMissionBatchID)
		}
	}
	if ps := esoq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := esoq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := esoq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := esoq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (esoq *ExtraServiceOrderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(esoq.driver.Dialect())
	t1 := builder.Table(extraserviceorder.Table)
	columns := esoq.ctx.Fields
	if len(columns) == 0 {
		columns = extraserviceorder.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if esoq.sql != nil {
		selector = esoq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if esoq.ctx.Unique != nil && *esoq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range esoq.modifiers {
		m(selector)
	}
	for _, p := range esoq.predicates {
		p(selector)
	}
	for _, p := range esoq.order {
		p(selector)
	}
	if offset := esoq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := esoq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (esoq *ExtraServiceOrderQuery) Modify(modifiers ...func(s *sql.Selector)) *ExtraServiceOrderSelect {
	esoq.modifiers = append(esoq.modifiers, modifiers...)
	return esoq.Select()
}

// ExtraServiceOrderGroupBy is the group-by builder for ExtraServiceOrder entities.
type ExtraServiceOrderGroupBy struct {
	selector
	build *ExtraServiceOrderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (esogb *ExtraServiceOrderGroupBy) Aggregate(fns ...AggregateFunc) *ExtraServiceOrderGroupBy {
	esogb.fns = append(esogb.fns, fns...)
	return esogb
}

// Scan applies the selector query and scans the result into the given value.
func (esogb *ExtraServiceOrderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, esogb.build.ctx, "GroupBy")
	if err := esogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExtraServiceOrderQuery, *ExtraServiceOrderGroupBy](ctx, esogb.build, esogb, esogb.build.inters, v)
}

func (esogb *ExtraServiceOrderGroupBy) sqlScan(ctx context.Context, root *ExtraServiceOrderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(esogb.fns))
	for _, fn := range esogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*esogb.flds)+len(esogb.fns))
		for _, f := range *esogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*esogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := esogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ExtraServiceOrderSelect is the builder for selecting fields of ExtraServiceOrder entities.
type ExtraServiceOrderSelect struct {
	*ExtraServiceOrderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (esos *ExtraServiceOrderSelect) Aggregate(fns ...AggregateFunc) *ExtraServiceOrderSelect {
	esos.fns = append(esos.fns, fns...)
	return esos
}

// Scan applies the selector query and scans the result into the given value.
func (esos *ExtraServiceOrderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, esos.ctx, "Select")
	if err := esos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExtraServiceOrderQuery, *ExtraServiceOrderSelect](ctx, esos.ExtraServiceOrderQuery, esos, esos.inters, v)
}

func (esos *ExtraServiceOrderSelect) sqlScan(ctx context.Context, root *ExtraServiceOrderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(esos.fns))
	for _, fn := range esos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*esos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := esos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (esos *ExtraServiceOrderSelect) Modify(modifiers ...func(s *sql.Selector)) *ExtraServiceOrderSelect {
	esos.modifiers = append(esos.modifiers, modifiers...)
	return esos
}

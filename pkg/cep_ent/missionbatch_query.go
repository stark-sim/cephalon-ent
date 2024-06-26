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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/extraserviceorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/mission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionbatch"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionconsumeorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// MissionBatchQuery is the builder for querying MissionBatch entities.
type MissionBatchQuery struct {
	config
	ctx                      *QueryContext
	order                    []missionbatch.OrderOption
	inters                   []Interceptor
	predicates               []predicate.MissionBatch
	withUser                 *UserQuery
	withMissionConsumeOrders *MissionConsumeOrderQuery
	withMissions             *MissionQuery
	withMissionOrders        *MissionOrderQuery
	withExtraServiceOrder    *ExtraServiceOrderQuery
	modifiers                []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MissionBatchQuery builder.
func (mbq *MissionBatchQuery) Where(ps ...predicate.MissionBatch) *MissionBatchQuery {
	mbq.predicates = append(mbq.predicates, ps...)
	return mbq
}

// Limit the number of records to be returned by this query.
func (mbq *MissionBatchQuery) Limit(limit int) *MissionBatchQuery {
	mbq.ctx.Limit = &limit
	return mbq
}

// Offset to start from.
func (mbq *MissionBatchQuery) Offset(offset int) *MissionBatchQuery {
	mbq.ctx.Offset = &offset
	return mbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mbq *MissionBatchQuery) Unique(unique bool) *MissionBatchQuery {
	mbq.ctx.Unique = &unique
	return mbq
}

// Order specifies how the records should be ordered.
func (mbq *MissionBatchQuery) Order(o ...missionbatch.OrderOption) *MissionBatchQuery {
	mbq.order = append(mbq.order, o...)
	return mbq
}

// QueryUser chains the current query on the "user" edge.
func (mbq *MissionBatchQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: mbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionbatch.Table, missionbatch.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, missionbatch.UserTable, missionbatch.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(mbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMissionConsumeOrders chains the current query on the "mission_consume_orders" edge.
func (mbq *MissionBatchQuery) QueryMissionConsumeOrders() *MissionConsumeOrderQuery {
	query := (&MissionConsumeOrderClient{config: mbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionbatch.Table, missionbatch.FieldID, selector),
			sqlgraph.To(missionconsumeorder.Table, missionconsumeorder.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, missionbatch.MissionConsumeOrdersTable, missionbatch.MissionConsumeOrdersColumn),
		)
		fromU = sqlgraph.SetNeighbors(mbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMissions chains the current query on the "missions" edge.
func (mbq *MissionBatchQuery) QueryMissions() *MissionQuery {
	query := (&MissionClient{config: mbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionbatch.Table, missionbatch.FieldID, selector),
			sqlgraph.To(mission.Table, mission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, missionbatch.MissionsTable, missionbatch.MissionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMissionOrders chains the current query on the "mission_orders" edge.
func (mbq *MissionBatchQuery) QueryMissionOrders() *MissionOrderQuery {
	query := (&MissionOrderClient{config: mbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionbatch.Table, missionbatch.FieldID, selector),
			sqlgraph.To(missionorder.Table, missionorder.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, missionbatch.MissionOrdersTable, missionbatch.MissionOrdersColumn),
		)
		fromU = sqlgraph.SetNeighbors(mbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExtraServiceOrder chains the current query on the "extra_service_order" edge.
func (mbq *MissionBatchQuery) QueryExtraServiceOrder() *ExtraServiceOrderQuery {
	query := (&ExtraServiceOrderClient{config: mbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionbatch.Table, missionbatch.FieldID, selector),
			sqlgraph.To(extraserviceorder.Table, extraserviceorder.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, missionbatch.ExtraServiceOrderTable, missionbatch.ExtraServiceOrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(mbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MissionBatch entity from the query.
// Returns a *NotFoundError when no MissionBatch was found.
func (mbq *MissionBatchQuery) First(ctx context.Context) (*MissionBatch, error) {
	nodes, err := mbq.Limit(1).All(setContextOp(ctx, mbq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{missionbatch.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mbq *MissionBatchQuery) FirstX(ctx context.Context) *MissionBatch {
	node, err := mbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MissionBatch ID from the query.
// Returns a *NotFoundError when no MissionBatch ID was found.
func (mbq *MissionBatchQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mbq.Limit(1).IDs(setContextOp(ctx, mbq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{missionbatch.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mbq *MissionBatchQuery) FirstIDX(ctx context.Context) int64 {
	id, err := mbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MissionBatch entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MissionBatch entity is found.
// Returns a *NotFoundError when no MissionBatch entities are found.
func (mbq *MissionBatchQuery) Only(ctx context.Context) (*MissionBatch, error) {
	nodes, err := mbq.Limit(2).All(setContextOp(ctx, mbq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{missionbatch.Label}
	default:
		return nil, &NotSingularError{missionbatch.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mbq *MissionBatchQuery) OnlyX(ctx context.Context) *MissionBatch {
	node, err := mbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MissionBatch ID in the query.
// Returns a *NotSingularError when more than one MissionBatch ID is found.
// Returns a *NotFoundError when no entities are found.
func (mbq *MissionBatchQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mbq.Limit(2).IDs(setContextOp(ctx, mbq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{missionbatch.Label}
	default:
		err = &NotSingularError{missionbatch.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mbq *MissionBatchQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := mbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MissionBatches.
func (mbq *MissionBatchQuery) All(ctx context.Context) ([]*MissionBatch, error) {
	ctx = setContextOp(ctx, mbq.ctx, "All")
	if err := mbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MissionBatch, *MissionBatchQuery]()
	return withInterceptors[[]*MissionBatch](ctx, mbq, qr, mbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mbq *MissionBatchQuery) AllX(ctx context.Context) []*MissionBatch {
	nodes, err := mbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MissionBatch IDs.
func (mbq *MissionBatchQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if mbq.ctx.Unique == nil && mbq.path != nil {
		mbq.Unique(true)
	}
	ctx = setContextOp(ctx, mbq.ctx, "IDs")
	if err = mbq.Select(missionbatch.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mbq *MissionBatchQuery) IDsX(ctx context.Context) []int64 {
	ids, err := mbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mbq *MissionBatchQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mbq.ctx, "Count")
	if err := mbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mbq, querierCount[*MissionBatchQuery](), mbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mbq *MissionBatchQuery) CountX(ctx context.Context) int {
	count, err := mbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mbq *MissionBatchQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mbq.ctx, "Exist")
	switch _, err := mbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mbq *MissionBatchQuery) ExistX(ctx context.Context) bool {
	exist, err := mbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MissionBatchQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mbq *MissionBatchQuery) Clone() *MissionBatchQuery {
	if mbq == nil {
		return nil
	}
	return &MissionBatchQuery{
		config:                   mbq.config,
		ctx:                      mbq.ctx.Clone(),
		order:                    append([]missionbatch.OrderOption{}, mbq.order...),
		inters:                   append([]Interceptor{}, mbq.inters...),
		predicates:               append([]predicate.MissionBatch{}, mbq.predicates...),
		withUser:                 mbq.withUser.Clone(),
		withMissionConsumeOrders: mbq.withMissionConsumeOrders.Clone(),
		withMissions:             mbq.withMissions.Clone(),
		withMissionOrders:        mbq.withMissionOrders.Clone(),
		withExtraServiceOrder:    mbq.withExtraServiceOrder.Clone(),
		// clone intermediate query.
		sql:  mbq.sql.Clone(),
		path: mbq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (mbq *MissionBatchQuery) WithUser(opts ...func(*UserQuery)) *MissionBatchQuery {
	query := (&UserClient{config: mbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mbq.withUser = query
	return mbq
}

// WithMissionConsumeOrders tells the query-builder to eager-load the nodes that are connected to
// the "mission_consume_orders" edge. The optional arguments are used to configure the query builder of the edge.
func (mbq *MissionBatchQuery) WithMissionConsumeOrders(opts ...func(*MissionConsumeOrderQuery)) *MissionBatchQuery {
	query := (&MissionConsumeOrderClient{config: mbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mbq.withMissionConsumeOrders = query
	return mbq
}

// WithMissions tells the query-builder to eager-load the nodes that are connected to
// the "missions" edge. The optional arguments are used to configure the query builder of the edge.
func (mbq *MissionBatchQuery) WithMissions(opts ...func(*MissionQuery)) *MissionBatchQuery {
	query := (&MissionClient{config: mbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mbq.withMissions = query
	return mbq
}

// WithMissionOrders tells the query-builder to eager-load the nodes that are connected to
// the "mission_orders" edge. The optional arguments are used to configure the query builder of the edge.
func (mbq *MissionBatchQuery) WithMissionOrders(opts ...func(*MissionOrderQuery)) *MissionBatchQuery {
	query := (&MissionOrderClient{config: mbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mbq.withMissionOrders = query
	return mbq
}

// WithExtraServiceOrder tells the query-builder to eager-load the nodes that are connected to
// the "extra_service_order" edge. The optional arguments are used to configure the query builder of the edge.
func (mbq *MissionBatchQuery) WithExtraServiceOrder(opts ...func(*ExtraServiceOrderQuery)) *MissionBatchQuery {
	query := (&ExtraServiceOrderClient{config: mbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mbq.withExtraServiceOrder = query
	return mbq
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
//	client.MissionBatch.Query().
//		GroupBy(missionbatch.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (mbq *MissionBatchQuery) GroupBy(field string, fields ...string) *MissionBatchGroupBy {
	mbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MissionBatchGroupBy{build: mbq}
	grbuild.flds = &mbq.ctx.Fields
	grbuild.label = missionbatch.Label
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
//	client.MissionBatch.Query().
//		Select(missionbatch.FieldCreatedBy).
//		Scan(ctx, &v)
func (mbq *MissionBatchQuery) Select(fields ...string) *MissionBatchSelect {
	mbq.ctx.Fields = append(mbq.ctx.Fields, fields...)
	sbuild := &MissionBatchSelect{MissionBatchQuery: mbq}
	sbuild.label = missionbatch.Label
	sbuild.flds, sbuild.scan = &mbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MissionBatchSelect configured with the given aggregations.
func (mbq *MissionBatchQuery) Aggregate(fns ...AggregateFunc) *MissionBatchSelect {
	return mbq.Select().Aggregate(fns...)
}

func (mbq *MissionBatchQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mbq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mbq); err != nil {
				return err
			}
		}
	}
	for _, f := range mbq.ctx.Fields {
		if !missionbatch.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if mbq.path != nil {
		prev, err := mbq.path(ctx)
		if err != nil {
			return err
		}
		mbq.sql = prev
	}
	return nil
}

func (mbq *MissionBatchQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MissionBatch, error) {
	var (
		nodes       = []*MissionBatch{}
		_spec       = mbq.querySpec()
		loadedTypes = [5]bool{
			mbq.withUser != nil,
			mbq.withMissionConsumeOrders != nil,
			mbq.withMissions != nil,
			mbq.withMissionOrders != nil,
			mbq.withExtraServiceOrder != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MissionBatch).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MissionBatch{config: mbq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(mbq.modifiers) > 0 {
		_spec.Modifiers = mbq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mbq.withUser; query != nil {
		if err := mbq.loadUser(ctx, query, nodes, nil,
			func(n *MissionBatch, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := mbq.withMissionConsumeOrders; query != nil {
		if err := mbq.loadMissionConsumeOrders(ctx, query, nodes,
			func(n *MissionBatch) { n.Edges.MissionConsumeOrders = []*MissionConsumeOrder{} },
			func(n *MissionBatch, e *MissionConsumeOrder) {
				n.Edges.MissionConsumeOrders = append(n.Edges.MissionConsumeOrders, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := mbq.withMissions; query != nil {
		if err := mbq.loadMissions(ctx, query, nodes,
			func(n *MissionBatch) { n.Edges.Missions = []*Mission{} },
			func(n *MissionBatch, e *Mission) { n.Edges.Missions = append(n.Edges.Missions, e) }); err != nil {
			return nil, err
		}
	}
	if query := mbq.withMissionOrders; query != nil {
		if err := mbq.loadMissionOrders(ctx, query, nodes,
			func(n *MissionBatch) { n.Edges.MissionOrders = []*MissionOrder{} },
			func(n *MissionBatch, e *MissionOrder) { n.Edges.MissionOrders = append(n.Edges.MissionOrders, e) }); err != nil {
			return nil, err
		}
	}
	if query := mbq.withExtraServiceOrder; query != nil {
		if err := mbq.loadExtraServiceOrder(ctx, query, nodes,
			func(n *MissionBatch) { n.Edges.ExtraServiceOrder = []*ExtraServiceOrder{} },
			func(n *MissionBatch, e *ExtraServiceOrder) {
				n.Edges.ExtraServiceOrder = append(n.Edges.ExtraServiceOrder, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mbq *MissionBatchQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*MissionBatch, init func(*MissionBatch), assign func(*MissionBatch, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MissionBatch)
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
func (mbq *MissionBatchQuery) loadMissionConsumeOrders(ctx context.Context, query *MissionConsumeOrderQuery, nodes []*MissionBatch, init func(*MissionBatch), assign func(*MissionBatch, *MissionConsumeOrder)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*MissionBatch)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(missionconsumeorder.FieldMissionBatchID)
	}
	query.Where(predicate.MissionConsumeOrder(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(missionbatch.MissionConsumeOrdersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MissionBatchID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "mission_batch_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mbq *MissionBatchQuery) loadMissions(ctx context.Context, query *MissionQuery, nodes []*MissionBatch, init func(*MissionBatch), assign func(*MissionBatch, *Mission)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*MissionBatch)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(mission.FieldMissionBatchID)
	}
	query.Where(predicate.Mission(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(missionbatch.MissionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MissionBatchID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "mission_batch_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mbq *MissionBatchQuery) loadMissionOrders(ctx context.Context, query *MissionOrderQuery, nodes []*MissionBatch, init func(*MissionBatch), assign func(*MissionBatch, *MissionOrder)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*MissionBatch)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(missionorder.FieldMissionBatchID)
	}
	query.Where(predicate.MissionOrder(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(missionbatch.MissionOrdersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MissionBatchID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "mission_batch_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mbq *MissionBatchQuery) loadExtraServiceOrder(ctx context.Context, query *ExtraServiceOrderQuery, nodes []*MissionBatch, init func(*MissionBatch), assign func(*MissionBatch, *ExtraServiceOrder)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*MissionBatch)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(extraserviceorder.FieldMissionBatchID)
	}
	query.Where(predicate.ExtraServiceOrder(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(missionbatch.ExtraServiceOrderColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MissionBatchID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "mission_batch_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mbq *MissionBatchQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mbq.querySpec()
	if len(mbq.modifiers) > 0 {
		_spec.Modifiers = mbq.modifiers
	}
	_spec.Node.Columns = mbq.ctx.Fields
	if len(mbq.ctx.Fields) > 0 {
		_spec.Unique = mbq.ctx.Unique != nil && *mbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mbq.driver, _spec)
}

func (mbq *MissionBatchQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(missionbatch.Table, missionbatch.Columns, sqlgraph.NewFieldSpec(missionbatch.FieldID, field.TypeInt64))
	_spec.From = mbq.sql
	if unique := mbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mbq.path != nil {
		_spec.Unique = true
	}
	if fields := mbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, missionbatch.FieldID)
		for i := range fields {
			if fields[i] != missionbatch.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mbq.withUser != nil {
			_spec.Node.AddColumnOnce(missionbatch.FieldUserID)
		}
	}
	if ps := mbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mbq *MissionBatchQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mbq.driver.Dialect())
	t1 := builder.Table(missionbatch.Table)
	columns := mbq.ctx.Fields
	if len(columns) == 0 {
		columns = missionbatch.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mbq.sql != nil {
		selector = mbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mbq.ctx.Unique != nil && *mbq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range mbq.modifiers {
		m(selector)
	}
	for _, p := range mbq.predicates {
		p(selector)
	}
	for _, p := range mbq.order {
		p(selector)
	}
	if offset := mbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mbq *MissionBatchQuery) Modify(modifiers ...func(s *sql.Selector)) *MissionBatchSelect {
	mbq.modifiers = append(mbq.modifiers, modifiers...)
	return mbq.Select()
}

// MissionBatchGroupBy is the group-by builder for MissionBatch entities.
type MissionBatchGroupBy struct {
	selector
	build *MissionBatchQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mbgb *MissionBatchGroupBy) Aggregate(fns ...AggregateFunc) *MissionBatchGroupBy {
	mbgb.fns = append(mbgb.fns, fns...)
	return mbgb
}

// Scan applies the selector query and scans the result into the given value.
func (mbgb *MissionBatchGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mbgb.build.ctx, "GroupBy")
	if err := mbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MissionBatchQuery, *MissionBatchGroupBy](ctx, mbgb.build, mbgb, mbgb.build.inters, v)
}

func (mbgb *MissionBatchGroupBy) sqlScan(ctx context.Context, root *MissionBatchQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mbgb.fns))
	for _, fn := range mbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mbgb.flds)+len(mbgb.fns))
		for _, f := range *mbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MissionBatchSelect is the builder for selecting fields of MissionBatch entities.
type MissionBatchSelect struct {
	*MissionBatchQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mbs *MissionBatchSelect) Aggregate(fns ...AggregateFunc) *MissionBatchSelect {
	mbs.fns = append(mbs.fns, fns...)
	return mbs
}

// Scan applies the selector query and scans the result into the given value.
func (mbs *MissionBatchSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mbs.ctx, "Select")
	if err := mbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MissionBatchQuery, *MissionBatchSelect](ctx, mbs.MissionBatchQuery, mbs, mbs.inters, v)
}

func (mbs *MissionBatchSelect) sqlScan(ctx context.Context, root *MissionBatchQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mbs.fns))
	for _, fn := range mbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mbs *MissionBatchSelect) Modify(modifiers ...func(s *sql.Selector)) *MissionBatchSelect {
	mbs.modifiers = append(mbs.modifiers, modifiers...)
	return mbs
}

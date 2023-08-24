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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/costbill"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionbatch"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionconsumeorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionproduceorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// MissionConsumeOrderQuery is the builder for querying MissionConsumeOrder entities.
type MissionConsumeOrderQuery struct {
	config
	ctx                      *QueryContext
	order                    []missionconsumeorder.OrderOption
	inters                   []Interceptor
	predicates               []predicate.MissionConsumeOrder
	withUser                 *UserQuery
	withCostBills            *CostBillQuery
	withMissionProduceOrders *MissionProduceOrderQuery
	withMissionBatch         *MissionBatchQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MissionConsumeOrderQuery builder.
func (mcoq *MissionConsumeOrderQuery) Where(ps ...predicate.MissionConsumeOrder) *MissionConsumeOrderQuery {
	mcoq.predicates = append(mcoq.predicates, ps...)
	return mcoq
}

// Limit the number of records to be returned by this query.
func (mcoq *MissionConsumeOrderQuery) Limit(limit int) *MissionConsumeOrderQuery {
	mcoq.ctx.Limit = &limit
	return mcoq
}

// Offset to start from.
func (mcoq *MissionConsumeOrderQuery) Offset(offset int) *MissionConsumeOrderQuery {
	mcoq.ctx.Offset = &offset
	return mcoq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mcoq *MissionConsumeOrderQuery) Unique(unique bool) *MissionConsumeOrderQuery {
	mcoq.ctx.Unique = &unique
	return mcoq
}

// Order specifies how the records should be ordered.
func (mcoq *MissionConsumeOrderQuery) Order(o ...missionconsumeorder.OrderOption) *MissionConsumeOrderQuery {
	mcoq.order = append(mcoq.order, o...)
	return mcoq
}

// QueryUser chains the current query on the "user" edge.
func (mcoq *MissionConsumeOrderQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: mcoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mcoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mcoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionconsumeorder.Table, missionconsumeorder.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, missionconsumeorder.UserTable, missionconsumeorder.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(mcoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCostBills chains the current query on the "cost_bills" edge.
func (mcoq *MissionConsumeOrderQuery) QueryCostBills() *CostBillQuery {
	query := (&CostBillClient{config: mcoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mcoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mcoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionconsumeorder.Table, missionconsumeorder.FieldID, selector),
			sqlgraph.To(costbill.Table, costbill.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, missionconsumeorder.CostBillsTable, missionconsumeorder.CostBillsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mcoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMissionProduceOrders chains the current query on the "mission_produce_orders" edge.
func (mcoq *MissionConsumeOrderQuery) QueryMissionProduceOrders() *MissionProduceOrderQuery {
	query := (&MissionProduceOrderClient{config: mcoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mcoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mcoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionconsumeorder.Table, missionconsumeorder.FieldID, selector),
			sqlgraph.To(missionproduceorder.Table, missionproduceorder.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, missionconsumeorder.MissionProduceOrdersTable, missionconsumeorder.MissionProduceOrdersColumn),
		)
		fromU = sqlgraph.SetNeighbors(mcoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMissionBatch chains the current query on the "mission_batch" edge.
func (mcoq *MissionConsumeOrderQuery) QueryMissionBatch() *MissionBatchQuery {
	query := (&MissionBatchClient{config: mcoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mcoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mcoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionconsumeorder.Table, missionconsumeorder.FieldID, selector),
			sqlgraph.To(missionbatch.Table, missionbatch.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, missionconsumeorder.MissionBatchTable, missionconsumeorder.MissionBatchColumn),
		)
		fromU = sqlgraph.SetNeighbors(mcoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MissionConsumeOrder entity from the query.
// Returns a *NotFoundError when no MissionConsumeOrder was found.
func (mcoq *MissionConsumeOrderQuery) First(ctx context.Context) (*MissionConsumeOrder, error) {
	nodes, err := mcoq.Limit(1).All(setContextOp(ctx, mcoq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{missionconsumeorder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mcoq *MissionConsumeOrderQuery) FirstX(ctx context.Context) *MissionConsumeOrder {
	node, err := mcoq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MissionConsumeOrder ID from the query.
// Returns a *NotFoundError when no MissionConsumeOrder ID was found.
func (mcoq *MissionConsumeOrderQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mcoq.Limit(1).IDs(setContextOp(ctx, mcoq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{missionconsumeorder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mcoq *MissionConsumeOrderQuery) FirstIDX(ctx context.Context) int64 {
	id, err := mcoq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MissionConsumeOrder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MissionConsumeOrder entity is found.
// Returns a *NotFoundError when no MissionConsumeOrder entities are found.
func (mcoq *MissionConsumeOrderQuery) Only(ctx context.Context) (*MissionConsumeOrder, error) {
	nodes, err := mcoq.Limit(2).All(setContextOp(ctx, mcoq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{missionconsumeorder.Label}
	default:
		return nil, &NotSingularError{missionconsumeorder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mcoq *MissionConsumeOrderQuery) OnlyX(ctx context.Context) *MissionConsumeOrder {
	node, err := mcoq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MissionConsumeOrder ID in the query.
// Returns a *NotSingularError when more than one MissionConsumeOrder ID is found.
// Returns a *NotFoundError when no entities are found.
func (mcoq *MissionConsumeOrderQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mcoq.Limit(2).IDs(setContextOp(ctx, mcoq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{missionconsumeorder.Label}
	default:
		err = &NotSingularError{missionconsumeorder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mcoq *MissionConsumeOrderQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := mcoq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MissionConsumeOrders.
func (mcoq *MissionConsumeOrderQuery) All(ctx context.Context) ([]*MissionConsumeOrder, error) {
	ctx = setContextOp(ctx, mcoq.ctx, "All")
	if err := mcoq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MissionConsumeOrder, *MissionConsumeOrderQuery]()
	return withInterceptors[[]*MissionConsumeOrder](ctx, mcoq, qr, mcoq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mcoq *MissionConsumeOrderQuery) AllX(ctx context.Context) []*MissionConsumeOrder {
	nodes, err := mcoq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MissionConsumeOrder IDs.
func (mcoq *MissionConsumeOrderQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if mcoq.ctx.Unique == nil && mcoq.path != nil {
		mcoq.Unique(true)
	}
	ctx = setContextOp(ctx, mcoq.ctx, "IDs")
	if err = mcoq.Select(missionconsumeorder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mcoq *MissionConsumeOrderQuery) IDsX(ctx context.Context) []int64 {
	ids, err := mcoq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mcoq *MissionConsumeOrderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mcoq.ctx, "Count")
	if err := mcoq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mcoq, querierCount[*MissionConsumeOrderQuery](), mcoq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mcoq *MissionConsumeOrderQuery) CountX(ctx context.Context) int {
	count, err := mcoq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mcoq *MissionConsumeOrderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mcoq.ctx, "Exist")
	switch _, err := mcoq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mcoq *MissionConsumeOrderQuery) ExistX(ctx context.Context) bool {
	exist, err := mcoq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MissionConsumeOrderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mcoq *MissionConsumeOrderQuery) Clone() *MissionConsumeOrderQuery {
	if mcoq == nil {
		return nil
	}
	return &MissionConsumeOrderQuery{
		config:                   mcoq.config,
		ctx:                      mcoq.ctx.Clone(),
		order:                    append([]missionconsumeorder.OrderOption{}, mcoq.order...),
		inters:                   append([]Interceptor{}, mcoq.inters...),
		predicates:               append([]predicate.MissionConsumeOrder{}, mcoq.predicates...),
		withUser:                 mcoq.withUser.Clone(),
		withCostBills:            mcoq.withCostBills.Clone(),
		withMissionProduceOrders: mcoq.withMissionProduceOrders.Clone(),
		withMissionBatch:         mcoq.withMissionBatch.Clone(),
		// clone intermediate query.
		sql:  mcoq.sql.Clone(),
		path: mcoq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (mcoq *MissionConsumeOrderQuery) WithUser(opts ...func(*UserQuery)) *MissionConsumeOrderQuery {
	query := (&UserClient{config: mcoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mcoq.withUser = query
	return mcoq
}

// WithCostBills tells the query-builder to eager-load the nodes that are connected to
// the "cost_bills" edge. The optional arguments are used to configure the query builder of the edge.
func (mcoq *MissionConsumeOrderQuery) WithCostBills(opts ...func(*CostBillQuery)) *MissionConsumeOrderQuery {
	query := (&CostBillClient{config: mcoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mcoq.withCostBills = query
	return mcoq
}

// WithMissionProduceOrders tells the query-builder to eager-load the nodes that are connected to
// the "mission_produce_orders" edge. The optional arguments are used to configure the query builder of the edge.
func (mcoq *MissionConsumeOrderQuery) WithMissionProduceOrders(opts ...func(*MissionProduceOrderQuery)) *MissionConsumeOrderQuery {
	query := (&MissionProduceOrderClient{config: mcoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mcoq.withMissionProduceOrders = query
	return mcoq
}

// WithMissionBatch tells the query-builder to eager-load the nodes that are connected to
// the "mission_batch" edge. The optional arguments are used to configure the query builder of the edge.
func (mcoq *MissionConsumeOrderQuery) WithMissionBatch(opts ...func(*MissionBatchQuery)) *MissionConsumeOrderQuery {
	query := (&MissionBatchClient{config: mcoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mcoq.withMissionBatch = query
	return mcoq
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
//	client.MissionConsumeOrder.Query().
//		GroupBy(missionconsumeorder.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (mcoq *MissionConsumeOrderQuery) GroupBy(field string, fields ...string) *MissionConsumeOrderGroupBy {
	mcoq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MissionConsumeOrderGroupBy{build: mcoq}
	grbuild.flds = &mcoq.ctx.Fields
	grbuild.label = missionconsumeorder.Label
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
//	client.MissionConsumeOrder.Query().
//		Select(missionconsumeorder.FieldCreatedBy).
//		Scan(ctx, &v)
func (mcoq *MissionConsumeOrderQuery) Select(fields ...string) *MissionConsumeOrderSelect {
	mcoq.ctx.Fields = append(mcoq.ctx.Fields, fields...)
	sbuild := &MissionConsumeOrderSelect{MissionConsumeOrderQuery: mcoq}
	sbuild.label = missionconsumeorder.Label
	sbuild.flds, sbuild.scan = &mcoq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MissionConsumeOrderSelect configured with the given aggregations.
func (mcoq *MissionConsumeOrderQuery) Aggregate(fns ...AggregateFunc) *MissionConsumeOrderSelect {
	return mcoq.Select().Aggregate(fns...)
}

func (mcoq *MissionConsumeOrderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mcoq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mcoq); err != nil {
				return err
			}
		}
	}
	for _, f := range mcoq.ctx.Fields {
		if !missionconsumeorder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if mcoq.path != nil {
		prev, err := mcoq.path(ctx)
		if err != nil {
			return err
		}
		mcoq.sql = prev
	}
	return nil
}

func (mcoq *MissionConsumeOrderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MissionConsumeOrder, error) {
	var (
		nodes       = []*MissionConsumeOrder{}
		_spec       = mcoq.querySpec()
		loadedTypes = [4]bool{
			mcoq.withUser != nil,
			mcoq.withCostBills != nil,
			mcoq.withMissionProduceOrders != nil,
			mcoq.withMissionBatch != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MissionConsumeOrder).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MissionConsumeOrder{config: mcoq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mcoq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mcoq.withUser; query != nil {
		if err := mcoq.loadUser(ctx, query, nodes, nil,
			func(n *MissionConsumeOrder, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := mcoq.withCostBills; query != nil {
		if err := mcoq.loadCostBills(ctx, query, nodes,
			func(n *MissionConsumeOrder) { n.Edges.CostBills = []*CostBill{} },
			func(n *MissionConsumeOrder, e *CostBill) { n.Edges.CostBills = append(n.Edges.CostBills, e) }); err != nil {
			return nil, err
		}
	}
	if query := mcoq.withMissionProduceOrders; query != nil {
		if err := mcoq.loadMissionProduceOrders(ctx, query, nodes,
			func(n *MissionConsumeOrder) { n.Edges.MissionProduceOrders = []*MissionProduceOrder{} },
			func(n *MissionConsumeOrder, e *MissionProduceOrder) {
				n.Edges.MissionProduceOrders = append(n.Edges.MissionProduceOrders, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := mcoq.withMissionBatch; query != nil {
		if err := mcoq.loadMissionBatch(ctx, query, nodes, nil,
			func(n *MissionConsumeOrder, e *MissionBatch) { n.Edges.MissionBatch = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mcoq *MissionConsumeOrderQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*MissionConsumeOrder, init func(*MissionConsumeOrder), assign func(*MissionConsumeOrder, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MissionConsumeOrder)
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
func (mcoq *MissionConsumeOrderQuery) loadCostBills(ctx context.Context, query *CostBillQuery, nodes []*MissionConsumeOrder, init func(*MissionConsumeOrder), assign func(*MissionConsumeOrder, *CostBill)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*MissionConsumeOrder)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(costbill.FieldReasonID)
	}
	query.Where(predicate.CostBill(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(missionconsumeorder.CostBillsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ReasonID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "reason_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mcoq *MissionConsumeOrderQuery) loadMissionProduceOrders(ctx context.Context, query *MissionProduceOrderQuery, nodes []*MissionConsumeOrder, init func(*MissionConsumeOrder), assign func(*MissionConsumeOrder, *MissionProduceOrder)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*MissionConsumeOrder)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(missionproduceorder.FieldMissionConsumeOrderID)
	}
	query.Where(predicate.MissionProduceOrder(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(missionconsumeorder.MissionProduceOrdersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MissionConsumeOrderID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "mission_consume_order_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mcoq *MissionConsumeOrderQuery) loadMissionBatch(ctx context.Context, query *MissionBatchQuery, nodes []*MissionConsumeOrder, init func(*MissionConsumeOrder), assign func(*MissionConsumeOrder, *MissionBatch)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MissionConsumeOrder)
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

func (mcoq *MissionConsumeOrderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mcoq.querySpec()
	_spec.Node.Columns = mcoq.ctx.Fields
	if len(mcoq.ctx.Fields) > 0 {
		_spec.Unique = mcoq.ctx.Unique != nil && *mcoq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mcoq.driver, _spec)
}

func (mcoq *MissionConsumeOrderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(missionconsumeorder.Table, missionconsumeorder.Columns, sqlgraph.NewFieldSpec(missionconsumeorder.FieldID, field.TypeInt64))
	_spec.From = mcoq.sql
	if unique := mcoq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mcoq.path != nil {
		_spec.Unique = true
	}
	if fields := mcoq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, missionconsumeorder.FieldID)
		for i := range fields {
			if fields[i] != missionconsumeorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mcoq.withUser != nil {
			_spec.Node.AddColumnOnce(missionconsumeorder.FieldUserID)
		}
		if mcoq.withMissionBatch != nil {
			_spec.Node.AddColumnOnce(missionconsumeorder.FieldMissionBatchID)
		}
	}
	if ps := mcoq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mcoq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mcoq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mcoq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mcoq *MissionConsumeOrderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mcoq.driver.Dialect())
	t1 := builder.Table(missionconsumeorder.Table)
	columns := mcoq.ctx.Fields
	if len(columns) == 0 {
		columns = missionconsumeorder.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mcoq.sql != nil {
		selector = mcoq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mcoq.ctx.Unique != nil && *mcoq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mcoq.predicates {
		p(selector)
	}
	for _, p := range mcoq.order {
		p(selector)
	}
	if offset := mcoq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mcoq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MissionConsumeOrderGroupBy is the group-by builder for MissionConsumeOrder entities.
type MissionConsumeOrderGroupBy struct {
	selector
	build *MissionConsumeOrderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mcogb *MissionConsumeOrderGroupBy) Aggregate(fns ...AggregateFunc) *MissionConsumeOrderGroupBy {
	mcogb.fns = append(mcogb.fns, fns...)
	return mcogb
}

// Scan applies the selector query and scans the result into the given value.
func (mcogb *MissionConsumeOrderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mcogb.build.ctx, "GroupBy")
	if err := mcogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MissionConsumeOrderQuery, *MissionConsumeOrderGroupBy](ctx, mcogb.build, mcogb, mcogb.build.inters, v)
}

func (mcogb *MissionConsumeOrderGroupBy) sqlScan(ctx context.Context, root *MissionConsumeOrderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mcogb.fns))
	for _, fn := range mcogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mcogb.flds)+len(mcogb.fns))
		for _, f := range *mcogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mcogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mcogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MissionConsumeOrderSelect is the builder for selecting fields of MissionConsumeOrder entities.
type MissionConsumeOrderSelect struct {
	*MissionConsumeOrderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mcos *MissionConsumeOrderSelect) Aggregate(fns ...AggregateFunc) *MissionConsumeOrderSelect {
	mcos.fns = append(mcos.fns, fns...)
	return mcos
}

// Scan applies the selector query and scans the result into the given value.
func (mcos *MissionConsumeOrderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mcos.ctx, "Select")
	if err := mcos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MissionConsumeOrderQuery, *MissionConsumeOrderSelect](ctx, mcos.MissionConsumeOrderQuery, mcos, mcos.inters, v)
}

func (mcos *MissionConsumeOrderSelect) sqlScan(ctx context.Context, root *MissionConsumeOrderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mcos.fns))
	for _, fn := range mcos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mcos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mcos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
